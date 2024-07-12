package handlers

import (
	"gateway-go/internal/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) SignUp(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if err := h.user.SignUp(username, password); err != nil {
		return err
	}

	h.redis.Incr(helper.GetKey(c.RealIP(), c.Path()))

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created successfully",
	})
}
