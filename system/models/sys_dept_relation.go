package models

type SysDeptRelation struct {
	Ancestor   int `xorm:"not null pk comment('祖先节点') index INT(11)"`
	Descendant int `xorm:"not null pk comment('后代节点') index INT(11)"`
}
