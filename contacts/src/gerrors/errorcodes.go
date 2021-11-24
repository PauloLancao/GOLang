package gerrors

// ErrorCodes consts
const (
	ErrDuplicate            = 1
	ErrUnknown              = 2
	ErrValidation           = 3
	ErrDecoding             = 4
	ErrMissingRequestParams = 5
	ErrTypeConversion       = 6
	ErrGeneric              = 7
	ErrArgumentException    = 8
	ErrInvalidPayload       = 9
)

var errorCodeText = map[int]string{
	ErrDuplicate:            "Error duplicate",
	ErrUnknown:              "Error unknown",
	ErrValidation:           "Error validation failed",
	ErrDecoding:             "Error decoding",
	ErrMissingRequestParams: "Error missing request parameters",
	ErrTypeConversion:       "Error type conversion",
	ErrGeneric:              "Error something went wrong internally",
	ErrArgumentException:    "Error argument exception",
	ErrInvalidPayload:       "Error invalid payload",
}

// ErrorCodeText func
func ErrorCodeText(code int) string {
	return errorCodeText[code]
}
