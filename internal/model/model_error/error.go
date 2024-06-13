package model_error

type Error struct {
	Code  ErrorCode `json:"code"`
	Error error     `json:"detail"`
}

type ErrorIface interface {
	GetErrorCode() int
	GetErrorCodeMessage() error
	GetHttpCode() int
}

func New(code ErrorCode) *Error {
	return &Error{
		Code:  code,
		Error: ErrorMap[code],
	}
}

func (ce Error) GetErrorCode() int {
	return int(ce.Code)
}

func (ce Error) GetErrorCodeMessage() error {
	return ErrorMap[ce.Code]
}

func (ce Error) GetHttpCode() int {
	return ErrorHttpMap[ce.Code]
}
