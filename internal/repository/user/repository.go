package user

import (
	"chat-app/internal/model"
	"context"
	"database/sql"

	"github.com/labstack/gommon/log"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) Repository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Insert(ctx context.Context, usr *model.User) (*model.User, error) {
	query := "INSERT INTO tbl_users(id, username, password, email) VALUES ($1, $2, $3, $4)"
	_, err := ur.db.QueryContext(ctx, query, usr.ID, usr.Username, usr.Password, usr.Email)
	if err != nil {
		log.Errorf("insert query err %v", err)
		return nil, err
	}
	return usr, nil
}
