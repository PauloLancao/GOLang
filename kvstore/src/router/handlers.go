package router

import (
	"encoding/json"
	"logging"
	"net/http"
	"response"
	"storage"
)

// Cmd commands for rest api
const (
	CmdGet    string = "get"
	CmdDelete string = "delete"
	CmdCreate string = "create"
	CmdUpdate string = "update"
)

// KVS func type defined for all handler calls
type kvs func(s storage.Store, in interface{}) (out interface{}, errc error)

// StoreExecutor func
func StoreExecutor(rc *Context, cmd string) func(s storage.Store, in interface{}) (out interface{}, errc error) {

	if cmd == CmdCreate || cmd == CmdUpdate {
		rc.WithBuilder(Builder{ExtractParams: true, ExtractBody: true})
	} else if cmd == CmdGet || cmd == CmdDelete {
		rc.WithBuilder(Builder{ExtractParams: true})
	}

	return func(s storage.Store, in interface{}) (out interface{}, errc error) {

		errCtx := "Handlers::StoreExecutor"

		params, ok := in.(map[string]interface{})
		if !ok {
			logging.Msg(rc.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrTypeConversion))
			rc.CtxCancel()
			return nil, response.New(response.ErrTypeConversion, errCtx)
		}

		key := params[ParamKey]
		if key == nil {
			logging.Msg(rc.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrMissingRequestParams))
			rc.CtxCancel()
			return nil, response.New(response.ErrMissingRequestParams, errCtx)
		}

		sKey := key.(string)
		resChan := make(chan storage.Result)
		var storeReq storage.Request

		if cmd == CmdGet || cmd == CmdDelete {

			if cmd == CmdGet {
				cmd = storage.CmdGetRecord
			} else if cmd == CmdDelete {
				cmd = storage.CmdDeleteRecord
			}

			storeReq = storage.Request{
				Cmd:      cmd,
				Key:      sKey,
				Response: resChan}

		} else if cmd == CmdCreate || cmd == CmdUpdate {

			payload := params[ParamBody]
			if payload == nil {
				logging.Msg(rc.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrMissingRequestParams))
				rc.CtxCancel()
				return nil, response.New(response.ErrMissingRequestParams, errCtx)
			}

			b, ok := payload.(*[]byte)
			if !ok {
				logging.Msg(rc.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrTypeConversion))
				rc.CtxCancel()
				return nil, response.New(response.ErrTypeConversion, errCtx)
			}

			if cmd == CmdCreate {
				cmd = storage.CmdAddRecord
			} else if cmd == CmdUpdate {
				cmd = storage.CmdUpdateRecord
			}

			storeReq = storage.Request{
				Cmd:      cmd,
				Key:      sKey,
				Value:    &storage.StoreValue{CType: rc.GetContentType(), CValue: *b},
				Response: resChan}

		} else {
			logging.Msg(rc.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrMissingCommand))
			return nil, response.New(response.ErrMissingCommand, errCtx)
		}

		s.Requests <- storeReq
		res := <-resChan

		return res.Value, res.Err
	}
}

// Handler of all kvs
func Handler(rc *Context, s storage.Store, storeExecutor kvs) {

	// Make sure all resources are cleanup after leaving
	cancelFunc := rc.CtxCancel
	defer cancelFunc()

	// Err channel shared across
	errc := make(chan error, 1)

	// Execution call
	out := make(chan interface{}, 1)

	go func() {
		defer close(out)
		defer close(errc)

		inb := <-httpExtractor(rc, errc)

		logging.MsgFields(
			rc.uID,
			"HttpExtractor",
			logging.Input,
			logging.Field{{logging.Payload, inb}})

		if rc.Cancelled() {
			logging.Msg(
				rc.uID,
				"Handler",
				logging.Cancelled,
				"kvs execution was cancelled from another goroutine")

			return
		}

		res, err := storeExecutor(s, inb)

		if err != nil {
			errc <- err
			cancelFunc()
			return
		}

		select {
		case out <- res:
		case <-rc.Ctx.Done():
			return
		}
	}()

	HTTPResponse(rc, out, errc)
}

// HTTPResponse handler
func HTTPResponse(rc *Context, out chan interface{}, err chan error) {

	select {
	case payload := <-out:
		sv, ok := payload.(*storage.StoreValue)
		if ok {
			SuccessResponse(rc, sv)
		} else {
			ErrorResponse(rc, http.StatusInternalServerError, response.New(response.ErrTypeConversion, "HTTPResponse"))
		}
	case errc := <-err:
		e, ok := errc.(response.APIError)
		if ok {
			ErrorResponse(rc, e.StatusCode(), e)
		} else {
			ErrorResponse(rc, http.StatusBadRequest, errc)
		}
	}
}

// SuccessResponse func
func SuccessResponse(rc *Context, sv *storage.StoreValue) {
	rc.W.Header().Set("Content-Type", sv.CType)
	rc.W.WriteHeader(http.StatusOK)
	rc.W.Write(sv.CValue)

	logResponseOutput(rc.uID, sv.CValue)
}

// ErrorResponse func
func ErrorResponse(rc *Context, statusCode int, payload interface{}) {

	bytePayload, err := json.Marshal(payload)

	rc.W.Header().Set("Content-Type", "application/json")

	if err != nil {
		logging.Msgf(
			rc.uID,
			"ErrorResponseJSON",
			logging.Error, "ErrorResponseJSON:: error marshal payload: %+v", err)

		rc.W.WriteHeader(http.StatusInternalServerError)
		rc.W.Write([]byte(err.Error()))
	} else {
		rc.W.WriteHeader(statusCode)
		rc.W.Write(bytePayload)
	}

	logResponseOutput(rc.uID, bytePayload)
}

func logResponseOutput(uuid string, logResponse interface{}) {
	logging.MsgFields(
		uuid,
		"HTTPResponse",
		logging.Output,
		logging.Field{{logging.Payload, logResponse}})
}

// HTTPExtractor generic extractor for parameters and body
func httpExtractor(rc *Context, errc chan error) <-chan interface{} {
	out := make(chan interface{}, 1)
	errCtx := "Handlers::HTTPExtractor"

	go func() {
		defer close(out)

		reqMap := make(map[string]interface{})
		reqBuilder := rc.Builder

		// Path variables extractor
		if reqBuilder.ExtractParams {
			params := rc.PathVars()

			if params == nil || len(params) == 0 {
				logging.Msg(rc.uID, "HttpExtractor", logging.Error, response.ErrorCodeText(response.ErrMissingRequestParams))
				errc <- response.New(response.ErrMissingRequestParams, errCtx)
				rc.CtxCancel()
				return
			}

			reqMap[ParamKey] = params[ParamKey]
		}

		// Body extractor
		if reqBuilder.ExtractBody {
			decBody, err := rc.Body()
			if err != nil {
				logging.Msg(rc.uID, "HttpExtractor", logging.Error, response.ErrorCodeText(response.ErrDecoding))
				errc <- response.New(response.ErrDecoding, errCtx)
				rc.CtxCancel()
				return
			}

			reqMap[ParamBody] = &decBody
		}

		// Send the data to the output channel but return early
		// if the context has been cancelled.
		select {
		case out <- reqMap:
		case <-rc.Ctx.Done():
			return
		}
	}()

	return out
}
