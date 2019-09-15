package models

import (
	"time"
)

type SysZuulRoute struct {
	Id                   int       `xorm:"not null pk autoincr comment('router Id') INT(11)"`
	Path                 string    `xorm:"not null comment('路由路径') VARCHAR(255)"`
	ServiceId            string    `xorm:"not null comment('服务名称') VARCHAR(255)"`
	Url                  string    `xorm:"default 'NULL' comment('url代理') VARCHAR(255)"`
	StripPrefix          string    `xorm:"default ''1'' comment('转发去掉前缀') CHAR(1)"`
	Retryable            string    `xorm:"default ''1'' comment('是否重试') CHAR(1)"`
	Enabled              string    `xorm:"default ''1'' comment('是否启用') CHAR(1)"`
	SensitiveheadersList string    `xorm:"default 'NULL' comment('敏感请求头') VARCHAR(255)"`
	CreateTime           time.Time `xorm:"default 'current_timestamp()' comment('创建时间') TIMESTAMP"`
	UpdateTime           time.Time `xorm:"default 'NULL' comment('更新时间') TIMESTAMP"`
	DelFlag              string    `xorm:"default ''0'' comment('删除标识（0-正常,1-删除）') CHAR(1)"`
}
