package middleware

import (
	"context"
	"fmt"
	"gorm/internal/utils"
	"net/http"
	"strings"
)

type JwtAuthMiddleware struct {
	SecretKey string
}

func NewJwtAuthMiddleware(secretKey string) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		SecretKey: secretKey,
	}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/admin/login") {
			next(w, r)
			return
		}
		// TODO generate middleware implement function, delete after code implementation
		// 检查 Authorization 头部
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}
		// 解析 Bearer Token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}
		// 使用 ParseJwtToken 验证 Token
		claims, err := utils.ParseJwtToken(tokenString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unauthorized: %v", err), http.StatusUnauthorized)
			return
		}
		// 将用户信息存储到上下文中
		r = r.WithContext(context.WithValue(r.Context(), "userID", claims["userID"]))

		next(w, r)
	}
}
