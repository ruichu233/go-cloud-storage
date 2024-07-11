package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
	JWTKey string
}

func NewJWT(jwtKey string) *JWT {
	return &JWT{
		JWTKey: jwtKey,
	}
}

type Claims struct {
	Uid int
	jwt.RegisteredClaims
}

// Award 生成token
func (j JWT) Award(uid int) (string, error) {
	// 过期时间 7 天
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(j.JWTKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// ParseToken 解析token
func (j JWT) ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.JWTKey, nil
	})
	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(Claims)
	if ok && token.Valid {
		return token, &claims, nil
	}
	return nil, nil, errors.New("the resolved claims are not utils.Claims")
}
