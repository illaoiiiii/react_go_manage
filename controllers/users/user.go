package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"react_go_manage/models"
	"react_go_manage/utils"
	"strings"
)

type UserController struct{}

func (con UserController) List(c *gin.Context) {
	userList := []models.User{}
	pageNum, _ := utils.Int(c.Query("pageNum"))
	pageSize, _ := utils.Int(c.Query("pageSize"))
	state, _ := utils.Int(c.Query("state"))
	userID, _ := utils.Int(c.Query("userId"))
	userName := c.Query("userName") //string就不用转

	if pageNum == 0 {
		pageNum = 1
	}
	/*var total int64
	query := models.DB
	query1 := models.DB.Model(&models.User{})
	if state != 0 {
		query = query.Where("state=?", state)
		query1 = query1.Where("state=?", state)
	}
	if userID != 0 {
		query = query.Where("user_id = ?", userID)
		query1 = query1.Where("user_id = ?", userID)
	}
	if userName != "" {
		query = query.Where("user_name LIKE ?", fmt.Sprintf("%%%s%%", userName))
		query1 = query1.Where("user_name LIKE ?", fmt.Sprintf("%%%s%%", userName))
	}
	query = query.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("user_id ASC")
	query1.Count(&total)
	query.Find(&userList)*/
	var total int64
	query := models.DB
	if state != 0 {
		query = query.Where("state=?", state)
	}
	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}
	if userName != "" {
		query = query.Where("user_name LIKE ?", fmt.Sprintf("%%%s%%", userName))
	}
	//之前的分页有问题，我们这的需求是分页之前统计总数
	countQuery := query
	countQuery.Model(&models.User{}).Count(&total)

	query.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("user_id ASC").Find(&userList)

	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"list": userList,
			"page": gin.H{
				"pageNum":  pageNum,
				"pageSize": pageSize,
				"total":    total,
			},
		},
		"msg": "获取成功",
	})
}

// 重要注释 这里的c.PostForm是获取前端传过来的post请求的form-data数据
// 但是前端传过来的是json数据，所以这里获取不到，所以要用c.BindJSON

func (con UserController) Create(c *gin.Context) {
	userName := c.PostForm("userName")
	if userName == "" {
		c.JSON(400, gin.H{
			"msg": "用户名称不能为空",
		})
		return
	}
	deptId := strings.Trim(c.PostForm("deptId"), " ")
	job := strings.Trim(c.PostForm("job"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	roleList := strings.Trim(c.PostForm("roleList"), "")
	state, _ := utils.Int(c.PostForm("state"))
	userEmail := strings.Trim(c.PostForm("userEmail"), "")
	userImg := strings.Trim(c.PostForm("userImg"), "")

	//fmt.Println(utils.GetDate())
	user := models.User{
		DeptId:    deptId,
		Job:       job,
		Mobile:    mobile,
		RoleList:  roleList,
		State:     state,
		UserEmail: userEmail,
		UserImg:   userImg,
		UserName:  userName,
		//上面前端传过来的数据
		CreateTime: utils.GetDateWithZone(),
		//只是创建还没有登录就不用给值
		//LastLoginTime: utils.GetDate(),
	}
	err := models.DB.Create(&user).Error
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

func (con UserController) Edit(c *gin.Context) {
	deptId := strings.Trim(c.PostForm("deptId"), " ")
	job := strings.Trim(c.PostForm("job"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	roleList := strings.Trim(c.PostForm("roleList"), "")
	state, _ := utils.Int(c.PostForm("state"))
	userEmail := strings.Trim(c.PostForm("userEmail"), "")
	userImg := strings.Trim(c.PostForm("userImg"), "")
	userName := strings.Trim(c.PostForm("userName"), "")
	//这个c.PostForm好像默认拿过来的就是string
	userId := c.PostForm("userId")

	//fmt.Println(utils.GetDate())
	user := models.User{
		DeptId:    deptId,
		Job:       job,
		Mobile:    mobile,
		RoleList:  roleList,
		State:     state,
		UserEmail: userEmail,
		UserImg:   userImg,
		UserName:  userName,
	}
	err := models.DB.Where("user_id=?", userId).Omit("create_time").Updates(&user).Error
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

// Delete 单个删除和批量删除
func (con UserController) Delete(c *gin.Context) {
	//这是前端用body里的form-data传过来的
	//userIds := c.PostForm("userIds")
	//result := utils.StringToArray(userIds)
	//这是post请求的data的数据，同时headers是content-type: application/json
	result := c.PostFormArray("userIds[]")
	fmt.Println(result)

	err1 := models.DB.Where("user_id IN (?)", result).Delete(&models.User{}).Error
	if err1 != nil {
		c.JSON(400, gin.H{
			"code": 0,
			"msg":  "删除用户失败",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "删除用户成功",
		})
	}
}

func (con UserController) UserInfo(c *gin.Context) {
	//这里不能用id只能用jwt来判断身份
	//userId := c.Query("userId")
	authorization := c.Request.Header.Get("Authorization")
	jwtWithoutBearer := strings.TrimPrefix(authorization, "Bearer ")
	userStruct, err := utils.ParseToken(jwtWithoutBearer)
	fmt.Println(userStruct)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "获取用户信息失败",
		})
		return
	}

	fmt.Println(userStruct.Id)
	user := models.User{}
	models.DB.Where("user_id=?", userStruct.Id).First(&user)
	c.JSON(200, gin.H{
		"code": 0,
		"data": user,
		"msg":  "",
	})
}

func (con UserController) AllList(c *gin.Context) {
	userList := []models.User{}
	err := models.DB.Select("user_id,user_email,user_name").Find(&userList).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"msg": "获取用户信息失败",
		})
		return
	}

	//处理role=0的情况，因为role=0的是超级管理员，这里不需要，但是别的地方需要

	c.JSON(200, gin.H{
		"code": 0,
		"data": userList,
		"msg":  "",
	})
}
