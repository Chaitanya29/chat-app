package main

import (
	"chat-app/adaptor/http/app"
	db "chat-app/database"
	"context"
	"time"

	"github.com/labstack/gommon/log"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	dbConn, err := db.NewDBConnection()
	if err != nil {
		log.Errorf("database connection error: %v", err)
		return
	}
	app.NewChatApp(ctx, dbConn.GetDB()).Start()

	// fmt.Println("fmt")
}
