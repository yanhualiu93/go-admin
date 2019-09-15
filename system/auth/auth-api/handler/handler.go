package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"go-admin/common/custom"
	"go-admin/common/errors"
	"go-admin/common/handler"
	"strings"
)

var cli client.XClient
var codeCli client.XClient
func SetXc(c client.XClient) {
	cli = c
}
func SetCodeCli(c client.XClient)  {
	codeCli = c
}
func Login(c *gin.Context) {
	var loginInfo LoginInfo
	err := c.ShouldBind(&loginInfo)
	if err!=nil {
		handler.HandlerBadRequest(c,err)
		return
	}
	res := &custom.Response{}
	Creq := &custom.Request{}
	Creq.SetParam("key",c.PostForm("key"))
	Creq.SetParam("code",c.PostForm("code"))
	err = codeCli.Call(context.Background(),"CheckCode",Creq,res)
	if err!=nil {
		handler.HandlerRes(c,nil,err)
		return
	}
	Treq := &custom.Request{}
	Treq.SetParam("username", c.PostForm("username"))
	Treq.SetParam("password", c.PostForm("password"))

	err = cli.Call(context.Background(), "Token", Treq, res)
	if err != nil {
		handler.HandlerRes(c, nil, err)
		return
	}
	handler.HandlerRes(c, res.Body, nil)
}
func CheckToken(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		handler.HandlerRes(c, nil, errors.NewUnauthorizedError())
		return
	}
	bearerToken = strings.TrimSpace(bearerToken)
	if !strings.HasPrefix(bearerToken, "Bearer") {
		handler.HandlerRes(c, nil, errors.NewUnauthorizedError())
		return
	}
	token := strings.TrimSpace(strings.Replace(bearerToken, "Bearer", "", 1))
	req := &custom.Request{}
	req.SetParam("token", token)
	res := &custom.Response{}
	err := cli.Call(context.Background(), "CheckToken", req, res)
	if err != nil {
		handler.HandlerRes(c, nil, err)
		return
	}
	handler.HandlerRes(c, res.Body, nil)
}
