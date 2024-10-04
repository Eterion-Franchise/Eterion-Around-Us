package errors

import "fmt"

type InternalError struct {
	Code string
	Msg  string
}

func (e InternalError) Error() string {
	return fmt.Sprintf("[Error code: %s] %s", e.Code, e.Msg)
}

var INVALID_FORM_STRING_TYPE = InternalError{
	Code: "INVALID_FORM_STRING_TYPE",
	Msg:  "Invalid data type recieved. Only types.WikiDataType is supported",
}
