package repository

import (
	"gateway-go/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type UsersRepository interface {
	Login(username string, password string) (models.User, error)
	SignUp(username string, password string) error
}

type UsersRepositoryImpl struct {
	db *sqlx.DB
}

func (u UsersRepositoryImpl) Login(username string, password string) (models.User, error) {
	var user models.User
	err := u.db.Get(&user, "SELECT * FROM users WHERE username = $1 AND password = $2", username, password)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	return user, nil
}

func (u UsersRepositoryImpl) SignUp(username string, password string) error {
	db := u.db.MustBegin()
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()

	return nil
}

func NewUsersRepository(db *sqlx.DB) UsersRepository {
	return &UsersRepositoryImpl{
		db: db,
	}
}
