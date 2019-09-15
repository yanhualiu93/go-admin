package models

import (
	"time"
)

type SysDict struct {
	Id          int       `xorm:"not null pk autoincr comment('编号') INT(64)"`
	Value       string    `xorm:"not null comment('数据值') index VARCHAR(100)"`
	Label       string    `xorm:"not null comment('标签名') index VARCHAR(100)"`
	Type        string    `xorm:"not null comment('类型') VARCHAR(100)"`
	Description string    `xorm:"not null comment('描述') VARCHAR(100)"`
	Sort        string    `xorm:"not null comment('排序（升序）') DECIMAL(10)"`
	CreateTime  time.Time `xorm:"not null default 'current_timestamp()' comment('创建时间') TIMESTAMP"`
	UpdateTime  time.Time `xorm:"not null default 'current_timestamp()' comment('更新时间') TIMESTAMP"`
	Remarks     string    `xorm:"default 'NULL' comment('备注信息') VARCHAR(255)"`
	DelFlag     string    `xorm:"not null default ''0'' comment('删除标记') index CHAR(1)"`
}
