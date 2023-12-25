package user

import (
	"chat-app/internal/model"
	"chat-app/internal/repository/user"
	"context"

	"time"

	"github.com/labstack/gommon/log"

	"github.com/google/uuid"
)

func NewUserService(repo user.Repository) UserService {
	return &userService{
		Repo:    repo,
		timeout: time.Duration(2) * time.Second,
	}
}

type userService struct {
	Repo    user.Repository
	timeout time.Duration
}

func (us *userService) CreateUser(ctx context.Context, createUser *model.CreateUserReq) (*model.CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, us.timeout)
	defer cancel()
	usr := model.User{
		ID:       uuid.New().String(),
		Username: createUser.Username,
		Password: createUser.Password,
		Email:    createUser.Email,
	}
	user, err := us.Repo.Insert(ctx, &usr)
	if err != nil {
		log.Errorf("insert error %v", err)
	}

	return &model.CreateUserRes{ID: user.ID, Username: user.Username, Email: user.Email}, nil
}
