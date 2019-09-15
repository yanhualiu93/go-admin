package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-admin/common/custom"
	"go-admin/common/handler"
	"go-admin/system/models"
	handler2 "go-admin/system/system-api/handler"
	"time"
)

type DelUserRequest struct {
	UserId int `json:"userId"  binding:"required"`
	Status int `json:"status"  binding:"required"`
}
type EditUserRequest struct {
	UserId     int       `json:"userId"xorm:"not null pk autoincr comment('主键ID') INT(11)" binding:"required"`
	Username   string    `json:"username"xorm:"not null comment('用户名') unique VARCHAR(64)" binding:"required"`
	Password   string    `json:"-"xorm:"not null VARCHAR(255)" binding:"required" binding:"required"`
	Salt       string    `json:"salt"xorm:"default 'NULL' comment('随机盐') VARCHAR(255)"`
	Phone      string    `json:"phone"xorm:"not null comment('简介') unique VARCHAR(20)"`
	Avatar     string    `json:"avatar"xorm:"default 'NULL' comment('头像') VARCHAR(255)"`
	DeptId     int       `json:"deptId"xorm:"default NULL comment('部门ID') INT(11)" binding:"required"`
	UpdateTime time.Time `json:"updateTime"xorm:"default 'NULL' comment('修改时间') TIMESTAMP"`
	DelFlag    string    `json:"delFlag"xorm:"default ''0'' comment('0-正常，1-删除') CHAR(1)" `
}
func AddUser(c *gin.Context) {
	u:=models.SysUser{}
	err := c.Bind(&u)
	if err!=nil {
		handler.HandlerBadRequest(c,err)
		return
	}
	req:=&custom.Request{}
	req.SetParam("userInfo",u)
	res :=&custom.Response{}
	err = handler2.Xclient.Call(context.Background(),"AddUser",req,res)
	if err!=nil {
		handler.HandlerRes(c,nil,err)
		return
	}
	handler.HandlerRes(c,res.Body,nil)
}
func EditUser(c *gin.Context)  {
	u:=models.SysUser{}
	err :=c.Bind(&u)
	if err!=nil {
		handler.HandlerBadRequest(c,err)
		return
	}
 	req:= &custom.Request{}
 	req.SetParam("userInfo",&u)
 	res := &custom.Response{}
 	err =handler2.Xclient.Call(context.Background(),"EditUser",req,res)
	if err!=nil {
		handler.HandlerBadRequest(c,err)
		return
	}
 	handler.HandlerRes(c,res.Body,err)
}
func DelUser(c *gin.Context)  {
	var d DelUserRequest
	if err:= c.Bind(&d);err!=nil {
		handler.HandlerBadRequest(c,err)
		return
	}
	req := &custom.Request{}
	req.SetParam("userId",d.UserId)
	req.SetParam("status",d.Status)
	res := &custom.Response{}
	err := handler2.Xclient.Call(context.Background(),"DelUser",req,res)
	if err!=nil {
		handler.HandlerBadRequest(c,err)
		return
	}
	handler.HandlerRes(c,res.Body,err)
}

type QueryUserRequest struct {
	Username string `json:"username"`
	Status  int `json:"status"`
}
func GetUser(c *gin.Context)  {
	var q  QueryUserRequest
	conds := make(map[string]interface{})
	if err := c.Bind(q);err!=nil {
		 handler.HandlerBadRequest(c,err)
		return
	}
	if q.Username!="" {
		conds["username"]=q.Username
	}
	status := c.Query("status")

}