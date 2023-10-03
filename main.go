package main

import (
	"github.com/gin-gonic/gin"
	"receipt-processor-module/api"
)

func main() {
	router := gin.Default()
	api.ImportRoutes(router)
	router.Run("localhost:8000")
}