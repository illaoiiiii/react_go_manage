package routers

import (
	"github.com/gin-gonic/gin"
	"react_go_manage/controllers/users"
)

func UserRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/users")
	{
		userRouters.GET("/getPermissionList", users.PermissionController{}.List)

		userRouters.GET("/list", users.UserController{}.List)
		userRouters.GET("/all/list", users.UserController{}.AllList)
		userRouters.POST("/create", users.UserController{}.Create)
		userRouters.POST("/edit", users.UserController{}.Edit)
		//总结起来，DELETE请求通常不支持在请求主体中传递数据，但你可以考虑使用POST或PUT请求来替代，并在请求主体中传递数组或其他数据。
		userRouters.POST("/delete", users.UserController{}.Delete)
		userRouters.GET("/getUserInfo", users.UserController{}.UserInfo)

		userRouters.POST("/login", users.LoginController{}.Login)
	}

}
