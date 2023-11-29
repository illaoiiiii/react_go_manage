package routers

import (
	"github.com/gin-gonic/gin"
	"react_go_manage/controllers/role"
)

func RolesRoutersInit(r *gin.Engine) {
	rolesRouters := r.Group("/roles")
	{
		rolesRouters.GET("/list", role.RoleController{}.List)
		rolesRouters.GET("/allList", role.RoleController{}.AllList)
		rolesRouters.POST("/create", role.RoleController{}.Create)
		rolesRouters.POST("/edit", role.RoleController{}.Edit)
		rolesRouters.POST("/update/permission", role.RoleController{}.UpdatePermission)
		rolesRouters.POST("/delete", role.RoleController{}.Delete)
	}
}
