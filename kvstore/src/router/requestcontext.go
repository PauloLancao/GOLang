package router

import (
	"bytes"
	"context"
	"logging"
	"net/http"
	"net/url"
	"strings"
)

// ParamKey const
const ParamKey string = "key"

// ParamBody const
const ParamBody string = "body"

// Context create one per request using encapsulation to pass request scoped entities
type Context struct {
	uID       string
	W         http.ResponseWriter
	R         *http.Request
	Ctx       context.Context
	CtxCancel context.CancelFunc
	Builder   Builder
}

// Builder struct used to extract all data from http request
type Builder struct {
	ExtractBody   bool
	ExtractParams bool
	Body          interface{}
	Params        map[string]interface{}
}

// NewContext per request func
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return new(Context).setResponse(w).setRequest(r).newContextWithCancel().newUUID()
}

// NewContextWithCancel func
func (c *Context) newContextWithCancel() *Context {
	c.Ctx, c.CtxCancel = context.WithCancel(context.Background())
	return c
}

// WithBuilder func
func (c *Context) WithBuilder(b Builder) *Context {
	c.Builder = b
	return c
}

func (c *Context) setResponse(w http.ResponseWriter) *Context {
	c.W = w
	return c
}

func (c *Context) setRequest(r *http.Request) *Context {
	c.R = r
	return c
}

func (c *Context) newUUID() *Context {
	c.uID = logging.UUID()
	return c
}

// HasPath aux func
// This funcs won't work for more complex paths
// Needs some care here
func (c *Context) HasPath() bool {
	split := splitURLPath(c.R.URL.Path)
	// path e.g. /kvs/:key
	return len(split) == 3
}

// PathVars func
func (c *Context) PathVars() map[string]string {
	split := splitURLPath(c.R.URL.Path)
	if len(split) == 3 {
		key := split[2]
		vars := make(map[string]string)
		vars[ParamKey] = key
		return vars
	}

	return nil
}

func splitURLPath(path string) []string {
	return strings.Split(path, "/")
}

// Cancelled call for cancel go rotines
func (c *Context) Cancelled() bool {
	select {
	case <-c.Ctx.Done():
		return true
	default:
		return false
	}
}

// Headers aux func
func (c *Context) Headers() map[string][]string {
	return c.R.Header
}

// GetContentType aux func
func (c *Context) GetContentType() string {

	contentType := c.R.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return contentType
}

// QueryString aux func
func (c *Context) QueryString() url.Values {
	return c.R.URL.Query()
}

// Body aux func
func (c *Context) Body() ([]byte, error) {

	buf := bytes.NewBuffer(make([]byte, 0, c.R.ContentLength))
	n, err := buf.ReadFrom(c.R.Body)
	defer c.R.Body.Close()

	if err != nil {
		logging.Msgf(
			c.uID,
			"RequestContext",
			"Body",
			"Error reading request: %+v bytes: %d", err, n)
		return nil, err
	}

	return buf.Bytes(), nil
}
