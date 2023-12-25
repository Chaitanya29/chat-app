package controllers

import (
	"chat-app/internal/model"
	"chat-app/internal/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user.UserService
}

func NewUserController(us user.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

func (us *UserController) AddUser(c *gin.Context) {
	var user model.CreateUserReq
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userResp, err := us.userService.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userResp)
}
