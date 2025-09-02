package controllers

import (
	"game-sale-backend/internal/interfaces/controllers/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {}

func NewUserController() * UserController {
	return &UserController{}
}

func (ctrl *UserController) Register(router *gin.Engine) {
	requestGroup := router.Group("/user")
	{
		requestGroup.POST("", create)
	}
}

func create(requestContext *gin.Context) {
	var user dtos.UserRequest

	if err := requestContext.ShouldBindJSON(&user); err != nil {
		requestContext.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})

		return;
	}

	requestContext.JSON(http.StatusCreated, user)
}