package main

import (
	"github.com/gin-gonic/gin"
	"receipt-processor-module/app/api"

	"receipt-processor-module/pkg/database"
	// "log"
)

func main() {

	database.InitMongoDBConnection()
	
	router := gin.Default()
	api.ImportRoutes(router)
	router.Run("localhost:8000")
}

	
	
	// Logger := log.New("DEBUG: ", log.Llongfile)