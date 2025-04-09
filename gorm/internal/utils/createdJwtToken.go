package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var SecretKey = []byte("GORM_JWT_SECRET_KEY")

func CreatedJwtToken(userId uint, Email string) (string, error) {
	// 设置jwt过期时间和内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  Email,
		"exp":    time.Now().Add(time.Hour * time.Duration(1)).Unix(), // 过期时间24小时
	})
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJwtToken(tokenString string) (map[string]interface{}, error) {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return SecretKey, nil

	})
	// 检查解析结果
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result := make(map[string]interface{})
		for key, value := range claims {
			result[key] = value
		}
		return result, nil
	}

	return nil, errors.New("invalid token")
}
