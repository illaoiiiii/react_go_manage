package routers

import (
	"github.com/gin-gonic/gin"
	"react_go_manage/controllers/order"
)

func OrderRoutersInit(r *gin.Engine) {
	orderRouter := r.Group("/order")
	{
		orderRouter.GET("/list", order.TypeController{}.List)
		orderRouter.GET("/vehicleList", order.TypeController{}.VehicleList)
		orderRouter.GET("/cityList", order.TypeController{}.CityList)
		orderRouter.POST("/orderExport", order.TypeController{}.OrderExport)
		orderRouter.POST("/orderImport", order.TypeController{}.OrderImport)
		orderRouter.POST("/delete", order.TypeController{}.Delete)
	}
}
