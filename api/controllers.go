package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/google/uuid"
	"receipt-processor-module/models"
	"receipt-processor-module/helpers"
	"strconv"
	"math"
)

func GetAllReceipts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Receipts)
}

func AddReceipt(c *gin.Context) {
	var newReceipt models.Receipt
	if err := c.ShouldBindJSON(&newReceipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	newReceipt.ID = uuid.New().String()
	models.Receipts = append(models.Receipts, newReceipt)
	c.JSON(http.StatusCreated, gin.H{"id": newReceipt.ID})
}

func GetReceiptPointsById(c *gin.Context) {
	var id = c.Param("id")
	var receipt models.Receipt = helpers.GetReceiptById(id, models.Receipts)
	fmt.Println(receipt)
	totalPoints := 0
	totalPoints += helpers.CountAlphanumeric(receipt.Retailer)
	if helpers.IsRoundedDollarAmount(receipt.Total) {
		totalPoints += 50
	}
	if helpers.IsMultipleOfQuarter(receipt.Total){
		totalPoints += 25
	}
	totalPoints += 5*(len(receipt.Items)/2)
	day := helpers.GetDayFromDate(receipt.PurchaseDate)
	if(day % 2 ==1){
		totalPoints += 6
	}
	if helpers.IsTimeBetween2And4PM(receipt.PurchaseTime) {
		totalPoints += 10
	}

	for _, item := range receipt.Items {
		trimmedDescription := item.ShortDescription
		if len(trimmedDescription) % 3 ==0{
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			points := math.Ceil(itemPrice*0.2)
			totalPoints += int(points) 
		}
	}
	c.JSON(http.StatusCreated, gin.H{"points": totalPoints})
}