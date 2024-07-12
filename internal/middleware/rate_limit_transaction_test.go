package middleware

import (
	"bytes"
	"gateway-go/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestMiddleware_RateLimitTransaction_Deny(t *testing.T) {
	e := echo.New()
	e.POST("/transaction/create", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	rateLimitRepo := mocks.NewRateLimiterRepository(t)
	rateLimitRepo.On("Get", mock.AnythingOfType("string")).Return(3, nil)

	e.Use(NewMiddleware(rateLimitRepo).RateLimitTransaction)
	go e.Start(":64321")
	defer e.Close()
	jsonStr := []byte(`{"amount":123.3,"status":"failed"}`)
	res, err := http.Post("http://localhost:64321/transaction/create", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("expected status code %d, got %d", http.StatusTooManyRequests, res.StatusCode)
	}

	jsonStr = []byte(`{"amount":123.3,"status":"success"}`)
	res, err = http.Post("http://localhost:64321/transaction/create", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestMiddleware_RateLimitTransaction_Accept(t *testing.T) {
	e := echo.New()
	e.POST("/transaction/create", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	rateLimitRepo := mocks.NewRateLimiterRepository(t)
	rateLimitRepo.On("Get", mock.AnythingOfType("string")).Return(2, nil)

	e.Use(NewMiddleware(rateLimitRepo).RateLimitTransaction)
	go e.Start(":64321")
	defer e.Close()
	jsonStr := []byte(`{"amount":123.3,"status":"failed"}`)
	res, err := http.Post("http://localhost:64321/transaction/create", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusTooManyRequests, res.StatusCode)
	}

	jsonStr = []byte(`{"amount":123.3,"status":"success"}`)
	res, err = http.Post("http://localhost:64321/transaction/create", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}
