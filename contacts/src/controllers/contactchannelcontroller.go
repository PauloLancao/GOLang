package controllers

import (
	"context"
	"net/http"

	"github.com/paulolancao/go-contacts/channels"
	"github.com/paulolancao/go-contacts/models"
	"github.com/paulolancao/go-contacts/repository"
	"github.com/paulolancao/go-contacts/utils"
)

// GetChannelContacts restapi endpoint
func GetChannelContacts(repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, <-channels.GetContacts(repo))
	}
}

// GetChannelContact restapi endpoint
// /contact/{id:[0-9]+}
func GetChannelContact(repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithCancel(context.Background())
		defer cancelFunc()

		var methodName = "GetChannelContact"
		rExt := channels.RequestExtractor{ExtractParams: true, Params: []string{"id"}}

		out, errc := channels.HTTPExtractor(ctx, &rExt, r)
		valErr := <-errc
		if valErr != nil {
			channels.HTTPResponse(w, out, valErr, methodName)
			return
		}

		out, errc = channels.GetContact(ctx, out, repo)
		valErr = <-errc
		channels.HTTPResponse(w, out, valErr, methodName)
	}
}

// CreateChannelContact restapi endpoint
func CreateChannelContact(repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithCancel(context.Background())
		defer cancelFunc()

		var methodName = "CreateChannelContact"
		var newContact = models.Contact{}
		rExt := channels.RequestExtractor{ExtractBody: true, Body: &newContact}

		out, errc := channels.HTTPExtractor(ctx, &rExt, r)
		valErr := <-errc
		if valErr != nil {
			channels.HTTPResponse(w, out, valErr, methodName)
			return
		}

		out, errc = channels.HTTPValidator(ctx, out)
		valErr = <-errc
		if valErr != nil {
			channels.HTTPResponse(w, out, valErr, methodName)
			return
		}

		out, errc = channels.CreateContact(ctx, out, repo)
		valErr = <-errc
		channels.HTTPResponse(w, out, valErr, methodName)
	}
}

// UpdateChannelContact restapi endpoint
// /contact/{id:[0-9]+}
func UpdateChannelContact(repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithCancel(context.Background())
		defer cancelFunc()

		var methodName = "UpdateChannelContact"
		var c models.Contact
		rExt := channels.RequestExtractor{ExtractBody: true, Body: &c,
			ExtractParams: true, Params: []string{"id"}}

		out, errc := channels.HTTPExtractor(ctx, &rExt, r)
		valErr := <-errc
		if valErr != nil {
			channels.HTTPResponse(w, out, valErr, methodName)
			return
		}

		out, errc = channels.UpdateContact(ctx, out, repo)
		valErr = <-errc
		channels.HTTPResponse(w, out, valErr, methodName)
	}
}

// DeleteChannelContact restapi endpoint
// /contact/{id:[0-9]+}
func DeleteChannelContact(repo repository.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithCancel(context.Background())
		defer cancelFunc()

		var methodName = "DeleteChannelContact"

		rExt := channels.RequestExtractor{ExtractParams: true, Params: []string{"id"}}

		out, errc := channels.HTTPExtractor(ctx, &rExt, r)
		valErr := <-errc
		if valErr != nil {
			channels.HTTPResponse(w, out, valErr, methodName)
			return
		}

		out, errc = channels.DeleteContact(ctx, out, repo)
		valErr = <-errc

		channels.HTTPResponse(w, out, valErr, methodName)
	}
}
