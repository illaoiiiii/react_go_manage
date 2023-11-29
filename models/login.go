package models

type Login struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

func (Login) TableName() string {
	return "users"
}
