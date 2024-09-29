package main

import (
	"ExchangeRateService/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	if err := r.Run("localhost:8080"); err != nil {
		fmt.Println(err.Error())
		return
	}
}
