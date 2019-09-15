package models

import (
	"time"
)

type JobExecutionLog struct {
	Id              string    `xorm:"not null pk VARCHAR(40)"`
	JobName         string    `xorm:"not null VARCHAR(100)"`
	TaskId          string    `xorm:"not null VARCHAR(255)"`
	Hostname        string    `xorm:"not null VARCHAR(255)"`
	Ip              string    `xorm:"not null VARCHAR(50)"`
	ShardingItem    int       `xorm:"not null INT(11)"`
	ExecutionSource string    `xorm:"not null VARCHAR(20)"`
	FailureCause    string    `xorm:"default 'NULL' VARCHAR(4000)"`
	IsSuccess       int       `xorm:"not null INT(11)"`
	StartTime       time.Time `xorm:"default 'NULL' TIMESTAMP"`
	CompleteTime    time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
