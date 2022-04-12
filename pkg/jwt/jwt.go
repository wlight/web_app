package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var mySigningKey = []byte("密码加盐")

var keyFunc jwt.Keyfunc

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

//
//func GenToken(userId int64, username string) (string, error) {
//	c := MyClaims{
//		Username: username,
//		UserId:   userId,
//		StandardClaims: jwt.StandardClaims{
//			Issuer:    "web_app",
//			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
//	return token.SignedString(mySigningKey)
//}

// GenToken 生成access token和refresh token
func GenToken(userId int64, username string) (aToken string, rToken string, err error) {
	c := MyClaims{
		Username: username,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "web_app",
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
		},
	}

	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySigningKey)

	// refresh token 不需要存任何自定义数据，只需要过期时间长就可以了

	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30 * 3600).Unix(),
		Issuer:    "app_web",
	}).SignedString(mySigningKey)

	return
}

// ParseToken 解析token
func ParseToken(tokenString string) (mc *MyClaims, err error) {
	mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return mc, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token 无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}

	// 从旧access token 中解析出claims 数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当 access token 是过期错误，并且refresh token 没有过期就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserId, claims.Username)
	}

	return
}
