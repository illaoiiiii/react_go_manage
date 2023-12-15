package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"react_go_manage/models"
	"react_go_manage/routers"
)

func main() {
	r := gin.Default()

	result := []models.Login{}
	models.DB.Find(&result)
	fmt.Printf("%#v", result)

	//处理OPTIONS请求
	r.OPTIONS("/", func(c *gin.Context) {
		// 设置 CORS 头部信息
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		// 返回成功状态码
		c.Status(http.StatusOK)
	})
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
