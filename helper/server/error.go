package server

import (
	"fmt"
	"net/http"
)

type Err struct {
	err        string
	Message    string
	StatusCode int
}

var (
	ErrorBadRequest     = Err{err: "bad request", StatusCode: http.StatusBadRequest}
	ErrorUnauthorized   = Err{err: "unauthorized", StatusCode: http.StatusUnauthorized}
	ErrorForbidden      = Err{err: "forbidden", StatusCode: http.StatusForbidden}
	ErrorInternalDB     = Err{err: "internal db", StatusCode: http.StatusInternalServerError}
	ErrorInternalDomain = Err{err: "internal domain", StatusCode: http.StatusInternalServerError}
)

func (err Err) Error() string {
	return fmt.Sprintf("%s - %s", err.err, err.Message)
}
