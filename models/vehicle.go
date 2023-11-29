package models

type Vehicle struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sum  int    `json:"sum"`
}

func (Vehicle) TableName() string {
	return "vehicle"
}
