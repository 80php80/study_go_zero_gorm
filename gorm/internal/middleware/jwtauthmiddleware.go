package middleware

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm/internal/utils"
	"net/http"
	"strings"
)

type JwtAuthMiddleware struct {
	SecretKey   string
	RedisClient *redis.Client
}

func NewJwtAuthMiddleware(secretKey string, RedisClient *redis.Client) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		SecretKey:   secretKey,
		RedisClient: RedisClient,
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
		// 打印 claims 内容以便调试
		fmt.Printf("Parsed claims: %+v\n", claims)

		// 获取用户 ID 并验证类型
		userIDValue, ok := claims["userId"]
		if !ok {
			http.Error(w, "Unauthorized: Missing user ID in token", http.StatusUnauthorized)
			return
		}

		var userID string
		switch v := userIDValue.(type) {
		case string:
			userID = v
		case float64:
			userID = fmt.Sprintf("%d", int(v))
		default:
			http.Error(w, "Unauthorized: Invalid type for user ID in token", http.StatusUnauthorized)
			return
		}

		if userID == "" {
			http.Error(w, "Unauthorized: Missing or empty user ID in token", http.StatusUnauthorized)
			return
		}

		// 从 Redis 验证 Token 是否存在
		storedToken, err := m.RedisClient.Get(context.Background(), userID).Result()
		if err == redis.Nil {
			http.Error(w, "Unauthorized: Token not found in Redis", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if storedToken != tokenString {
			http.Error(w, "Unauthorized: Invalid or expired token", http.StatusUnauthorized)
			return
		}
		// 将用户信息存储到上下文中
		r = r.WithContext(context.WithValue(r.Context(), "userID", claims["userID"]))
		next(w, r)
	}
}
