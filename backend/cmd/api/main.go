package main

import (
	"game-sale-backend/internal/infrastructure/config/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.RegisterControllers(router)
	router.Run(":8080")
}