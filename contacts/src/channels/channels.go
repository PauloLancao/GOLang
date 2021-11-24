package channels

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/paulolancao/go-contacts/gerrors"
	"github.com/paulolancao/go-contacts/models"
	"github.com/paulolancao/go-contacts/repository"
	"github.com/paulolancao/go-contacts/utils"
)

// RequestParams struct used from Extractor http request params / body
// to be passed to processing channels
type RequestParams struct {
	Params map[string]interface{}
}

// RequestExtractor struct used to extract all data from http request
type RequestExtractor struct {
	ExtractBody   bool
	ExtractParams bool
	Body          interface{}
	Params        []string
}

// HTTPExtractor generic extractor for parameters and body
func HTTPExtractor(ctx context.Context, e *RequestExtractor, r *http.Request) (
	<-chan interface{}, <-chan error) {

	out := make(chan interface{}, 1)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errc)

		reqMap := make(map[string]interface{})

		if e.ExtractParams && len(e.Params) > 0 {
			params := mux.Vars(r)
			for _, s := range e.Params {
				reqMap[s] = params[s]
			}
		}

		if e.ExtractBody && e.Body != nil {
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&e.Body)

			if err != nil {
				log.Print("HTTPExtractor::failed to decode payload body")

				// Handle an error that occurs during the goroutine.
				errc <- gerrors.New(http.StatusUnprocessableEntity, gerrors.ErrDecoding)
				return
			}

			defer r.Body.Close()

			reqMap["body"] = &e.Body
		}

		// Send the data to the output channel but return early
		// if the context has been cancelled.
		select {
		case out <- &RequestParams{Params: reqMap}:
		case <-ctx.Done():
			return
		}
	}()

	return out, errc
}

// HTTPValidator channel makes sure the payload is correct to be processed
func HTTPValidator(ctx context.Context, in <-chan interface{}) (
	<-chan interface{}, <-chan error) {

	out := make(chan interface{}, 1)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errc)

		inb := <-in
		params, ok := inb.(*RequestParams)
		if !ok {
			log.Print("HTTPValidator::missing request params")

			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrMissingRequestParams)
			return
		}

		contact := params.Params["body"]
		if contact == nil {
			log.Print("HTTPValidator::failed to decode payload body")

			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrMissingRequestParams)
			return
		}

		cInt := *contact.(*interface{})
		c, ok := cInt.(*models.Contact)
		if !ok {
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrTypeConversion)
			return
		}

		if validErrs := c.Validate(); len(validErrs) > 0 {
			errc <- gerrors.NewErrors(http.StatusBadRequest, gerrors.ErrValidation, validErrs)
			return
		}

		// Send the data to the output channel but return early
		// if the context has been cancelled.
		select {
		case out <- inb:
		case <-ctx.Done():
			return
		}
	}()

	return out, errc
}

// HTTPResponse channel response
func HTTPResponse(w http.ResponseWriter, out <-chan interface{}, err error, msg string) {
	if err != nil {
		e, ok := err.(gerrors.Channelerror)
		response := map[string]interface{}{"error": e.Error(), "errorcode": e.ErrorCode(), "errors": e.Errors()}
		if ok {
			log.Printf("%s::%s - %d", msg, e.Error(), e.StatusCode())
			utils.RespondWithError(w, e.StatusCode(), response)
		} else {
			log.Printf("%s::%+v", msg, err)
			utils.RespondWithError(w, http.StatusBadRequest, response)
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, <-out)
}

// CreateContact createContact
func CreateContact(ctx context.Context, in <-chan interface{}, r repository.Repository) (
	<-chan interface{}, <-chan error) {

	out := make(chan interface{}, 1)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errc)

		inb := <-in
		params, ok := inb.(*RequestParams)
		if !ok {
			log.Print("CreateContact::missing request params")

			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrMissingRequestParams)
			return
		}

		contact := params.Params["body"]
		if contact == nil {
			log.Print("CreateContact::missing contact body")

			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrMissingRequestParams)
			return
		}

		cInt := *contact.(*interface{})
		c, ok := cInt.(*models.Contact)
		if !ok {
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrTypeConversion)
			return
		}

		err := r.CreateContact(*c)

		if err != nil {
			// Handle an error that occurs during the goroutine.
			errc <- err
			return
		}

		// Send the data to the output channel but return early
		// if the context has been cancelled.
		select {
		case out <- struct{}{}:
		case <-ctx.Done():
			return
		}
	}()

	return out, errc
}

