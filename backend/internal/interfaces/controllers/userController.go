package controllers

import (
	application_user "game-sale-backend/internal/application/User"
	"game-sale-backend/internal/interfaces/controllers/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService  application_user.UserService
}

func NewUserController(router *gin.Engine, userService  application_user.UserService) * UserController {
	controller := &UserController{
		userService: userService,
	}

	requestGroup := router.Group("/user")
	{
		requestGroup.POST("", controller.create)
	}

	return controller;
}

func (uc *UserController) create(requestContext *gin.Context) {
	var user dtos.UserRequest

	if err := requestContext.ShouldBindJSON(&user); err != nil {
		requestContext.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})

		return;
	}

	var command = application_user.CreateCommand(user);
	var domainUser = uc.userService.Create(command);

	requestContext.JSON(http.StatusCreated, domainUser.ToUserResponse())
}

