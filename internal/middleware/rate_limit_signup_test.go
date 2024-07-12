package middleware

import (
	"bytes"
	"gateway-go/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestMiddleware_RateLimitSignup_Deny(t *testing.T) {
	e := echo.New()
	e.POST("/signup", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	rateLimitRepo := mocks.NewRateLimiterRepository(t)
	rateLimitRepo.On("Get", mock.AnythingOfType("string")).Return(20, nil)

	e.Use(NewMiddleware(rateLimitRepo).RateLimitSignup)
	go e.Start(":64321")
	defer e.Close()
	jsonStr := []byte(`{"login":"login","password":"password"}`)
	res, err := http.Post("http://localhost:64321/signup", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("expected status code %d, got %d", http.StatusTooManyRequests, res.StatusCode)
	}
}

func TestMiddleware_RateLimitSignup_Accept(t *testing.T) {
	e := echo.New()
	e.POST("/signup", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	rateLimitRepo := mocks.NewRateLimiterRepository(t)
	rateLimitRepo.On("Get", mock.AnythingOfType("string")).Return(19, nil)

	e.Use(NewMiddleware(rateLimitRepo).RateLimitSignup)
	go e.Start(":64321")
	defer e.Close()
	jsonStr := []byte(`{"login":"login","password":"password"}`)
	res, err := http.Post("http://localhost:64321/signup", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}
