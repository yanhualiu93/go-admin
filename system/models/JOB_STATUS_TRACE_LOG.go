package models

import (
	"time"
)

type JobStatusTraceLog struct {
	Id             string    `xorm:"not null pk VARCHAR(40)"`
	JobName        string    `xorm:"not null VARCHAR(100)"`
	OriginalTaskId string    `xorm:"not null VARCHAR(255)"`
	TaskId         string    `xorm:"not null index(TASK_ID_STATE_INDEX) VARCHAR(255)"`
	SlaveId        string    `xorm:"not null VARCHAR(50)"`
	Source         string    `xorm:"not null VARCHAR(50)"`
	ExecutionType  string    `xorm:"not null VARCHAR(20)"`
	ShardingItem   string    `xorm:"not null VARCHAR(100)"`
	State          string    `xorm:"not null index(TASK_ID_STATE_INDEX) VARCHAR(20)"`
	Message        string    `xorm:"default 'NULL' VARCHAR(4000)"`
	CreationTime   time.Time `xorm:"default 'NULL' TIMESTAMP"`
}
