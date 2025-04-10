package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/paymentmethod"
)

// PaymentMethod represents a user's payment method
type PaymentMethod struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	StripeID     string    `json:"stripe_id"`
	CardBrand    string    `json:"card_brand"`
	CardLast4    string    `json:"card_last4"`
	CardExpMonth int       `json:"card_exp_month"`
	CardExpYear  int       `json:"card_exp_year"`
	IsDefault    bool      `json:"is_default"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Order represents a user's order
type Order struct {
	ID              int64        `json:"id"`
	UserID          int64        `json:"user_id"`
	Total           float64      `json:"total"`
	StripePaymentID string       `json:"stripe_payment_id"`
	Status          string       `json:"status"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	Items           []*OrderItem `json:"items,omitempty"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID          int64     `json:"id"`
	OrderID     int64     `json:"order_id"`
	ProductID   int64     `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

// OrderRequest represents a request to buy products
type OrderRequest struct {
	PaymentMethodID int64              `json:"payment_method_id"`
	Items           []OrderItemRequest `json:"items"`
}

// OrderItemRequest represents a product to be purchased
type OrderItemRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

// AddPaymentMethod adds a new payment method for a user
func AddPaymentMethod(db *sql.DB, userID int64, stripeToken string) (*PaymentMethod, error) {
	// Get user's Stripe customer ID or create one if it doesn't exist
	var stripeCustomerID string
	err := db.QueryRow("SELECT stripe_customer_id FROM users WHERE id = $1", userID).Scan(&stripeCustomerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// If user doesn't have a Stripe customer ID yet, create one
	if stripeCustomerID == "" {
		// Get user information for creating Stripe customer
		var username, email string
		err := db.QueryRow("SELECT username, email FROM users WHERE id = $1", userID).Scan(&username, &email)
		if err != nil {
			return nil, err
		}

		// Create Stripe customer
		customerParams := &stripe.CustomerParams{
			Name:  stripe.String(username),
			Email: stripe.String(email),
		}
		c, err := customer.New(customerParams)
		if err != nil {
			return nil, err
		}
		stripeCustomerID = c.ID

		// Save Stripe customer ID to user record
		_, err = db.Exec("UPDATE users SET stripe_customer_id = $1 WHERE id = $2", stripeCustomerID, userID)
		if err != nil {
			return nil, err
		}
	}

	// Attach the payment method to the customer
	pmParams := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(stripeCustomerID),
	}
	pm, err := paymentmethod.Attach(stripeToken, pmParams)
	if err != nil {
		return nil, err
	}

	// Check if this is the user's first payment method (make it default if so)
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM payment_methods WHERE user_id = $1", userID).Scan(&count)
	if err != nil {
		return nil, err
	}
	isDefault := count == 0

	// Save payment method details to database
	var paymentMethod PaymentMethod
	now := time.Now()

	err = db.QueryRow(
		`INSERT INTO payment_methods
		(user_id, stripe_id, card_brand, card_last4, card_exp_month, card_exp_year, is_default, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, user_id, stripe_id, card_brand, card_last4, card_exp_month, card_exp_year, is_default, created_at, updated_at`,
		userID, pm.ID, pm.Card.Brand, pm.Card.Last4, pm.Card.ExpMonth, pm.Card.ExpYear, isDefault, now, now,
	).Scan(
		&paymentMethod.ID, &paymentMethod.UserID, &paymentMethod.StripeID,
		&paymentMethod.CardBrand, &paymentMethod.CardLast4,
		&paymentMethod.CardExpMonth, &paymentMethod.CardExpYear,
		&paymentMethod.IsDefault, &paymentMethod.CreatedAt, &paymentMethod.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &paymentMethod, nil
}

// DeletePaymentMethod deletes a user's payment method
func DeletePaymentMethod(db *sql.DB, userID, paymentMethodID int64) error {
	// Get the payment method
	var stripeID string
	var isDefault bool
	err := db.QueryRow(
		"SELECT stripe_id, is_default FROM payment_methods WHERE id = $1 AND user_id = $2",
		paymentMethodID, userID,
	).Scan(&stripeID, &isDefault)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("payment method not found")
		}
		return err
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Delete the payment method from the database
	_, err = tx.Exec("DELETE FROM payment_methods WHERE id = $1", paymentMethodID)
	if err != nil {
		return err
	}

	// If this was the default payment method, make another one default if available
	if isDefault {
		var newDefaultID int64
		err = tx.QueryRow(
			"SELECT id FROM payment_methods WHERE user_id = $1 LIMIT 1",
			userID,
		).Scan(&newDefaultID)

		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if err != sql.ErrNoRows {
			// Set the new default payment method
			_, err = tx.Exec(
				"UPDATE payment_methods SET is_default = true WHERE id = $1",
				newDefaultID,
			)
			if err != nil {
				return err
			}
		}
	}

	// Detach the payment method from Stripe
	_, err = paymentmethod.Detach(stripeID, nil)
	if err != nil {
		return err
	}

	// Commit the transaction
	return tx.Commit()
}

