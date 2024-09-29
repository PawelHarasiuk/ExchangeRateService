package routes

import (
	"ExchangeRateService/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("symbols", api.GetSymbols)
	r.GET("convert", api.Convert)
	r.GET("rate", api.GetRate)
}
