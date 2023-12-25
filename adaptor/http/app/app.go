package app

import (
	repo "chat-app/internal/repository/user"
	"chat-app/internal/services/user"
	"context"
	"database/sql"
	"log"
	"net/http"
)

type ChatApp struct {
	UserService user.UserService
}

func NewChatApp(ctx context.Context, db *sql.DB) *ChatApp {
	userService := user.NewUserService(
		repo.NewUserRepository(db),
	)
	return &ChatApp{
		UserService: userService,
	}
}

func (app *ChatApp) Start() {
	router := app.LoadRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
