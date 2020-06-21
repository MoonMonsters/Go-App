package util

import (
	"Go-App/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	ID     uint
	Mobile string
	jwt.StandardClaims
}

/**
生成token
*/
func GeneratorToken(id uint, mobile string) (string, error) {
	nowTime := time.Now()
	// 过期时间
	expireTime := nowTime.Add(setting.AppSetting.JwtExpireTime * time.Hour)

	claims := Claims{
		ID:     id,
		Mobile: mobile,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

/**
解析token
*/
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
