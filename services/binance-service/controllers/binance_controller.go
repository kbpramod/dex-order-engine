package controllers

import (
	"net/http"

	"binance-service/models"
	"binance-service/services"

	"github.com/gin-gonic/gin"
)

func GetPrice(c *gin.Context) {

	symbol := c.Param("symbol")

	price, err := services.GetPrice(symbol)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, price)
}

func CreateOrder(c *gin.Context) {

	var data models.OrderRequest

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := services.CreateOrder(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}