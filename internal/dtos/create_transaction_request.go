package dtos

type CreateTransactionRequest struct {
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
