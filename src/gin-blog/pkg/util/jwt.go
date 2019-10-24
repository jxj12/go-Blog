package util

import (
	"fmt"
	"../setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)
type MyCustomClaims struct {
	username string `json:"username"`
	password string `json:"password"`
	jwt.StandardClaims
}

func  GenerateToken(username string ,password string)(string, error){
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := MyCustomClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),//过期时间
			Issuer:    "gin-blog", //填写对应项目
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)//获取完整的签名令牌
	fmt.Printf("%v %v", token , err)
	return token,err
}

func ParseToken(token string) (*MyCustomClaims, error){ //解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}



