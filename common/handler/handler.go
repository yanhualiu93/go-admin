package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-admin/common/errors"
	"net/http"
	"strings"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func handlerErr(err error) (int, string) {
	if err == nil {
		return http.StatusOK, "ok"
	}
	errno := &errors.Errno{}
	e := json.Unmarshal([]byte(err.Error()), &errno)
	if e != nil {
		return http.StatusInternalServerError, "internal server error"
	}
	errNo := errno.ErrNo
	msg := errno.Message
	switch errNo {
	case http.StatusBadRequest:
		return http.StatusBadRequest, msg
	case http.StatusUnauthorized:
		return http.StatusUnauthorized, msg
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
func HandlerRes(c *gin.Context, data interface{}, err error) {
	code, msg := handlerErr(err)
	c.JSON(code, &Response{
		Message: msg,
		Data:    data,
	})
}
func HandlerBadRequest(c *gin.Context,err error)  {
	if err!=nil {
		errString :=err.Error()
		c.JSON(http.StatusBadRequest,&Response{
			Message: "bad request",
			Data:strings.Split(errString,"\n"),
		})
	}
}

