package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"receipt-processor-module/models"
)


func sendErrorResponse(c *gin.Context, message string){
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": message}) // Sending error response based on error received
}
func sendCustomErrorResponse(c * gin.Context, err error) {
	customError := err.(models.CustomError)
	c.IndentedJSON(customError.HttpCode, gin.H{"error": customError.Message})
}


