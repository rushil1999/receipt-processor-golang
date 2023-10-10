package api

import (
	"github.com/gin-gonic/gin"
)

func ImportRoutes(routerEngine *gin.Engine) {
	routerEngine.GET("/receipts", GetAllReceipts) // Route to get all receipts 
	routerEngine.POST("/receipts/process", AddReceipt) //Route to add a receipt
	routerEngine.GET("/receipts/:id/process", GetReceiptPointsById) //Route to calculate points in a receipt
	routerEngine.PUT("/receipts/process", UpdateReceipt) //Router to update a receipt
}
