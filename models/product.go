package models

import (
	"database/sql"
	"errors"
	"fmt"
	"task/requests"
	"time"
)

// Product represents a product in the system
type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateProduct creates a new product in the database
func CreateProduct(db *sql.DB, data requests.ProductRequest) (*Product, error) {
	var product Product
	now := time.Now()
	err := db.QueryRow(
		`INSERT INTO products (name, description, price, stock, created_at, updated_at)
                VALUES ($1, $2, $3, $4, $5, $6)
                RETURNING id, name, description, price, stock, created_at, updated_at`,
		data.Name, data.Description, data.Price, data.Stock, now, now,
	).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

// UpdateProduct updates an existing product in the database
func UpdateProduct(db *sql.DB, id int64, data requests.ProductRequest) (*Product, error) {
	var product Product
	now := time.Now()
	err := db.QueryRow(
		`UPDATE products SET name = $1, description = $2, price = $3, stock = $4, updated_at = $5
                WHERE id = $6
                RETURNING id, name, description, price, stock, created_at, updated_at`,
		data.Name, data.Description, data.Price, data.Stock, now, id,
	).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(db *sql.DB, id int64) error {
	result, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

// GetProductByID retrieves a product by its ID
func GetProductByID(db *sql.DB, id int64) (*Product, error) {
	var product Product
	err := db.QueryRow(
		`SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE id = $1`,
		id,
	).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

// ListProducts retrieves all products from the database
func ListProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query(
		`SELECT id, name, description, price, stock, created_at, updated_at FROM products ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return []*Product{}, nil
	}

	return products, nil
}

// UpdateProductStock updates the stock of a product
func UpdateProductStock(db *sql.DB, id int64, quantity int, tx *sql.Tx) error {
	var query string
	var err error

	query = `UPDATE products SET stock = stock - $1, updated_at = $2 WHERE id = $3 AND stock >= $1`

	now := time.Now()

	var result sql.Result
	if tx != nil {
		result, err = tx.Exec(query, quantity, now, id)
	} else {
		result, err = db.Exec(query, quantity, now, id)
	}

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		// Check if product exists
		var exists bool
		if tx != nil {
			err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)", id).Scan(&exists)
		} else {
			err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)", id).Scan(&exists)
		}

		if err != nil {
			return err
		}

		if !exists {
			return errors.New("product not found")
		}

		return errors.New("insufficient stock")
	}

	return nil
}

// ProductSale represents a product sale record with user information
type ProductSale struct {
	OrderID      int64     `json:"order_id"`
	ProductID    int64     `json:"product_id"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	Total        float64   `json:"total"`
	UserID       int64     `json:"user_id"`
	Username     string    `json:"username"`
	PurchaseDate time.Time `json:"purchase_date"`
}

// GetProductSales retrieves product sales with optional filtering
func GetProductSales(db *sql.DB, fromDate, toDate time.Time, username string) ([]*ProductSale, error) {
	query := `
                SELECT
                        o.id as order_id,
                        p.id as product_id,
                        p.name as product_name,
                        oi.quantity,
                        oi.price,
                        (oi.quantity * oi.price) as total,
                        u.id as user_id,
                        u.username,
                        o.created_at as purchase_date
                FROM
                        orders o
                JOIN
                        order_items oi ON o.id = oi.order_id
                JOIN
                        products p ON oi.product_id = p.id
                JOIN
                        users u ON o.user_id = u.id
                WHERE
                        1=1
        `

	var args []interface{}
	var argIndex = 1

	// Add date filters if provided
	if !fromDate.IsZero() {
		query += fmt.Sprintf(" AND o.created_at >= $%d", argIndex)
		args = append(args, fromDate)
		argIndex++
	}

	if !toDate.IsZero() {
		query += fmt.Sprintf(" AND o.created_at <= $%d", argIndex)
		args = append(args, toDate)
		argIndex++
	}

	// Add username filter if provided
	if username != "" {
		query += fmt.Sprintf(" AND u.username = $%d", argIndex)
		args = append(args, username)
		argIndex++
	}

	// Order by purchase date (newest first)
	query += " ORDER BY o.created_at DESC"

	// Execute the query
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []*ProductSale
	for rows.Next() {
		var sale ProductSale
		err := rows.Scan(
			&sale.OrderID,
			&sale.ProductID,
			&sale.ProductName,
			&sale.Quantity,
			&sale.Price,
			&sale.Total,
			&sale.UserID,
			&sale.Username,
			&sale.PurchaseDate,
		)
		if err != nil {
			return nil, err
		}
		sales = append(sales, &sale)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(sales) == 0 {
		return []*ProductSale{}, nil
	}

	return sales, nil
}
