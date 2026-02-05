package middleware

import (
	"net/http"
	"strings"

	"workshop-platform/backend/internal/auth"
)

func Auth(j *auth.JWT) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" || !strings.HasPrefix(header, "Bearer ") {
				http.Error(w, `{"error":"missing or invalid authorization"}`, http.StatusUnauthorized)
				return
			}
			token := strings.TrimPrefix(header, "Bearer ")
			claims, err := j.Verify(token)
			if err != nil {
				http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
				return
			}
			ctx := auth.WithClaims(r.Context(), claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
