package middleware

import (
	"github.com/gin-gonic/gin"
	"react_go_manage/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "token 解析错误",
			})
			c.Abort()
			return
		}

		// 将用户信息传递给后续处理函数
		c.Set("token", tokenStr)
		c.Set("id", claims.Id)
		c.Set("userName", claims.UserName)

		c.Next()

	}
}
