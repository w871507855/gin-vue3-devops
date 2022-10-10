package helper

import (
	"crypto/md5"
	"fmt"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

// 生成md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

type UserClaims struct {
	UUID int64  `json:"uuid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// token加密密钥
var myKey = []byte("ops-key")

// GenerateToken 生成token
func GenerateToken(uuid int64, username string) (string, error) {
	UserClaim := &UserClaims{
		UUID:           uuid,
		Name:           username,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// AnalyseToken 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, err
	}
	return userClaim, nil
}

// GetUUID 生成uuid
func GetUUID() string {
	return uuid.NewV4().String()
}
