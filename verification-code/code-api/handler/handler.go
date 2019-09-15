package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"go-admin/common/custom"
	"go-admin/common/handler"
	"net/http"
	"strconv"
)

var cli client.XClient

func SetCli(c client.XClient) {
	cli = c
}
func CreateCode(ctx *gin.Context) {
	codeType := ctx.Query("codeType")
	w := ctx.Query("w")
	h := ctx.Query("h")
	req := &custom.Request{}
	if codeType != "" {
		t, err := strconv.ParseInt(codeType, 10, 0)
		if err != nil {
			req.SetParam("codeType", 1)
		} else {
			req.SetParam("codeType", t)
		}
	} else {
		req.SetParam("codeType", 1)
	}
	if w!=""&&h!="" {
		req.SetParam("w",w).SetParam("h",h)
	}
	res := &custom.Response{}
	err := cli.Call(context.Background(), "CreateCode", &req, res)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res.Body,
	})

}
func CheckCode(ctx *gin.Context) {
	req := &custom.Request{}
	res := &custom.Response{}
	fmt.Println(ctx.Query("key"))
	req.SetParam("key", ctx.PostForm("key"))
	req.SetParam("code", ctx.PostForm("code"))
	err := cli.Call(context.Background(), "CheckCode", req, res)
	if err != nil {
		handler.HandlerRes(ctx, nil, err)
		return
	}
	handler.HandlerRes(ctx, res.Body, nil)
}
