package errors

import (
	"go-admin/common/errors"
	"net/http"
)

func NewTokenError() error {
	return &errors.Errno{
		ErrNo:   http.StatusUnauthorized,
		Message: "token validation failed",
	}
}
func NewCreateTokenError() error {
	return &errors.Errno{
		ErrNo:   http.StatusInternalServerError,
		Message: "token creation failed",
	}
}
func NewUserNoFoundError() error {
	return &errors.Errno{
		ErrNo:   http.StatusBadRequest,
		Message: "user no found",
	}
}
func NewPasswordError() error {
	return &errors.Errno{
		ErrNo:   http.StatusBadRequest,
		Message: "Password authentication failed",
	}
}
