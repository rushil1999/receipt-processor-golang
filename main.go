package main

import (
	"github.com/gin-gonic/gin"
	"receipt-processor-module/api"
	// "log"
)

func main() {
	router := gin.Default()
	api.ImportRoutes(router)
	router.Run("localhost:8000")
	// Logger := log.New("DEBUG: ", log.Llongfile)
}