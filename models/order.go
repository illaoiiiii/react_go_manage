package models

import "time"

type Order struct {
	ID            string    `gorm:"column:_id" json:"_id"`
	OrderId       string    `json:"orderId"`
	CityName      string    `json:"cityName"`
	UserName      string    `json:"userName"`
	StartAddress  string    `json:"startAddress"`
	EndAddress    string    `json:"endAddress"`
	OrderAmount   int       `json:"orderAmount"`
	UserPayAmount int       `json:"userPayAmount"`
	DriverAmount  int       `json:"driverAmount"`
	PayType       int       `json:"payType"`
	VehicleName   string    `json:"vehicleName"`
	DriverName    string    `json:"driverName"`
	State         int       `json:"state"`
	UseTime       time.Time `json:"useTime"`
	EndTime       time.Time `json:"endTime"`
	CreateId      int       `json:"createId"`
	CreateTime    time.Time `json:"createTime"`
}

func (Order) TableName() string {
	return "order_list"
}
