package util

import (
	"fmt"
	"net/http"
)

//BadRequestError
type BadRequestError struct {
	Message string
}

func (i *BadRequestError) Error() string {
	return fmt.Sprintf("BadRequest: %v", i.Message)
}

//NotFoundError
type NotFoundError struct {
	Message string
}

func (i *NotFoundError) Error() string {
	return fmt.Sprintf("NotFound: %v", i.Message)
}

type ErrorWrapper struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Error2Wrapper(err error) (status int, errBody interface{}) {
	switch err.(type) {
	case *BadRequestError:
		return http.StatusBadRequest, ErrorWrapper{Code: "400", Message: err.Error()}
	case *NotFoundError:
		return http.StatusNotFound, ErrorWrapper{Code: "404", Message: err.Error()}
	default:
		return http.StatusInternalServerError, ErrorWrapper{Code: "500", Message: err.Error()}
	}
}
