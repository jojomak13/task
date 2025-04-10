package config

import (
	"errors"
	"os"
	"time"
)

type Config struct {
	JWTSecret           string
	JWTExpiration       time.Duration
	StripeSecretKey     string
	StripeWebhookSecret string
}

func Load() (*Config, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret"
	}

	stripeKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripeKey == "" {
		return nil, errors.New("STRIPE_SECRET_KEY is missing")
	}

	stripeWebhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	return &Config{
		JWTSecret:           jwtSecret,
		JWTExpiration:       time.Hour * 24,
		StripeSecretKey:     stripeKey,
		StripeWebhookSecret: stripeWebhookSecret,
	}, nil
}

// // InitStripe initializes the Stripe client with the provided API key
// func InitStripe(key string) {
// 	stripe.Key = key
// }
