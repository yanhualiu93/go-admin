package models

import (
	"time"
)

type SysDept struct {
	DeptId     int       `xorm:"not null pk autoincr INT(20)"`
	Name       string    `xorm:"default 'NULL' comment('部门名称') VARCHAR(50)"`
	OrderNum   int       `xorm:"default NULL comment('排序') INT(11)"`
	CreateTime time.Time `xorm:"default 'current_timestamp()' comment('创建时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"default 'NULL' comment('修改时间') TIMESTAMP"`
	DelFlag    string    `xorm:"default ''0'' comment('是否删除  -1：已删除  0：正常') CHAR(1)"`
	ParentId   int       `xorm:"default NULL INT(11)"`
}

func GetAllDept()  (*[]*SysDept, error)  {
	var depts []*SysDept
	err:= engine.Find(&depts)
	if(err!=nil){
		return  nil,err
	}
	return  &depts,nil
}
func AddDept(dept *SysDept) error {
	_,err:=engine.InsertOne(dept)
	return  err
}
func EditDept(dept *SysDept)  {

}
func GetDeptOne() *SysDept  {
	
}