# E-Commerce API Documentation

This document provides information on how to use the E-Commerce API for user management, product operations, and payment processing.

## Installation
It will be easy to install it all you need is to have docker installed on you machine after that run

1. run `cp .env.example .env`
2. run `docker-compoer up -d`

## Base URL

The base URL for all API endpoints is:
```
http://localhost:5000/api
```

## Authentication

Most endpoints require authentication. To authenticate:

1. Sign up or login to get a JWT token
2. Include the token in the Authorization header of your requests:
```
Authorization: Bearer YOUR_TOKEN_HERE
```

## User Management Endpoints

### Sign Up

Register a new user.

**URL**: `/signup`
**Method**: `POST`
**Auth required**: No

**Request Body**:
```json
{
  "username": "username",
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Success Response**: `201 Created`
```json
{
  "token": "jwt_token",
  "user": {
    "id": 1,
    "username": "username",
    "email": "user@example.com",
    "role": "user",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
}
```

### Login

Authenticate a user.

**URL**: `/login`
**Method**: `POST`
**Auth required**: No

**Request Body**:
```json
{
  "username": "username",
  "password": "securepassword"
}
```

**Success Response**: `200 OK`
```json
{
  "token": "jwt_token",
  "user": {
    "id": 1,
    "username": "username",
    "email": "user@example.com",
    "role": "user",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
}
```

## User Routes (Require Authentication)

### Add Payment Method

Add a credit card to the user account.

**URL**: `/user/payment-methods`
**Method**: `POST`
**Auth required**: Yes

**Request Body**:
```json
{
  "stripe_token": "stripe_token_from_frontend"
}
```

**Success Response**: `201 Created`
```json
{
  "id": 1,
  "user_id": 1,
  "stripe_id": "pm_...",
  "card_brand": "visa",
  "card_last4": "4242",
  "card_exp_month": 12,
  "card_exp_year": 2024,
  "is_default": true,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Delete Payment Method

Remove a credit card from the user account.

**URL**: `/user/payment-methods/{id}`
**Method**: `DELETE`
**Auth required**: Yes

**Success Response**: `200 OK`
```json
{
  "message": "Payment method deleted successfully"
}
```

### List Products

Get all available products.

**URL**: `/user/products`
**Method**: `GET`
**Auth required**: Yes

**Success Response**: `200 OK`
```json
[
  {
    "id": 1,
    "name": "Product Name",
    "description": "Product description",
    "price": 99.99,
    "stock": 10,
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
]
```

### Buy Products

Purchase products.

**URL**: `/user/purchase`
**Method**: `POST`
**Auth required**: Yes

**Request Body**:
```json
{
  "payment_method_id": 1,
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 2,
      "quantity": 1
    }
  ]
}
```

**Success Response**: `201 Created`
```json
{
  "id": 1,
  "user_id": 1,
  "total": 299.97,
  "stripe_payment_id": "pi_...",
  "status": "completed",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z",
  "items": [
    {
      "id": 1,
      "order_id": 1,
      "product_id": 1,
      "product_name": "Product Name",
      "quantity": 2,
      "price": 99.99,
      "created_at": "2023-01-01T00:00:00Z"
    },
    {
      "id": 2,
      "order_id": 1,
      "product_id": 2,
      "product_name": "Another Product",
      "quantity": 1,
      "price": 99.99,
      "created_at": "2023-01-01T00:00:00Z"
    }
  ]
}
```

### Get Orders

Retrieve a user's order history.

**URL**: `/user/orders`
**Method**: `GET`
**Auth required**: Yes

**Success Response**: `200 OK`
```json
[
  {
    "id": 1,
    "user_id": 1,
    "total": 299.97,
    "stripe_payment_id": "pi_...",
    "status": "completed",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z",
    "items": [
      {
        "id": 1,
        "order_id": 1,
        "product_id": 1,
        "product_name": "Product Name",
        "quantity": 2,
        "price": 99.99,
        "created_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
]
```

## Admin Routes (Require Admin Authentication)

### Create Product

Add a new product.

**URL**: `/admin/products`
**Method**: `POST`
**Auth required**: Yes (Admin only)

**Request Body**:
```json
{
  "name": "New Product",
  "description": "Product description",
  "price": 99.99,
  "stock": 100
}
```

**Success Response**: `201 Created`
```json
{
  "id": 1,
  "name": "New Product",
  "description": "Product description",
  "price": 99.99,
  "stock": 100,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Update Product

Update an existing product.

**URL**: `/admin/products/{id}`
**Method**: `PUT`
**Auth required**: Yes (Admin only)

**Request Body**:
```json
{
  "name": "Updated Product",
  "description": "Updated description",
  "price": 149.99,
  "stock": 75
}
```

**Success Response**: `200 OK`
```json
{
  "id": 1,
  "name": "Updated Product",
  "description": "Updated description",
  "price": 149.99,
  "stock": 75,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-02T00:00:00Z"
}
```

### Delete Product

Remove a product.

**URL**: `/admin/products/{id}`
**Method**: `DELETE`
**Auth required**: Yes (Admin only)

**Success Response**: `200 OK`
```json
{
  "message": "Product deleted successfully"
}
```

### Get Product Sales

Retrieve product sales with optional filtering.

**URL**: `/admin/product-sales`
**Method**: `GET`
**Auth required**: Yes (Admin only)

**Query Parameters**:
- `from_date`: Filter sales from this date (YYYY-MM-DD)
- `to_date`: Filter sales to this date (YYYY-MM-DD)
- `username`: Filter sales by username

**Success Response**: `200 OK`
```json
[
  {
    "order_id": 1,
    "product_id": 1,
    "product_name": "Product Name",
    "quantity": 2,
    "price": 99.99,
    "total": 199.98,
    "user_id": 1,
    "username": "username",
    "purchase_date": "2023-01-01T00:00:00Z"
  }
]
```

## Error Responses

**400 Bad Request**
```json
{
  "error": "Invalid request payload"
}
```

**401 Unauthorized**
```json
{
  "error": "Invalid credentials"
}
```

**403 Forbidden**
```json
{
  "error": "Admin privileges required"
}
```

**500 Internal Server Error**
```json
{
  "error": "Error message"
}
```
