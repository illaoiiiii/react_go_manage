package users

import (
	"github.com/gin-gonic/gin"
	"react_go_manage/models"
	"react_go_manage/utils"
)

type LoginController struct{}

func (con LoginController) Login(c *gin.Context) {
	//TODO 这里由于前端改了所以ShouldBindJSON用不到了
	user := models.Login{}

	username := c.PostForm("userName")
	password := c.PostForm("userPwd")
	//if err := c.ShouldBindJSON(&user); err != nil {
	//	c.JSON(400, gin.H{
	//		"code": 0,
	//		"msg":  "无效的请求负载",
	//	})
	//	return
	//}
	//username := user.UserName
	//password := user.UserPwd

	models.DB.Where("user_name=? AND user_pwd=?", username, password).First(&user)

	if user.UserId <= 0 {
		c.JSON(403, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}
	token, err := utils.CreateToken(user.UserId, user.UserName)
	if err != nil {
		c.JSON(403, gin.H{
			"msg": "生成token失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": token,
		"msg":  "登陆成功",
	})
}
