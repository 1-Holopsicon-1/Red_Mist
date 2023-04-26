package middleware

import (
	"RedMist/internal/app/models"
	"RedMist/internal/app/utils"
	"net/http"
)

func AuthzMiddleware(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value("user").(*models.Users)
			if !utils.Contains(roles, user.Role) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
