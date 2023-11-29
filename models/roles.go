package models

import "time"

type Role struct {
	ID       string `gorm:"column:_id;primaryKey" json:"_id"`
	RoleName string `gorm:"column:role_name" json:"roleName"`

	//下面这俩是为了前端展示用的,gorm查数据的时候不要查询这个字段
	PermissionList Keys `gorm:"foreignKey:ID;references:ID" json:"permissionList"`

	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type Keys struct {
	ID              string `gorm:"column:_id;primaryKey" json:"-"`
	CheckedKeys     string `gorm:"column:checked_keys" json:"-"`
	HalfCheckedKeys string `gorm:"column:half_checked_keys" json:"-"`
	//下面这俩是为了前端展示用的,gorm查数据的时候不要查询这俩字段
	CheckedKeysArray     []string `gorm:"-" json:"checkedKeys"`
	HalfCheckedKeysArray []string `gorm:"-" json:"halfCheckedKeys"`
}

func (Role) TableName() string {
	return "roles"
}

func (Keys) TableName() string {
	return "roles"
}
