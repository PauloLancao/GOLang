package router

import (
	"fmt"
	"logging"
	"net/http"
	"net/url"
	"regexp"
	"response"
	"storage"
	"strings"
)

var acceptableEndpoints = url.Values{"/kvs": []string{"^/kvs$"}, "/kvs/": []string{"^/kvs/[a-zA-Z0-9_]+$"}}

// IsAcceptableURLPath func
func isAcceptableURLPath(path string) bool {
	var validURL bool
	for k := range acceptableEndpoints {
		if strings.HasPrefix(path, k) {
			v := acceptableEndpoints.Get(k)
			match, err := regexp.MatchString(v, path)
			if err != nil {
				logging.Msgf(logging.UUID(), "Endpoints", "IsAcceptableURLPath", "regexp matchstring failed %+v", err)
				return false
			}
			validURL = validURL || match
		}
	}

	return validURL
}

func httpHandler(s storage.Store) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		errCtx := "Router::HttpHandler"
		rc := NewContext(w, r)

		if !isAcceptableURLPath(r.URL.Path) {
			ErrorResponse(rc, http.StatusNotFound, response.New(response.ErrInvalidURL, errCtx))
		} else {
			switch r.Method {
			case "GET":
				if rc.HasPath() {
					Handler(rc, s, StoreExecutor(rc, CmdGet))
				}
			case "POST":
				Handler(rc, s, StoreExecutor(rc, CmdCreate))
			case "PUT":
				Handler(rc, s, StoreExecutor(rc, CmdUpdate))
			case "DELETE":
				Handler(rc, s, StoreExecutor(rc, CmdDelete))
			default:
				ErrorResponse(rc, http.StatusBadRequest, response.New(response.ErrNotFound, errCtx))
			}
		}
	}
}

const port = "9000"
const domain = "localhost"

// Start webserver
func Start(store storage.Store, done chan bool) {
	uuid := logging.UUID()
	logging.Msgf(uuid, "Start", "Router", "server domain: %s - port: %s", domain, port)

	http.HandleFunc("/kvs", httpHandler(store))
	http.HandleFunc("/kvs/", httpHandler(store))

	logging.Msg(uuid, "Start", "Router", "HTTP server started successfully")
	logging.Fatalf(uuid,
		"Start",
		"ListenAndServe",
		"Error:: %+v", http.ListenAndServe(fmt.Sprintf("%s:%s", domain, port), nil))

	done <- true
}
