package models

import (
	"time"
)

type ZipkinDependencies struct {
	Day       time.Time `xorm:"not null pk unique(day) unique(day_2) DATE"`
	Parent    string    `xorm:"not null pk unique(day) unique(day_2) VARCHAR(255)"`
	Child     string    `xorm:"not null pk unique(day) unique(day_2) VARCHAR(255)"`
	CallCount int64     `xorm:"default NULL BIGINT(20)"`
}
