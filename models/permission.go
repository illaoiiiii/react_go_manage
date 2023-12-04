package models

import "time"

type Permission struct {
	CreateId   int       `json:"create_id"`
	CreateTime time.Time `gorm:"default:null" json:"createTime"`
	Icon       string    `json:"icon"`
	MenuCode   string    `json:"menuCode"`
	MenuName   string    `json:"menuName"`
	MenuState  int       `json:"menuState"`
	MenuType   int       `json:"menuType"`
	OrderBy    int       `json:"orderBy"`
	ParentId   string    `json:"parentId"`
	Path       string    `json:"path"`
	UpdateTime time.Time `json:"updateTime"`
	Id         string    `gorm:"column:_id" gorm:"''" json:"_id"`
	//递归便利字段
	Children []Permission `gorm:"foreignKey:ParentId;references:Id" json:"children"`
}

func (Permission) TableName() string {
	return "permission"
}
