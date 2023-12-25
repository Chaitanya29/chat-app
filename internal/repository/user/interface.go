package user

import (
	"chat-app/internal/model"
	"context"
)

type Repository interface {
	Reader
	Writer
}

type Reader interface {
}
type Writer interface {
	Insert(ctx context.Context, user *model.User) (*model.User, error)
}
