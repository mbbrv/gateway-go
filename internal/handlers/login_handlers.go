package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var JwtConfig = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(JwtCustomClaims)
	},
	SigningKey: []byte("secret"),
}

type JwtCustomClaims struct {
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

func (h *handler) Login(c echo.Context) error {
	username := c.FormValue("username")
	pass := c.FormValue("password")

	user, err := h.user.Login(username, pass)
	if err != nil {
		return c.String(http.StatusUnauthorized, "")
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
