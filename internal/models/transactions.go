package models

import "time"

type Transaction struct {
	ID        int       `db:"id" json:"id"`
	Amount    float64   `db:"amount" json:"amount"`
	UserName  string    `db:"user_name" json:"username"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