// UpdateContact update Contact
func UpdateContact(ctx context.Context, in <-chan interface{}, r repository.Repository) (
	<-chan interface{}, <-chan error) {

	out := make(chan interface{}, 1)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errc)

		inb := <-in
		params, ok := inb.(*RequestParams)
		if !ok {
			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrMissingRequestParams)
			return
		}

		contact := params.Params["body"]
		id := params.Params["id"]
		if contact == nil || id == nil {
			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrMissingRequestParams)
			return
		}

		aID, err := strconv.Atoi(id.(string))
		if err != nil {
			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrTypeConversion)
			return
		}

		cInt := *contact.(*interface{})
		c := cInt.(*models.Contact)
		c.ID = uint(aID)

		err = r.UpdateContact(*c)

		if err != nil {
			// Handle an error that occurs during the goroutine.
			errc <- err
			return
		}

		// Send the data to the output channel but return early
		// if the context has been cancelled.
		select {
		case out <- c:
		case <-ctx.Done():
			return
		}
	}()

	return out, errc
}

// DeleteContact update Contact
func DeleteContact(ctx context.Context, in <-chan interface{}, r repository.Repository) (
	<-chan interface{}, <-chan error) {

	out := make(chan interface{}, 1)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errc)

		inb := <-in
		params, ok := inb.(*RequestParams)
		if !ok {
			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrMissingRequestParams)
			return
		}

		id := params.Params["id"]
		aID, err := strconv.Atoi(id.(string))
		if err != nil {
			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrTypeConversion)
			return
		}

		c, err := r.DeleteContact(uint(aID))

		if err != nil {
			// Handle an error that occurs during the goroutine.
			errc <- err
			return
		}

		// Send the data to the output channel but return early
		// if the context has been cancelled.
		select {
		case out <- c:
		case <-ctx.Done():
			return
		}
	}()

	return out, errc
}

// GetContact getContact by Id channel
func GetContact(ctx context.Context, in <-chan interface{}, r repository.Repository) (
	<-chan interface{}, <-chan error) {

	out := make(chan interface{}, 1)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errc)

		inb := <-in
		params, ok := inb.(*RequestParams)
		if !ok {
			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrMissingRequestParams)
			return
		}

		id := params.Params["id"]
		aID, err := strconv.Atoi(id.(string))
		if err != nil {
			// Handle an error that occurs during the goroutine.
			errc <- gerrors.New(http.StatusInternalServerError, gerrors.ErrTypeConversion)
			return
		}

		res, err := r.GetContact(uint(aID))
		if err != nil {
			// Handle an error that occurs during the goroutine.
			errc <- err
			return
		}

		// Send the data to the output channel but return early
		// if the context has been cancelled.
		select {
		case out <- res:
		case <-ctx.Done():
			return
		}
	}()

	return out, errc
}

// GetContacts getContacts channel
func GetContacts(r repository.Repository) <-chan interface{} {
	out := make(chan interface{}, 1)

	go func() {
		defer close(out)
		res := r.GetContacts()
		out <- res
	}()

	return out
}

// WaitForPipeline waits for results from all error channels.
// It returns early on the first error.
func WaitForPipeline(errs ...<-chan error) error {
	errc := MergeErrors(errs...)
	for err := range errc {
		if err != nil {
			return err
		}
	}
	return nil
}

// MergeErrors merges multiple channels of errors.
// Based on https://blog.golang.org/pipelines.
func MergeErrors(cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	// We must ensure that the output channel has the capacity to
	// hold as many errors
	// as there are error channels.
	// This will ensure that it never blocks, even
	// if WaitForPipeline returns early.
	out := make(chan error, len(cs))
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls
	// wg.Done.
	output := func(c <-chan error) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	// Start a goroutine to close out once all the output goroutines
	// are done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
