package user

import (
	"chat-app/internal/model"
	"context"
)

type UserService interface {
	CreateUser(c context.Context, req *model.CreateUserReq) (*model.CreateUserRes, error)
}
