package api

import (
	"ExchangeRateService/utils"
	"github.com/gin-gonic/gin"
)

func GetSymbols(c *gin.Context) {
	stringJson := utils.GetSymbols()
	c.Data(200, "application/json", []byte(stringJson))
}

func Convert(c *gin.Context) {
	to := c.Query("to")
	from := c.DefaultQuery("from", "EUR")
	amount := c.Query("amount")
	result := utils.GetConvertedValue(to, from, amount)
	c.String(200, "%.2f", result)
}

func GetRate(c *gin.Context) {
	to := c.Query("to")
	from := c.Query("from")

	result := utils.GetRate(to, from)
	c.String(200, "%.2f", result)
}
