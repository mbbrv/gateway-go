package handlers

import (
	"fmt"
	"gateway-go/internal/dtos"
	"gateway-go/internal/helper"
	"gateway-go/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) CreateTransaction(c echo.Context) error {
	newTransactionReq := new(dtos.CreateTransactionRequest)
	err := c.Bind(&newTransactionReq)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	userName := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims).UserName
	fmt.Println(newTransactionReq)
	err = h.transaction.CreateTransaction(models.Transaction{
		Amount:   newTransactionReq.Amount,
		Status:   newTransactionReq.Status,
		UserName: userName,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if newTransactionReq.Status == "failed" {
		h.redis.Incr(helper.GetKey(userName, c.Path()))
	}

	return c.String(http.StatusOK, "Transaction created successfully!")
}
