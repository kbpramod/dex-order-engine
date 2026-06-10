package routes

import (
	"binance-service/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/price/:symbol", controllers.GetPrice)
	router.POST("/order", controllers.CreateOrder)
}