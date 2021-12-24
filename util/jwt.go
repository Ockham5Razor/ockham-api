package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"ockham-api/config"
	"time"
)

// JwtClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
// https://www.liwenzhou.com/posts/Go/jwt_in_gin/
type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtConfig = config.GetConfig().Auth.Jwt

var tokenExpireDuration = time.Second * jwtConfig.ExpireSeconds

var jwtSecret = []byte(jwtConfig.Secret)

// GenToken 生成JWT
func GenToken(username string, jti string) (string, error) {
	now := time.Now()
	c := JwtClaims{
		username,
		jwt.StandardClaims{
			Id:        jti,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(tokenExpireDuration).Unix(), // 过期时间
			Issuer:    jwtConfig.Issuer,                    // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*JwtClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
