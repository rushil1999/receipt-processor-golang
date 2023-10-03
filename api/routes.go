package api

import (
	"github.com/gin-gonic/gin"
)

func ImportRoutes(routerEngine *gin.Engine) {
	routerEngine.GET("/receipts", GetAllReceipts)
	routerEngine.POST("/receipts/process", AddReceipt)
	routerEngine.GET("/receipts/:id/process", GetReceiptPointsById)
}
