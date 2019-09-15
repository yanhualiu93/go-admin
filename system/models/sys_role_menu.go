package models

type SysRoleMenu struct {
	RoleId int `xorm:"not null pk comment('角色ID') INT(11)"`
	MenuId int `xorm:"not null pk comment('菜单ID') INT(11)"`
}
