package gerrors

// Channelerror exported error handling
// Update to support ErrorCodes includes, code + message of error
// more generic way to tell the client what went wrong
type Channelerror interface {
	StatusCode() int
	Error() string
	ErrorCode() int
	Errors() interface{}
}

type channelerror struct {
	statusCode   int
	errorMessage string
	errorCode    int
	errors       interface{}
}

// New func for custom ChannelError
func New(statucode int, errorcode int) error {
	return NewErrors(statucode, errorcode, struct{}{})
}

// NewErrors func for custom ChannelError
func NewErrors(statucode int, errorcode int, errors interface{}) error {
	return &channelerror{statucode, ErrorCodeText(errorcode), errorcode, errors}
}

// Create a function Error() string and associate it to the struct.
func (e *channelerror) Error() string {
	return e.errorMessage
}

// Code exposed to error handling
func (e *channelerror) StatusCode() int {
	return e.statusCode
}

// ErrorCode exposed to error handling
func (e *channelerror) ErrorCode() int {
	return e.errorCode
}

// Errors exposed to error handling
func (e *channelerror) Errors() interface{} {
	return e.errors
}
