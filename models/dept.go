package models

import "time"

type Dept struct {
	Id         string    `gorm:"column:_id" gorm:"''" json:"_id"`
	DeptName   string    `json:"deptName"`
	UserName   string    `json:"userName"`
	ParentId   string    `json:"parentId"`
	CreateId   int       `json:"createId"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`

	Children []Dept `gorm:"foreignKey:ParentId;references:Id" json:"children"`
}

func (Dept) TableName() string {
	return "department"
}
