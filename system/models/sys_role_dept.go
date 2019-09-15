package models

type SysRoleDept struct {
	Id     int `xorm:"not null pk autoincr INT(20)"`
	RoleId int `xorm:"default NULL comment('角色ID') INT(20)"`
	DeptId int `xorm:"default NULL comment('部门ID') INT(20)"`
}
