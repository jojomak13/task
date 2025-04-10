package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"task/config"
	"task/utils"

	"github.com/dgrijalva/jwt-go"
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

// UserIDKey is the key used to store the user ID in the request context
const UserIDKey contextKey = "userID"

// UserRoleKey is the key used to store the user role in the request context
const UserRoleKey contextKey = "userRole"

// Authenticate middleware validates JWT tokens and sets user info in the request context
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header is required")
			return
		}

		// Extract the token
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Token is required")
			return
		}

		// Parse and validate the token
		token, err := parseToken(tokenString)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Get claims from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		// Extract user ID and role from claims
		userID, ok := claims["id"].(float64)
		if !ok {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid user ID in token")
			return
		}

		userRole, ok := claims["role"].(string)
		if !ok {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid user role in token")
			return
		}

		// Set user info in context for downstream handlers
		ctx := context.WithValue(r.Context(), UserIDKey, int64(userID))
		ctx = context.WithValue(ctx, UserRoleKey, userRole)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserIDFromContext extracts the user ID from the request context
func GetUserIDFromContext(ctx context.Context) int64 {
	userID, ok := ctx.Value(UserIDKey).(int64)
	if !ok {
		return 0 // Default value if not found
	}
	return userID
}

// GetUserRoleFromContext extracts the user role from the request context
func GetUserRoleFromContext(ctx context.Context) string {
	role, ok := ctx.Value(UserRoleKey).(string)
	if !ok {
		return "" // Default value if not found
	}
	return role
}

// parseToken validates and parses a JWT token
func parseToken(tokenString string) (*jwt.Token, error) {
	cfg, _ := config.Load()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// Return the secret key
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
