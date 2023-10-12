package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"receipt-processor-module/models"
	"receipt-processor-module/helpers"
)

func GetAllReceipts(c *gin.Context) { // Controller to get all receipts
	c.IndentedJSON(http.StatusOK, models.Receipts)
}

func AddReceipt(c *gin.Context) { // Controller to add a receipt
	var newReceipt models.Receipt
	if err := c.ShouldBindJSON(&newReceipt); err != nil {
		customError := models.CustomError {
			Message: "Bad Input Provided",
			DebugMessage: "Unable to parse the json input",
			HttpCode: 404,
		}
		sendCustomErrorResponse(c, customError)
		return
	}
	newReceiptId, err := helpers.AddReceipt(newReceipt)
	if err != nil {
		customError := models.CustomError {
			Message: "Internal Server Error",
			DebugMessage: "Could not create receipt",
			HttpCode: 500,
		}
		sendCustomErrorResponse(c, customError)
		return	
	}
	c.JSON(http.StatusCreated, gin.H{"id": newReceiptId})
}


func GetReceiptPointsById(c *gin.Context)  { // Controller to get points for a receipt
	var id = c.Param("id")
	receipt, _, err := helpers.GetReceiptById(id) // Getting the receipt from ID
	if err != nil{
		sendCustomErrorResponse(c, err) // Returing error is ID is invalid
		return
	} 
	if(receipt.Points > 0) { // Checking if the points for the receipt are already claculated to save re-calculation
		c.JSON(http.StatusCreated, gin.H{"points": receipt.Points}) 
		return
	}
	
	totalPoints, err := helpers.CalculateReceiptPoints(receipt)
	if err != nil {
		sendCustomErrorResponse(c, err)
		return
	}
	receipt.Points = totalPoints
	c.JSON(http.StatusCreated, gin.H{"points": totalPoints}) // Seinding total points in JSON
}


func GetReceiptById(c *gin.Context) {
	receiptId := c.Param("id")
	receipt, _, err := helpers.GetReceiptById(receiptId)
	if err != nil {
		sendCustomErrorResponse(c, err)
		return
	}
	c.IndentedJSON(200, receipt)  
}

func UpdateReceipt(c *gin.Context) { // Controller to update the receipt
	var updatedReceipt models.Receipt
	if err:= c.ShouldBindJSON(&updatedReceipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	receipt, err := helpers.UpdateReceipt(updatedReceipt)
	if err != nil 	{
		sendCustomErrorResponse(c, err)
	}
	c.IndentedJSON(http.StatusCreated, receipt)

}

func DeleteReceipt(c * gin.Context) {
	id := c.Param("id")
	deletedReceiptId, err := helpers.DeleteReceipt(id)
	if err != nil {
		sendCustomErrorResponse(c, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"recetipDeleteId": deletedReceiptId})
}