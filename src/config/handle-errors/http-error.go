package http_error

import "net/http"

type HttpError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (err *HttpError) Error() string {
	return err.Message
}

func NewHttpError(message, err string, code int, causes []Causes) *HttpError {
	return &HttpError{Message: message, Err: err, Code: code, Causes: causes}
}

func NewBadRequestError(message string) *HttpError {
	return &HttpError{Message: message, Err: "bad_request", Code: http.StatusBadRequest}
}

func NewBadRequestValidationError(message string, causes []Causes) *HttpError {
	return &HttpError{Message: message, Err: "bad_request", Code: http.StatusBadRequest, Causes: causes}
}

func NewInternalServerError(message string) *HttpError {
	return &HttpError{Message: message, Err: "internal_server_error", Code: http.StatusBadRequest}
}

func NewNotFoundError(message string) *HttpError {
	return &HttpError{Message: message, Err: "not_found", Code: http.StatusNotFound}
}

func NewForbiddenError(message string) *HttpError {
	return &HttpError{Message: message, Err: "forbidden", Code: http.StatusForbidden}
}

func NewUnauthorizedError(message string) *HttpError {
	return &HttpError{Message: message, Err: "unauthorized", Code: http.StatusUnauthorized}
}
