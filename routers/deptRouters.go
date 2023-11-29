package routers

import (
	"github.com/gin-gonic/gin"
	"react_go_manage/controllers/dept"
)

func DeptRoutersInit(r *gin.Engine) {
	deptRouters := r.Group("/dept")
	{
		deptRouters.GET("/list", dept.DeptController{}.List)
		deptRouters.POST("/create", dept.DeptController{}.Create)
		deptRouters.POST("/edit", dept.DeptController{}.Edit)
		deptRouters.POST("/delete", dept.DeptController{}.Delete)
	}
}