// GetUserPaymentMethods retrieves all payment methods for a user
func GetUserPaymentMethods(db *sql.DB, userID int64) ([]*PaymentMethod, error) {
	rows, err := db.Query(`
		SELECT id, user_id, stripe_id, card_brand, card_last4, card_exp_month, card_exp_year, is_default, created_at, updated_at
		FROM payment_methods
		WHERE user_id = $1
		ORDER BY is_default DESC, created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []*PaymentMethod
	for rows.Next() {
		var pm PaymentMethod
		err := rows.Scan(
			&pm.ID, &pm.UserID, &pm.StripeID, &pm.CardBrand, &pm.CardLast4,
			&pm.CardExpMonth, &pm.CardExpYear, &pm.IsDefault, &pm.CreatedAt, &pm.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, &pm)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return paymentMethods, nil
}

// CreateOrder creates a new order with the given items
func CreateOrder(db *sql.DB, userID int64, paymentMethodID int64, items []OrderItemRequest) (*Order, error) {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Get the payment method
	var paymentMethod PaymentMethod
	err = tx.QueryRow(`
		SELECT id, user_id, stripe_id FROM payment_methods
		WHERE id = $1 AND user_id = $2`,
		paymentMethodID, userID,
	).Scan(&paymentMethod.ID, &paymentMethod.UserID, &paymentMethod.StripeID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("payment method not found")
		}
		return nil, err
	}

	// Verify each product and calculate total
	var total float64 = 0
	var orderItems []*OrderItem

	for _, item := range items {
		if item.Quantity <= 0 {
			return nil, errors.New("quantity must be greater than zero")
		}

		// Get product details
		var product Product
		err = tx.QueryRow(`
			SELECT id, name, price, stock FROM products
			WHERE id = $1`,
			item.ProductID,
		).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("product with ID %d not found", item.ProductID)
			}
			return nil, err
		}

		// Check if enough stock is available
		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product: %s", product.Name)
		}

		// Update product stock
		err = UpdateProductStock(db, product.ID, item.Quantity, tx)
		if err != nil {
			return nil, err
		}

		// Add to running total
		itemTotal := product.Price * float64(item.Quantity)
		total += itemTotal

		// Create order item
		orderItem := &OrderItem{
			ProductID:   product.ID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			Price:       product.Price,
		}
		orderItems = append(orderItems, orderItem)
	}

	// Create the order
	var order Order
	now := time.Now()

	// At this point we should process the payment with Stripe, but since we're focusing on the backend structure,
	// we'll leave a comment for where this would happen and use a placeholder "stripe_payment_id"
	// In a real application, you would create a PaymentIntent with Stripe here

	stripePaymentID := "placeholder_stripe_payment_id"

	err = tx.QueryRow(`
		INSERT INTO orders (user_id, total, stripe_payment_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, user_id, total, stripe_payment_id, status, created_at, updated_at`,
		userID, total, stripePaymentID, "completed", now, now,
	).Scan(
		&order.ID, &order.UserID, &order.Total, &order.StripePaymentID,
		&order.Status, &order.CreatedAt, &order.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Create order items
	for _, item := range orderItems {
		err = tx.QueryRow(`
			INSERT INTO order_items (order_id, product_id, product_name, quantity, price, created_at)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id, created_at`,
			order.ID, item.ProductID, item.ProductName, item.Quantity, item.Price, now,
		).Scan(&item.ID, &item.CreatedAt)

		if err != nil {
			return nil, err
		}

		item.OrderID = order.ID
	}

	// Assign items to order
	order.Items = orderItems

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// GetUserOrders retrieves all orders for a user
func GetUserOrders(db *sql.DB, userID int64) ([]*Order, error) {
	// First get all orders
	rows, err := db.Query(`
		SELECT id, user_id, total, stripe_payment_id, status, created_at, updated_at
		FROM orders
		WHERE user_id = $1
		ORDER BY created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*Order
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID, &order.UserID, &order.Total, &order.StripePaymentID,
			&order.Status, &order.CreatedAt, &order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return []*Order{}, nil
	}

	return orders, nil
}
