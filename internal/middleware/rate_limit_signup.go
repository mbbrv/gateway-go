package middleware

import (
	"gateway-go/internal/helper"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Middleware) RateLimitSignup(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := helper.GetKey(c.RealIP(), c.Path())

		signUpCounts, err := m.rateLimiterRepository.Get(key)
		if err != nil && err != redis.Nil {
			return err
		}

		if signUpCounts >= 20 {
			return c.String(http.StatusTooManyRequests, "Too many requests")
		}

		return next(c)
	}
}
