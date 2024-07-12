package main

import (
	"context"
	"gateway-go/internal/dbs/postgres"
	"gateway-go/internal/handlers"
	middleware2 "gateway-go/internal/middleware"
	"gateway-go/internal/repository"
	"github.com/go-redis/redis"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	e := echo.New()
	ctx := context.Background()
	// Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	db, err := postgres.New(postgres.Config{
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Ssl:      os.Getenv("DB_SSL"),
	}, ctx)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	m := middleware2.NewMiddleware(repository.NewRedisRepository(redisClient))

	handler := handlers.NewHandler(repository.NewTransactionsRepository(db), repository.NewRedisRepository(redisClient), repository.NewUsersRepository(db))
	e.POST("/signup", handler.SignUp, m.RateLimitSignup)
	e.POST("/login", handler.Login)

	a := e.Group("")
	// JWT
	a.Use(echojwt.WithConfig(handlers.JwtConfig))
	a.POST("/transactions/create", handler.CreateTransaction, m.RateLimitTransaction)
	a.POST("/messages/send", handler.SendMessage, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1)))

	e.Logger.Fatal(e.Start(":8000"))
}
