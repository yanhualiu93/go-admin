package models

import (
	"time"
)

type SysUser struct {
	UserId     int       `json:"userId"xorm:"not null pk autoincr comment('主键ID') INT(11)" `
	Username   string    `json:"username"xorm:"not null comment('用户名') unique VARCHAR(64)" binding:"required"`
	Password   string    `json:"-"xorm:"not null VARCHAR(255)" binding:"required" binding:"required"`
	Salt       string    `json:"salt"xorm:"default 'NULL' comment('随机盐') VARCHAR(255)"`
	Phone      string    `json:"phone"xorm:"not null comment('简介') unique VARCHAR(20)"`
	Avatar     string    `json:"avatar"xorm:"default 'NULL' comment('头像') VARCHAR(255)"`
	DeptId     int       `json:"deptId"xorm:"default NULL comment('部门ID') INT(11)" binding:"required"`
	CreateTime time.Time `json:"createTime"xorm:"default 'current_timestamp()' comment('创建时间') TIMESTAMP"`
	UpdateTime time.Time `json:"updateTime"xorm:"default 'NULL' comment('修改时间') TIMESTAMP"`
	DelFlag    string    `json:"delFlag"xorm:"default ''0'' comment('0-正常，1-删除') CHAR(1)" `
}

func GetUser(username string) *SysUser {
	u := &SysUser{}
	r := engine.Where("username=? ", username).GetFirst(u)
	if !r.Has {
		return nil
	}
	return u
}
func AddUser(v interface{}) error {
	_, err := engine.InsertOne(v)
	return err
}
func DelUser(userId int) (err error) {
	user := &SysUser{UserId: userId}
	_, err = engine.ID(userId).Update(user)
	return
}
func UpdateUser(userId int, v interface{}) (err error) {
	_, err = engine.ID(userId).Update(&SysUser{}, v)
	return
}
func GetUserConds(m map[string]interface{})[]*SysUser {
	session := engine.NewSession()
	if username, ok := m["username"]; ok {
		 session = engine.Where("username like %?%", username)
	}
	if status, ok := m["status"]; ok {
		session = session.Where("delFag = ?", status)
	}
	var count int64
	session.Count(&count)
	var users []*SysUser
	session.Limit().Find(&users)
	return users
}
