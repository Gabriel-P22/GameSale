package api

import (
	application_user "game-sale-backend/internal/application/User"
	"game-sale-backend/internal/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterControllers(router *gin.Engine) {
	userService := application_user.NewUserService()
	controllers.NewUserController(router, *userService)
}