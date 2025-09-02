package api

import (
	"game-sale-backend/internal/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)


func RegisterControllers(router *gin.Engine) {
	controllers.NewUserController().Register(router)
}