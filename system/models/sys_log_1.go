package models

import (
	"time"
)

type SysLog1 struct {
	Id         int64     `xorm:"pk comment('编号') BIGINT(64)"`
	Type       string    `xorm:"default ''1'' comment('日志类型') index CHAR(1)"`
	Title      string    `xorm:"default '''' comment('日志标题') VARCHAR(255)"`
	ServiceId  string    `xorm:"default 'NULL' comment('服务ID') VARCHAR(32)"`
	CreateBy   string    `xorm:"default 'NULL' comment('创建者') index VARCHAR(64)"`
	CreateTime time.Time `xorm:"default 'current_timestamp()' comment('创建时间') index DATETIME"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('更新时间') DATETIME"`
	RemoteAddr string    `xorm:"default 'NULL' comment('操作IP地址') VARCHAR(255)"`
	UserAgent  string    `xorm:"default 'NULL' comment('用户代理') VARCHAR(1000)"`
	RequestUri string    `xorm:"default 'NULL' comment('请求URI') index VARCHAR(255)"`
	Method     string    `xorm:"default 'NULL' comment('操作方式') VARCHAR(10)"`
	Params     string    `xorm:"default 'NULL' comment('操作提交的数据') TEXT"`
	Time       string    `xorm:"default 'NULL' comment('执行时间') MEDIUMTEXT"`
	DelFlag    string    `xorm:"default ''0'' comment('删除标记') CHAR(1)"`
	Exception  string    `xorm:"default 'NULL' comment('异常信息') TEXT"`
}
