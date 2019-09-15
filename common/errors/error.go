package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Errno struct {
	ErrNo   int    `json:"err_no"`
	Message string `json:"message"`
}

func (p *Errno) Error() string {
	str, _ := json.Marshal(p)
	return string(str)
}
func New(code int,message string) error {
	return  &Errno{ErrNo:code,Message:message}
}
func NewParamRequiredError(paramKey string) error {
	return &Errno{
		ErrNo:   http.StatusBadRequest,
		Message: fmt.Sprintf("%s is required", paramKey),
	}
}
func NewParamNonstandardError(paramKey string) error {
	return &Errno{
		ErrNo:   http.StatusBadRequest,
		Message: fmt.Sprintf("param %s do not meet the requirements", paramKey),
	}
}
func NewUnauthorizedError() error {
	return &Errno{
		ErrNo:   http.StatusUnauthorized,
		Message: "Authorization authentication failed",
	}
}
