package tcp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"logging"
	"net"
	"response"
	"storage"
	"strings"
)

// Internal to TCP
type context struct {
	uID    string
	conn   net.Conn
	store  storage.Store
	cancel chan struct{}
	stop   chan struct{}
}

// Cancelled call for cancel go rotines
func (c *context) cancelled() bool {
	select {
	case <-c.cancel:
		return true
	default:
		return false
	}
}

// Stopped command for clients to disconnect
func (c *context) stopped() bool {
	select {
	case <-c.stop:
		return true
	default:
		return false
	}
}

// Cmd commands for tcp
const (
	CmdGet    string = "get"
	CmdDelete string = "delete"
	CmdCreate string = "create"
	CmdUpdate string = "update"
	CmdStop   string = "stop"
)

var acceptableCommands = [...]string{CmdGet, CmdCreate, CmdUpdate, CmdDelete, CmdStop}

// Executor of commands
func executor(ctx context, in chan interface{}, errc chan error) chan interface{} {

	errCtx := "TCP::Executor"

	out := make(chan interface{}, 1)

	go func() {

		defer close(out)
		defer close(errc)

		input := <-in

		if ctx.cancelled() {
			logging.Msg(
				ctx.uID,
				"Executor",
				logging.Cancelled,
				"Execution was cancelled from another goroutine")

			return
		}

		paramMap, ok := input.(map[string]interface{})

		if ok {
			cmd := strings.ToLower(paramMap[ParamCmd].(string))
			key := paramMap[ParamKey]

			var resp interface{} = struct{}{}
			var err error = nil
			resChan := make(chan storage.Result)

			if cmd == CmdGet || cmd == CmdDelete {

				var command string
				if cmd == CmdGet {
					command = storage.CmdGetRecord
				} else if cmd == CmdDelete {
					command = storage.CmdDeleteRecord
				}

				ctx.store.Requests <- storage.Request{
					Cmd:      command,
					Key:      key.(string),
					Response: resChan}

			} else if cmd == CmdCreate || cmd == CmdUpdate {

				var command string
				if cmd == CmdCreate {
					command = storage.CmdAddRecord
				} else if cmd == CmdUpdate {
					command = storage.CmdUpdateRecord
				}

				body := paramMap[ParamBody]
				b, ok := body.(string)
				if !ok {
					logging.Msg(ctx.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrTypeConversion))
					close(ctx.cancel)
					errc <- response.New(response.ErrTypeConversion, errCtx)
					return
				}

				ctx.store.Requests <- storage.Request{
					Cmd:      command,
					Key:      key.(string),
					Value:    &storage.StoreValue{CValue: []byte(b)},
					Response: resChan}

			} else if cmd == CmdStop {
				errc <- response.New(response.ErrStopCommand, errCtx)
				close(ctx.stop)
				return
			} else {
				errc <- response.New(response.ErrMissingCommand, errCtx)
				return
			}

			// Wait for answer from Store
			res := <-resChan
			resp, err = res.Value, res.Err

			// Check error from Store
			if err != nil {
				logging.Msg(ctx.uID, errCtx, logging.Error, err.Error())
				errc <- err
				close(ctx.cancel)
				return
			}

			out <- resp

		} else {
			logging.Msg(ctx.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrTypeConversion))
			errc <- response.New(response.ErrTypeConversion, errCtx)
		}
	}()

	return out
}

// Make sure the cmd are passing the correct values to the executor
func validator(ctx context, in chan interface{}, errc chan error) chan interface{} {

	errCtx := "TPC::Validator"

	out := make(chan interface{}, 1)

	go func() {
		defer close(out)

		paramMap := <-in

		if ctx.cancelled() {
			logging.Msg(
				ctx.uID,
				"Validator",
				logging.Cancelled,
				"Execution was cancelled from another goroutine")

			return
		}

		parsedData, ok := paramMap.(map[string]interface{})

		// if type conversion fails
		if !ok {
			logging.Msg(ctx.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrTypeConversion))
			errc <- response.New(response.ErrTypeConversion, errCtx)
			close(ctx.cancel)
			return
		}

		// validate what comes up from cmd
		if paramMap == nil || len(parsedData) == 0 && parsedData[ParamCmd] == nil {
			logging.Msg(ctx.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrArgumentException))
			errc <- response.New(response.ErrArgumentException, errCtx)
			close(ctx.cancel)
			return
		}

		// validate commands
		cmd, ok := parsedData[ParamCmd].(string)
		validCmd := false
		for c := range acceptableCommands {
			if cmd == acceptableCommands[c] {
				validCmd = true
			}
		}

		if !validCmd || !ok {
			logging.Msg(ctx.uID, errCtx, logging.Error, response.ErrorCodeText(response.ErrArgumentException))
			errc <- response.New(response.ErrArgumentException, errCtx)
			close(ctx.cancel)
			return
		}

		// validate commands with required params
		key, ok := parsedData[ParamKey].(string)

		switch strings.ToLower(cmd) {
		case "get", "delete":
			if !ok || key == "" {
				errc <- response.New(response.ErrArgumentException, errCtx)
				close(ctx.cancel)
				return
			}
		case "create", "update":
			body := parsedData[ParamBody]
			if !ok || key == "" || body == nil || body == "" {
				errc <- response.New(response.ErrArgumentException, errCtx)
				close(ctx.cancel)
				return
			}
		}

		out <- paramMap
	}()

	return out
}

