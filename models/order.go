package models

import "time"

type Order struct {
	OrderId       string    `json:"orderId"`
	CityName      string    `json:"cityName"`
	UserName      string    `json:"userName"`
	StartAddress  string    `json:"startAddress"`
	EndAddress    string    `json:"endAddress"`
	OrderAmount   int       `json:"orderAmount"`
	UserPayAmount int       `json:"userPayAmount"`
	PayType       int       `json:"payType"`
	VehicleName   string    `json:"vehicleName"`
	DriverName    string    `json:"driverName"`
	State         int       `json:"state"`
	EndTime       time.Time `json:"endTime"`
	CreateId      int       `json:"createId" gorm:"-"`
	CreateTime    time.Time `json:"createTime"`
}

func (Order) TableName() string {
	return "order"
}
