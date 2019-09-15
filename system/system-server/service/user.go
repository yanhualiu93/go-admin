package service

import (
	"context"
	"go-admin/common/custom"
	"go-admin/common/errors"
	"go-admin/system/models"
)

type UserService struct {

}

func (s *UserService)AddUser(ctx context.Context, request custom.Request,response *custom.Response ) error {
	u,err:=request.GetParam("userInfo")
	if err!=nil {
		return  err
	}
	err = models.AddUser(u)
	if err!=nil {
		return err
	}
	response.SUCCESS()
	return nil
}
func (s *UserService)DelUser(ctx context.Context,request *custom.Request,response *custom.Response)error  {
	return nil
}
func (s *UserService)UpdateUser(ctx context.Context,request *custom.Request,response *custom.Response)error  {
	userId,err :=request.GetInt64("userId")
	if err!=nil {
		return  nil
	}
	u,err:= request.GetParam("userInfo")
	if err!=nil {
		return nil
	}
	err = models.UpdateUser(int(userId),u)
	if err!=nil {
		return err
	}
	response.SUCCESS()
	return nil
}
func (u *UserService)GetUserByUserName(ctx context.Context,request *custom.Request,response *custom.Response)error  {
	username,err := request.GetString("username")
	if err!=nil {
		return  nil
	}
	user:= models.GetUser(username)
	response.SetValue("user",user)
	response.SUCCESS()
	return nil
}

func (u *UserService)GetUserAll(ctx context.Context,request *custom.Request,response *custom.Response) error {
	m ,err:=request.GetMap("criteria")
	if err!=nil {
		return errors.NewParamRequiredError("criteria")
	}
	users:=  models.GetUserConds(m)
	response.SetValue("users",users)
	response.SUCCESS()
	return nil
}