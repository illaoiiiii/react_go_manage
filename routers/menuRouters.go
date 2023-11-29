package routers

import (
	"github.com/gin-gonic/gin"
	"react_go_manage/controllers/menu"
)

func MenuRoutersInit(r *gin.Engine) {
	menuRouters := r.Group("/menu")
	{
		menuRouters.GET("/list", menu.MenuController{}.List)
		menuRouters.POST("/create", menu.MenuController{}.Create)
		menuRouters.POST("/edit", menu.MenuController{}.Edit)
		menuRouters.POST("/delete", menu.MenuController{}.Delete)
	}
}
