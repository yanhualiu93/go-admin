package models

import (
	"time"
)

type SysRole struct {
	RoleId     int       `xorm:"not null pk autoincr INT(11)"`
	RoleName   string    `xorm:"not null VARCHAR(64)"`
	RoleCode   string    `xorm:"not null unique VARCHAR(64)"`
	RoleDesc   string    `xorm:"default 'NULL' VARCHAR(255)"`
	CreateTime time.Time `xorm:"not null default 'current_timestamp()' TIMESTAMP"`
	UpdateTime time.Time `xorm:"default 'NULL' TIMESTAMP"`
	DelFlag    string    `xorm:"default ''0'' comment('删除标识（0-正常,1-删除）') CHAR(1)"`
}
