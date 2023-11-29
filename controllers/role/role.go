package role

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"react_go_manage/models"
	"react_go_manage/utils"
	"strings"
)

type RoleController struct{}

func (con RoleController) List(c *gin.Context) {
	roleName := c.Query("roleName")
	pageNum, _ := utils.Int(c.Query("pageNum"))
	pageSize, _ := utils.Int(c.Query("pageSize"))
	if pageNum == 0 {
		pageNum = 1
	}

	roleList := []models.Role{}
	var total int64
	query := models.DB.Preload("PermissionList")
	if roleName != "" {
		query = query.Where("role_name LIKE ?", fmt.Sprintf("%%%s%%", roleName))
	}
	countQuery := query
	countQuery.Model(&models.Role{}).Count(&total)

	query.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&roleList)

	for i, role := range roleList {
		roleList[i].PermissionList.CheckedKeysArray = strings.Split(role.PermissionList.CheckedKeys, ",")
		roleList[i].PermissionList.HalfCheckedKeysArray = strings.Split(role.PermissionList.HalfCheckedKeys, ",")
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"list": roleList,
			"page": gin.H{
				"pageNum":  pageNum,
				"pageSize": pageSize,
				"total":    total,
			},
		},
		"msg": "获取成功",
	})
}

func (con RoleController) AllList(c *gin.Context) {
	roleList := []models.Role{}
	// 查询时只包含 _id 和 role_name 两个字段
	models.DB.Select("_id,role_name").Find(&roleList)

	c.JSON(200, gin.H{
		"code": 0,
		"data": roleList,
		"msg":  "获取成功",
	})
}

func (con RoleController) UpdatePermission(c *gin.Context) {
	id := c.PostForm("_id")
	//记住这个写法不标准，但是是我的一次尝试，可以用来参考
	var checkedKeys, halfCheckedKeys []string

	// 获取所有 checkedKeys 的值
	i := 0
	for {
		key := c.PostForm(fmt.Sprintf("permissionList[checkedKeys][%d]", i))
		if key == "" {
			break
		}
		checkedKeys = append(checkedKeys, key)
		i++
	}

	// 获取所有 halfCheckedKeys 的值
	j := 0
	for {
		key := c.PostForm(fmt.Sprintf("permissionList[halfCheckedKeys][%d]", j))
		if key == "" {
			break
		}
		halfCheckedKeys = append(halfCheckedKeys, key)
		j++
	}

	fmt.Println("id:", id)
	fmt.Println("checkedKeys:", checkedKeys)
	fmt.Println("halfCheckedKeys:", halfCheckedKeys)
	//处理数据

	role := models.Role{
		ID:         id,
		UpdateTime: utils.GetDateWithZone(),
	}

	keys := models.Keys{
		ID:              id,
		CheckedKeys:     strings.Join(checkedKeys, ","),
		HalfCheckedKeys: strings.Join(halfCheckedKeys, ","),
	}
	err1 := models.DB.Updates(&keys).Error
	if err1 != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  err1,
		})
		return
	}
	err := models.DB.Updates(&role).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "修改成功",
	})
}

func (con RoleController) Edit(c *gin.Context) {
	id := c.PostForm("_id")
	roleName := c.PostForm("roleName")

	role := models.Role{
		RoleName: roleName,
	}
	err := models.DB.Where("_id=?", id).Omit("create_time").Updates(&role).Error
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

func (con RoleController) Create(c *gin.Context) {
	roleName := c.PostForm("roleName")
	if roleName == "" {
		c.JSON(400, gin.H{
			"msg": "用户名称不能为空",
		})
		return
	}
	id := uuid.NewSHA1(uuid.Nil, []byte(roleName)).String()
	role := models.Role{
		ID:         id,
		RoleName:   roleName,
		CreateTime: utils.GetDateWithZone(),
		UpdateTime: utils.GetDateWithZone(),
	}
	err := models.DB.Create(&role).Error
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

func (con RoleController) Delete(c *gin.Context) {
	id := c.PostForm("_id")
	models.DB.Where("_id=?", id).Delete(&models.Role{})
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "删除用户成功",
	})
}
