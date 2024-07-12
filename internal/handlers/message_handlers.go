package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) SendMessage(c echo.Context) error {
	return c.String(http.StatusOK, "Message sent successfully!")
}
