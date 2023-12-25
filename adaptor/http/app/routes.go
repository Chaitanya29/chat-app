package app

import (
	"chat-app/adaptor/http/controllers"

	"github.com/gin-gonic/gin"
)

func (app *ChatApp) LoadRoutes() *gin.Engine {
	router := gin.Default()
	userController := controllers.NewUserController(app.UserService)
	router.POST("/signup", userController.AddUser)
	return router
}
