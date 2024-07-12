package middleware

import (
	"bytes"
	"encoding/json"
	"gateway-go/internal/handlers"
	"gateway-go/internal/helper"
	"gateway-go/internal/models"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (m *Middleware) RateLimitTransaction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var bodyBytes []byte
		if c.Request().Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request().Body)
		}
		// Restore the io.ReadCloser to its original state
		c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		transaction := models.Transaction{}
		err := json.Unmarshal(bodyBytes, &transaction)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if transaction.Status == "failed" {
			key := helper.GetKey("", c.Path())
			if c.Get("user") != nil {
				userName := c.Get("user").(*jwt.Token).Claims.(*handlers.JwtCustomClaims).UserName
				key = helper.GetKey(userName, c.Path())
			}
			failedTransactionCounts, err := m.rateLimiterRepository.Get(key)
			if err != nil && err != redis.Nil {
				return err
			}

			if failedTransactionCounts >= 3 {
				return c.String(http.StatusTooManyRequests, "Too many requests")
			}
		}

		return next(c)
	}
}
