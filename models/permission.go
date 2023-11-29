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

	Children []Permission `gorm:"foreignKey:ParentId;references:Id" json:"children"`
	//这个字段是为了前端方便，把children复制到buttons
	//Buttons  []Permission `gorm:"foreignKey:ParentId;references:Id" json:"buttons"`
}

func (Permission) TableName() string {
	return "permission"
}
