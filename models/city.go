package models

type City struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Id    int    `json:"id"`
}

func (City) TableName() string {
	return "city"
}
