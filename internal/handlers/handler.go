package handlers

import (
	"gateway-go/internal/models"
	"github.com/labstack/echo/v4"
)

// RateLimiterRepository is an interface to store rate limiting data
type RateLimiterRepository interface {
	// Get returns the current rate limit for the given identifier
	Get(identifier string) (int, error)

	// Incr increments the rate limit for the given identifier
	Incr(identifier string) error
}

type TransactionsRepository interface {
	CreateTransaction(transaction models.Transaction) error
	GetFailedUserTransactions(username string) []models.Transaction
}

type UsersRepository interface {
	Login(username string, password string) (models.User, error)
	SignUp(username string, password string) error
}

type Handler interface {
	Login(c echo.Context) error
	CreateTransaction(c echo.Context) error
	SendMessage(c echo.Context) error
	SignUp(c echo.Context) error
}

type handler struct {
	transaction TransactionsRepository
	redis       RateLimiterRepository
	user        UsersRepository
}

func NewHandler(transaction TransactionsRepository, redis RateLimiterRepository, user UsersRepository) Handler {
	return &handler{
		transaction: transaction,
		redis:       redis,
		user:        user,
	}
}
