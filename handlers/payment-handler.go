package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"task/core"
	"task/middlewares"
	"task/models"
	"task/utils"
)

// AddCreditCardRequest represents the request body for adding a credit card
type AddCreditCardRequest struct {
	StripeToken string `json:"stripe_token"`
}

// AddCreditCard handles adding a credit card for a user
func AddCreditCard(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID := middlewares.GetUserIDFromContext(r.Context())

	var req AddCreditCardRequest

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate input
	if req.StripeToken == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Stripe token is required")
		return
	}

	// Add payment method
	paymentMethod, err := models.AddPaymentMethod(core.DB, userID, req.StripeToken)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, paymentMethod)
}

// DeleteCreditCard handles deleting a credit card for a user
func DeleteCreditCard(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID := middlewares.GetUserIDFromContext(r.Context())

	// Get payment method ID from URL
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid payment method ID")
		return
	}

	// Delete payment method
	err = models.DeletePaymentMethod(core.DB, userID, id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Payment method deleted successfully"})
}

// BuyProducts handles purchasing products
func BuyProducts(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID := middlewares.GetUserIDFromContext(r.Context())

	var req models.OrderRequest

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate input
	if req.PaymentMethodID <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Payment method ID is required")
		return
	}

	if len(req.Items) == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "At least one product item is required")
		return
	}

	// Create order
	order, err := models.CreateOrder(core.DB, userID, req.PaymentMethodID, req.Items)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, order)
}
