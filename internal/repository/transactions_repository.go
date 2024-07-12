package repository

import (
	"gateway-go/internal/models"
	"github.com/jmoiron/sqlx"
)

type TransactionsRepository interface {
	CreateTransaction(transaction models.Transaction) error
	GetFailedUserTransactions(username string) []models.Transaction
}

type TransactionsRepositoryImpl struct {
	db *sqlx.DB
}

func (t TransactionsRepositoryImpl) CreateTransaction(transaction models.Transaction) error {
	db := t.db.MustBegin()
	_, err := db.NamedExec("INSERT INTO transactions (amount, user_name, status) VALUES (:amount, :user_name, :status)", transaction)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()

	return nil
}

func (t TransactionsRepositoryImpl) GetFailedUserTransactions(username string) []models.Transaction {
	var transactions []models.Transaction
	t.db.Select(&transactions, "SELECT * FROM transactions WHERE user_name = $1 AND status = 'failed'", username)
	return transactions
}

func NewTransactionsRepository(db *sqlx.DB) TransactionsRepository {
	return &TransactionsRepositoryImpl{
		db: db,
	}
}
