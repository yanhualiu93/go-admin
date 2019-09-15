package service

import (
	"context"
	"github.com/smallnest/rpcx/log"
	"go-admin/common/custom"
	"go-admin/verification-code/code-server/errors"
	"go-admin/verification-code/code-server/model"
)

type Code struct {
}

func (c *Code) CreateCode(ctx context.Context, req *custom.Request, res *custom.Response) error {
	log.Infof("run codeService.CreateCode req:%v", req)
	codeType, err := req.GetInt64("codeType")
	if err != nil {
		return err
	}
	w,_ := req.GetInt64("w")
	h,_ := req.GetInt64("h")
	key, base64Str := model.CreateCode(codeType,w,h)
	res.SetValue("key", key).SetValue("base64Str", base64Str)
	return nil
}
func (c *Code) CheckCode(ctx context.Context, req *custom.Request, res *custom.Response) error {
	key, err := req.GetString("key")
	if err != nil {
		return err
	}
	code, err := req.GetString("code")
	if err != nil {
		return err
	}
	if model.CheckCode(key, code) {
		res.SetValue("IsCheck", true)
	} else {
		return errors.NewCodeFailedErr()
	}
	return nil
}
