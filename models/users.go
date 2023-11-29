package models

import "time"

type User struct {
	CreateId   int       `json:"createId,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	DeptId     string    `json:"deptId,omitempty"`
	DeptName   string    `json:"deptName,omitempty"`
	Job        string    `json:"job,omitempty"`
	//这里只是为了演示一下，试一下加上时区怎么存储
	LastLoginTime string `gorm:"default:null" json:"lastLoginTime,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
	Role          int    `json:"role"`
	RoleList      string `json:"roleList,omitempty"`
	State         int    `json:"state,omitempty"`
	UserEmail     string `json:"userEmail,omitempty"`
	UserId        int    `gorm:"default:null" json:"userId,omitempty"`
	UserImg       string `json:"userImg,omitempty"`
	UserName      string `json:"userName,omitempty"`
}

func (User) TableName() string {
	return "users"
}
