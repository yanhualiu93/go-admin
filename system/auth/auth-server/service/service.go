package service

import (
	"context"
	"github.com/smallnest/rpcx/log"
	"go-admin/common/custom"
	"go-admin/system/auth/auth-server/errors"
	"go-admin/system/auth/auth-server/token"
	"go-admin/system/auth/auth-server/utils"
	"go-admin/system/models"
)

type TokenService struct {
}

func (t *TokenService) Token(ctx context.Context, request *custom.Request, response *custom.Response) error {
	log.Infof("run token : %v ", request)
	password, err := request.GetString("password")
	if err != nil {
		return err
	}
	username, err := request.GetString("username")
	if err != nil {
		return err
	}
	u := models.GetUser(username)
	if u == nil {
		return errors.NewUserNoFoundError()
	}
	if utils.Compare(u.Password, password) != nil {
		return errors.NewPasswordError()
	}
	c := &token.Content{
		Id:       u.UserId,
		Username: u.Username,
	}
	tokenStr, err := token.Sign(c)
	if err != nil {
		return errors.NewCreateTokenError()
	}
	response.SetValue("token", tokenStr)
	log.Infof("token success : %v", response)
	return nil
}
func (t *TokenService) CheckToken(ctx context.Context, request custom.Request, response *custom.Response) error {
	log.Infof("run CheckToken:%v", request)
	tokenStr, err := request.GetString("token")
	if err != nil {
		return err
	}
	c, err := token.VerifyToken(tokenStr)
	if err != nil {
		return errors.NewTokenError()
	}
	response.SetValue("id", c["id"])
	response.SetValue("username", c["username"])
	log.Infof("CheckToken success:%v", response)
	return nil
}
