package errors

import (
	"go-admin/common/errors"
	"net/http"
)

func NewCodeFailedErr() error {
	return &errors.Errno{
		ErrNo:   http.StatusBadRequest,
		Message: "Verification code failed",
	}
}
