package middlewares

import (
	"net/http"

	"task/models"
	"task/utils"
)

// RequireAdmin middleware checks if the authenticated user has admin role
func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get user role from context (set by auth middleware)
		role := GetUserRoleFromContext(r.Context())

		// Check if role is admin
		if role != string(models.RoleAdmin) {
			utils.RespondWithError(w, http.StatusForbidden, "Admin privileges required")
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
