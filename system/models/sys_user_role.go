package models

type SysUserRole struct {
	UserId int `xorm:"not null pk comment('用户ID') INT(11)"`
	RoleId int `xorm:"not null pk comment('角色ID') INT(11)"`
}
