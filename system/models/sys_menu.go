package models

import (
	"time"
)

type SysMenu struct {
	MenuId     int       `xorm:"not null pk comment('菜单ID') INT(11)"`
	Name       string    `xorm:"not null comment('菜单名称') VARCHAR(32)"`
	Permission string    `xorm:"default 'NULL' comment('菜单权限标识') VARCHAR(32)"`
	Path       string    `xorm:"default 'NULL' comment('前端URL') VARCHAR(128)"`
	Url        string    `xorm:"default 'NULL' comment('请求链接') VARCHAR(128)"`
	Method     string    `xorm:"default 'NULL' comment('请求方法') VARCHAR(32)"`
	ParentId   int       `xorm:"default NULL comment('父菜单ID') INT(11)"`
	Icon       string    `xorm:"default 'NULL' comment('图标') VARCHAR(32)"`
	Component  string    `xorm:"default 'NULL' comment('VUE页面') VARCHAR(64)"`
	Sort       int       `xorm:"default 1 comment('排序值') INT(11)"`
	Type       string    `xorm:"default 'NULL' comment('菜单类型 （0菜单 1按钮）') CHAR(1)"`
	CreateTime time.Time `xorm:"default 'current_timestamp()' comment('创建时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') TIMESTAMP"`
	DelFlag    string    `xorm:"default ''0'' comment('0--正常 1--删除') CHAR(1)"`
}
