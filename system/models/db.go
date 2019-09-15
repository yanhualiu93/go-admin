package models

import "github.com/xormplus/xorm"

var engine *xorm.Engine

func SetDB(e *xorm.Engine) {
	engine = e
}
