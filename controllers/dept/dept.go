package dept

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"react_go_manage/models"
	"react_go_manage/utils"
)

type DeptController struct{}

func (con DeptController) List(c *gin.Context) {
	deptList := []models.Dept{}
	getDepartmentList(&deptList, "")

	c.JSON(200, gin.H{
		"code": 0,
		"data": deptList,
		"msg":  "",
	})
}

func (con DeptController) Create(c *gin.Context) {
	deptName := c.PostForm("deptName")
	parentId := c.PostForm("parentId")
	userName := c.PostForm("userName")

	id := uuid.NewSHA1(uuid.Nil, []byte(deptName)).String()

	dept := models.Dept{
		Id:         id,
		DeptName:   deptName,
		UserName:   userName,
		ParentId:   parentId,
		CreateTime: utils.GetDateWithZone(),
		UpdateTime: utils.GetDateWithZone(),
	}
	err := models.DB.Create(&dept).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "增加用户失败",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "增加用户成功",
		})
	}
}

func (con DeptController) Edit(c *gin.Context) {
	deptName := c.PostForm("deptName")
	parentId := c.PostForm("parentId")
	userName := c.PostForm("userName")

	id := c.PostForm("_id")

	dept := models.Dept{
		Id:         id,
		DeptName:   deptName,
		UserName:   userName,
		ParentId:   parentId,
		UpdateTime: utils.GetDateWithZone(),
	}

	err := models.DB.Where("_id=?", id).Omit("create_time").Updates(&dept).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "修改用户失败",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "修改用户成功",
		})
	}
}

func (con DeptController) Delete(c *gin.Context) {
	id := c.PostForm("_id")
	err := models.DB.Where("_id=?", id).Delete(&models.Dept{}).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "删除用户失败",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "删除用户成功",
		})
	}
}

func getDepartmentList(deptList *[]models.Dept, id string) {
	models.DB.Where("parent_id = ?", id).Find(deptList)
	for i := range *deptList {
		getDepartmentList(&(*deptList)[i].Children, (*deptList)[i].Id)
	}
}
