package order

import (
	"github.com/gin-gonic/gin"
	"react_go_manage/models"
)

type TypeController struct{}

type Page struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

func (con TypeController) List(c *gin.Context) {
	var orderList models.Order
	if err := c.ShouldBindJSON(&orderList); err != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"err":  "获取参数失败",
		})
		return
	}
	var page Page
	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"err":  "获取分页参数失败",
		})
		return
	}
	orderId := orderList.OrderId
	driverName := orderList.DriverName
	pageNum := page.PageNum
	pageSize := page.PageSize

	query := models.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	if orderId != "" {
		query = query.Where("order_id=?", orderId)
	}
	if driverName != "" {
		query = query.Where("driver_name=?", driverName)
	}
	query.Find(&orderList)

	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"list": orderList,
			"page": gin.H{
				"pageNum":  pageNum,
				"pageSize": pageSize,
			},
		},
		"msg": "获取成功",
	})
}

func (con TypeController) VehicleList(c *gin.Context) {
	var vehicleList []models.Vehicle
	models.DB.Find(&vehicleList)
	c.JSON(200, gin.H{
		"code": 0,
		"data": vehicleList,
	})
}

func (con TypeController) CityList(c *gin.Context) {
	var cityList []models.City
	models.DB.Find(&cityList)
	c.JSON(200, gin.H{
		"code": 0,
		"data": cityList,
	})
}
