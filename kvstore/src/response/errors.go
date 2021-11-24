package response

// APIError struct
type APIError struct {
	errorContext string
	statusCode   int
	ErrorMessage string      `json:"error"`
	ErrorCode    int         `json:"errorcode"`
	Errors       interface{} `json:"errors"`
}

// New func for custom ChannelError
func New(errorcode int, errorctx string) error {
	return NewErrors(errorctx, ErrorCodeHTTPStatus(errorcode), errorcode, struct{}{})
}

// NewErrors func for custom ChannelError
func NewErrors(errorctx string, statucode int, errorcode int, errors interface{}) error {
	return APIError{errorctx, statucode, ErrorCodeText(errorcode), errorcode, errors}
}

// Compatible with the Error interface
func (e APIError) Error() string {
	return e.ErrorMessage
}

// StatusCode exposed to error handling
func (e APIError) StatusCode() int {
	return e.statusCode
}

// ErrorContext exposed to error handling
func (e APIError) ErrorContext() string {
	return e.errorContext
}
