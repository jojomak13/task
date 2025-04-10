package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"task/config"
	"task/core"
	"task/middlewares"
	"task/models"
	requests "task/requests/auth"
	"task/utils"
)

type TokenResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var data requests.SignUpRequest

	json.NewDecoder(r.Body).Decode(&data)

	// Validate input
	if err := core.NewValidator(data); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": err,
		})
		return
	}

	user, err := models.CreateUser(core.DB, data)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Generate JWT token
	cfg, _ := config.Load()
	token, err := generateToken(user, cfg.JWTSecret, int64(cfg.JWTExpiration))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Return token and user info
	utils.RespondWithJSON(w, http.StatusCreated, TokenResponse{
		Token: token,
		User:  user,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data requests.LoginRequest

	json.NewDecoder(r.Body).Decode(&data)

	// Validate input
	if err := core.NewValidator(data); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": err,
		})
		return
	}

	// Get user by username
	user, err := models.GetUserByUsername(core.DB, data.Username)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Verify password
	if !user.VerifyPassword(data.Password) {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate JWT token
	cfg, _ := config.Load()
	token, err := generateToken(user, cfg.JWTSecret, int64(cfg.JWTExpiration.Seconds()))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Return token and user info
	utils.RespondWithJSON(w, http.StatusOK, TokenResponse{
		Token: token,
		User:  user,
	})
}

func GetUserOrders(w http.ResponseWriter, r *http.Request) {
	userID := middlewares.GetUserIDFromContext(r.Context())

	orders, err := models.GetUserOrders(core.DB, userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"data": orders,
	})
}

func generateToken(user *models.User, secret string, expiration int64) (string, error) {
	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Second * time.Duration(expiration)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
