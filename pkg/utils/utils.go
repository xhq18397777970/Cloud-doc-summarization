package utils

//将用户的信息打包成token，供给其他需要身份验证的功能使用
import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// json文件密钥
var JWTsecret = []byte("ABAB")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	//claims不应该带password，不安全
	// Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, password string) (string, error) {

	notTime := time.Now()
	expireTime := notTime.Add(24 * time.Hour)

	////带时间戳的token
	claims := Claims{
		Id:       id,
		UserName: username,
		// Password:password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo_list",
		},
	}
	//结构体转token类型
	tokenCliams := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//token转字符串，并使用事先声明的密钥加密
	token, err := tokenCliams.SignedString(JWTsecret)

	return token, err

}

// parseToken验证用户token
func ParseToken(token string) (*Claims, error) {
	//传入token，解析为claim结构体类型
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
