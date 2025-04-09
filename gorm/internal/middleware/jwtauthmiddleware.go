package middleware

import (
	"fmt"
	"net/http"
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
		// TODO generate middleware implement function, delete after code implementation

		//我输出一个资环
		fmt.Println("lai l ma")

		// Passthrough to next handler if need
		next(w, r)
	}
}
