package model_error

import "net/http"

type ErrorCode int

const (
	ErrorCodeInternalServer ErrorCode = iota + 5000
)

var ErrorHttpMap = map[ErrorCode]int{
	ErrorCodeInternalServer: http.StatusInternalServerError,
}

var ErrorMap = map[ErrorCode]error{
	ErrorCodeInternalServer: errorInternalServer,
}
