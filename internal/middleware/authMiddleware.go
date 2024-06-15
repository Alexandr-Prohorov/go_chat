package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func AuthMiddleware(jwtKey string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Добавляем информацию о пользователе в контекст запроса
		//ctx := context.WithValue(r.Context(), "username", claims.Username)
		//r = r.WithContext(ctx)

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
