package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"react_go_manage/models"
	"react_go_manage/utils"
	"strings"
)

type PermissionController struct{}

func (con PermissionController) List(c *gin.Context) {
	//通过jwt判断角色来返回对应的权限列表
	authorization := c.Request.Header.Get("Authorization")
	jwtWithoutBearer := strings.TrimPrefix(authorization, "Bearer ")
	userStruct, err := utils.ParseToken(jwtWithoutBearer)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 50001,
			"msg":  "获取用户权限失败",
		})
		return
	}

	//获取用户的Id在查询他的系统角色，最后获取权限列表
	fmt.Println(userStruct.Id)
	role := models.User{}
	models.DB.Where("user_id = ?", userStruct.Id).First(&role)
	fmt.Println(role.RoleList)
	permission := models.Keys{}
	models.DB.Where("_id = ?", role.RoleList).First(&permission)
	fmt.Println(permission)

	//查询到权限对应的两个数组了，下面要做的就是把他们转成数组，然后再递归的时候判断id是不是符合的
	permission.CheckedKeysArray = strings.Split(permission.CheckedKeys, ",")
	permission.HalfCheckedKeysArray = strings.Split(permission.HalfCheckedKeys, ",")
	//合并这两个数组为一个数组
	AllKeysArray := append(permission.CheckedKeysArray, permission.HalfCheckedKeysArray...)

	//模拟递归
	permissionList := []models.Permission{}
	buttonList := []string{}
	getPermissionList(&permissionList, "", AllKeysArray, &buttonList)

	//标准递归 我parentId存的是string类型的id
	//testList := []models.Permission{}
	//getPermissionList2(&testList, "")
	//用一下原生sql
	//models.DB.Raw("select menu_code from react_go_manage.permission where menu_code is not null").Scan(&buttonList)

	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"buttonList": buttonList,
			"menuList":   permissionList,
			//"testList":   testList,
		},
		"msg": "",
	})
}

// 老版本没有带过滤的
//func getPermissionList2(permissionList *[]models.Permission, id string) {
//	models.DB.Where("parent_id = ?", id).Order("order_by asc").Find(permissionList)
//	for i := range *permissionList {
//		getPermissionList2(&(*permissionList)[i].Children, (*permissionList)[i].Id)
//	}
//}

func getPermissionList(permissionList *[]models.Permission, id string, AllKeysArray []string, buttonList *[]string) {
	//判断id是不是在checkedKeysArray或者HalfCheckedKeysArray里面
	models.DB.Where("parent_id = ? AND _id IN (?)", id, AllKeysArray).Order("order_by asc").Find(permissionList)
	for i := range *permissionList {
		if (*permissionList)[i].MenuCode != "" {
			*buttonList = append(*buttonList, (*permissionList)[i].MenuCode)
		}
		getPermissionList(&(*permissionList)[i].Children, (*permissionList)[i].Id, AllKeysArray, buttonList)
	}
}
