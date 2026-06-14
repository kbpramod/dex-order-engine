package main

import (
	"services/internal/binance/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}