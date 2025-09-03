package main

import (
	api "game-sale-backend/internal/interfaces/config/DI/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.RegisterControllers(router)
	router.Run(":8080")
}