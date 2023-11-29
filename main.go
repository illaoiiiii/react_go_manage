package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"react_go_manage/models"
	"react_go_manage/routers"
)

func main() {
	r := gin.Default()

	result := []models.Login{}
	models.DB.Find(&result)
	fmt.Printf("%#v", result)

	// 添加CORS中间件  允许跨域不会被浏览器同源策略拒绝
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	//JWT

	routers.OrderRoutersInit(r)
	routers.UserRoutersInit(r)
	routers.MenuRoutersInit(r)
	routers.RolesRoutersInit(r)
	routers.DeptRoutersInit(r)
	r.Run()
}
