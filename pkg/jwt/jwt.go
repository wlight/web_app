package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var mySigningKey = []byte("密码加盐")

const TokenExpireDuration = time.Hour * 2

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

type MyClaims struct {
	Username string `json:"username"`
	UserId   int64  `json:"user_id"`
	jwt.StandardClaims
}

func GenToken(userId int64, username string) (string, error) {
	c := MyClaims{
		Username: username,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "web_app",
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySigningKey)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	c := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return c, nil
	}

	return nil, errors.New("Invalid token")
}
