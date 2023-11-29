package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	jwt.StandardClaims // 标准Claims结构体，可设置8个标准字段
	Id                 int
	UserName           string
}

func CreateToken(id int, username string) (string, error) {
	// 创建MyClaims结构体实例
	claims := &MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为一天后
		},
		Id:       id,
		UserName: username,
	}

	// 创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成token字符串
	signedToken, err := token.SignedString([]byte("your-secret-key")) // 使用你自己的密钥
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil // 使用你自己的密钥
	})
	if err != nil {
		return nil, err
	}

	// 获取token中的claims
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
