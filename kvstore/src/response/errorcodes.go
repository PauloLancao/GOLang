package response

import "net/http"

// ErrorCodes consts
const (
	ErrDuplicate             = 1
	ErrUnknown               = 2
	ErrValidation            = 3
	ErrDecoding              = 4
	ErrMissingRequestParams  = 5
	ErrTypeConversion        = 6
	ErrGeneric               = 7
	ErrArgumentException     = 8
	ErrInvalidPayload        = 9
	ErrNotFound              = 10
	ErrInvalidURL            = 11
	ErrRequestMethodNotFound = 12
	ErrMissingCommand        = 13
	ErrStopCommand           = 14
	ErrTCPNetConnection      = 15
	ErrLoggingService        = 16
)

var errorCodeText = map[int]string{
	ErrDuplicate:             "Error duplicate",
	ErrUnknown:               "Error unknown",
	ErrValidation:            "Error validation failed",
	ErrDecoding:              "Error decoding",
	ErrMissingRequestParams:  "Error missing request parameters",
	ErrTypeConversion:        "Error type conversion",
	ErrGeneric:               "Error something went wrong internally",
	ErrArgumentException:     "Error argument exception",
	ErrInvalidPayload:        "Error invalid payload",
	ErrNotFound:              "Error not found",
	ErrInvalidURL:            "Error invalid URL",
	ErrRequestMethodNotFound: "Error requet METHOD not found",
	ErrMissingCommand:        "Error missing command",
	ErrStopCommand:           "Stop Command triggered from client",
	ErrTCPNetConnection:      "Error on TCP net connection",
	ErrLoggingService:        "Error starting logging service",
}

var errorCodeHTTPStatus = map[int]int{
	ErrDuplicate:             http.StatusBadRequest,
	ErrUnknown:               http.StatusBadRequest,
	ErrNotFound:              http.StatusBadRequest,
	ErrInvalidURL:            http.StatusNotFound,
	ErrRequestMethodNotFound: http.StatusNotFound,
	ErrMissingRequestParams:  http.StatusInternalServerError,
	ErrTypeConversion:        http.StatusInternalServerError,
	ErrDecoding:              http.StatusUnprocessableEntity,
	ErrMissingCommand:        http.StatusBadRequest,
	ErrStopCommand:           http.StatusAccepted,
	ErrTCPNetConnection:      http.StatusInternalServerError,
	ErrLoggingService:        http.StatusInternalServerError,
}

// ErrorCodeText func
func ErrorCodeText(code int) string {
	return errorCodeText[code]
}

// ErrorCodeHTTPStatus func
func ErrorCodeHTTPStatus(code int) int {
	return errorCodeHTTPStatus[code]
}