// Main handler
func handler(ctx context) {
	errCtx := "TCP::Handler"

	reader := bufio.NewReader(ctx.conn)

	defer ctx.conn.Close()

	for {
		if ctx.stopped() {
			logging.Msg(
				ctx.uID,
				"Handler",
				logging.Stopped,
				"Client requested STOP execution")

			return
		}

		// Read / wait for new requests
		input, err := reader.ReadString('\n')

		if err == io.EOF {
			logging.Msg(
				ctx.uID,
				"Handler",
				logging.Error,
				"Error reading from bufio client EOF")

			return
		}

		// Remove delimiter character \n
		input = strings.TrimSuffix(input, "\n")

		// Each request will have a new UUID
		ctx.uID = logging.UUID()

		// New chan per connection
		errc := make(chan error, 1)
		ctx.cancel = make(chan struct{})

		logging.MsgFields(
			ctx.uID,
			"Handler",
			logging.Input,
			logging.Field{{logging.Payload, input}})

		if err != nil {
			logging.Msgf(
				ctx.uID,
				"Handler",
				logging.Error,
				"Error reading from bufio client request %+v", err)

			errc <- response.New(response.ErrTCPNetConnection, errCtx)
			close(ctx.cancel)
		}

		// Execution
		out := parser(ctx, input, errc)
		out = validator(ctx, out, errc)
		out = executor(ctx, out, errc)
		writer(ctx, out, errc)
	}
}

func parser(ctx context, input string, errc chan error) chan interface{} {

	out := make(chan interface{}, 1)

	go func() {

		defer close(out)

		if ctx.cancelled() {
			logging.Msg(
				ctx.uID,
				"ParseInput",
				logging.Cancelled,
				"Execution was cancelled from another goroutine")

			return
		}

		parsedData, err := Parser{Content: string(input)}.ParsePipe()
		if err != nil {
			logging.Msgf(
				ctx.uID,
				"ParseInput",
				logging.Error,
				"Error parsing client request %+v", err)

			errc <- err
			close(ctx.cancel)
			return
		}

		out <- parsedData
	}()

	return out
}

func writer(ctx context, out chan interface{}, err chan error) {
	select {
	case payload := <-out:
		sv, ok := payload.(*storage.StoreValue)
		if ok {
			successResponse(ctx, sv)
		} else {
			errorResponse(ctx, response.New(response.ErrTypeConversion, "Writer"))
		}
	case errc := <-err:
		e, ok := errc.(response.APIError)
		if ok {
			errorResponse(ctx, e)
		} else {
			errorResponse(ctx, errc)
		}
	}
}

func successResponse(ctx context, sv *storage.StoreValue) {
	// Signal client to read the response with \n
	sbPayload := append(sv.CValue, "\n"...)

	io.Copy(ctx.conn, bytes.NewBuffer(sbPayload))

	logging.MsgFields(
		ctx.uID,
		"SuccessResponse",
		logging.Output,
		logging.Field{{logging.Payload, sv}})
}

func errorResponse(ctx context, out interface{}) {
	bytePayload, err := json.Marshal(out)

	if err != nil {
		logging.Msgf(
			ctx.uID,
			"ErrorResponse",
			logging.Error,
			"Error: %+v", err)

		// Signal client to read the response with \n
		errPayload := err.Error() + "\n"

		io.Copy(ctx.conn, bytes.NewBufferString(errPayload))

	} else {
		// Signal client to read the response with \n
		bytePayload = append(bytePayload, "\n"...)

		io.Copy(ctx.conn, bytes.NewBuffer(bytePayload))
	}

	logging.MsgFields(
		ctx.uID,
		"ErrorResponse",
		logging.Output,
		logging.Field{{logging.Payload, out}})
}

// Listen net
func Listen(domain string, port string) net.Listener {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", domain, port))
	if err != nil {
		logging.Msgf(
			logging.UUID(),
			"Listen",
			logging.Error,
			"Start Listener error: %+v", err)

		return nil
	}

	return l
}

// Accept conn
func Accept(l net.Listener, storage storage.Store) {
	for {
		uuid := logging.UUID()
		conn, err := l.Accept()
		if err != nil {
			logging.Msgf(
				uuid,
				"Accept",
				logging.Error,
				"Start Accept error: %+v", err)

			return
		}

		logging.Msg(uuid, "Accept", "New connection", "Accepting new connection")

		// Create new context per connected client
		ctx := context{
			conn:   conn,
			store:  storage,
			cancel: make(chan struct{}),
			stop:   make(chan struct{})}

		go handler(ctx)
	}
}

const tcpport = "9001"
const tcpdomain = "localhost"

// Start TCP server
func Start(store storage.Store, done chan bool) {
	Accept(Listen(tcpdomain, tcpport), store)
	done <- true
}
