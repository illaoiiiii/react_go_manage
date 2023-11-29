package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"react_go_manage/models"
	"react_go_manage/utils"
)

type MenuController struct{}

func (con MenuController) List(c *gin.Context) {
	var permissionList []models.Permission

	menuName := c.Query("menuName")

	if menuName != "" {
		models.DB.Where("menu_name like ?", "%"+menuName+"%").Order("order_by asc").Find(&permissionList)
	} else {
		getPermissionList(&permissionList, "")
	}

	//把所有的children等于空数组变成null 这里是后端处理了，另一个地方是前端处理
	GetMenuTree(&permissionList)

	c.JSON(200, gin.H{
		"code": 0,
		"data": permissionList,
		"msg":  "",
	})
}

func (con MenuController) Create(c *gin.Context) {
	icon := c.PostForm("icon")
	menuName := c.PostForm("menuName")
	menuCode := c.PostForm("menuCode")
	menuState, _ := utils.Int(c.PostForm("menuState"))
	menuType, _ := utils.Int(c.PostForm("menuType"))
	orderBy, _ := utils.Int(c.PostForm("orderBy"))
	parentId := c.PostForm("parentId")
	path := c.PostForm("path")
	//根据NewSHA1生成uuid
	id := uuid.NewSHA1(uuid.Nil, []byte(menuName)).String()
	menu := models.Permission{
		Id:        id,
		Icon:      icon,
		MenuCode:  menuCode,
		MenuName:  menuName,
		MenuState: menuState,
		MenuType:  menuType,
		OrderBy:   orderBy,
		ParentId:  parentId,
		Path:      path,
		//上面前端传过来的数据
		CreateTime: utils.GetDateWithZone(),
		UpdateTime: utils.GetDateWithZone(),
	}
	err := models.DB.Create(&menu).Error
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

func (con MenuController) Edit(c *gin.Context) {
	id := c.PostForm("_id")
	icon := c.PostForm("icon")
	menuName := c.PostForm("menuName")
	menuCode := c.PostForm("menuCode")
	menuState, _ := utils.Int(c.PostForm("menuState"))
	menuType, _ := utils.Int(c.PostForm("menuType"))
	orderBy, _ := utils.Int(c.PostForm("orderBy"))
	parentId := c.PostForm("parentId")
	path := c.PostForm("path")

	menu := models.Permission{
		Icon:      icon,
		MenuCode:  menuCode,
		MenuName:  menuName,
		MenuState: menuState,
		MenuType:  menuType,
		OrderBy:   orderBy,
		ParentId:  parentId,
		Path:      path,
		//上面前端传过来的数据
		UpdateTime: utils.GetDateWithZone(),
	}
	err := models.DB.Where("_id = ?", id).Omit("create_time").Updates(&menu).Error
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

func (con MenuController) Delete(c *gin.Context) {
	id := c.PostForm("_id")
	err := models.DB.Where("_id = ?", id).Delete(models.Permission{}).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "删除失败",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "删除成功",
		})
	}
}

func getPermissionList(permissionList *[]models.Permission, id string) {
	models.DB.Where("parent_id = ?", id).Order("order_by asc").Find(permissionList)
	for i := range *permissionList {
		getPermissionList(&(*permissionList)[i].Children, (*permissionList)[i].Id)
	}
}

func GetMenuTree(permissionList *[]models.Permission) {
	if len(*permissionList) == 0 {
		return
	} else {
		for i := range *permissionList {
			if len((*permissionList)[i].Children) == 0 {
				(*permissionList)[i].Children = nil
			}
			GetMenuTree(&(*permissionList)[i].Children)
		}
	}
}
