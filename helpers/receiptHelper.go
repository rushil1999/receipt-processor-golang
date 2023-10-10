package helpers

import (
	"fmt"
	"receipt-processor-module/models"
	"math"
	"strconv"
	"strings"
)

func GetReceiptById(id string) (models.Receipt, int, string) {
	errorMsg := "Receipt not found"
	receipts := models.Receipts
	for i, receipt := range(receipts) {
		if receiptId := receipt.ID; receiptId == id {
			return receipt, i, ""
		}
	}
	var emptyReceipt models.Receipt
	return emptyReceipt, -1, errorMsg
}

func GetItemPoints(items []models.Item) int { 
	totalPoints := 0
	for _, item := range items { // Looping through all the items to get the points
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription) % 3 ==0{
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			points := math.Ceil(itemPrice*0.2)
			totalPoints += int(points) 
		}
	}
	return totalPoints
}

func CalculateReceiptPoints(receipt models.Receipt) (int, string) {
	totalPoints := 5*(len(receipt.Items)/2) // Adding points for receipts pairs
	totalPoints += CountAlphanumeric(receipt.Retailer) // Counting alphanumeric characters
	isRoundedDollarAmount, errorString := IsRoundedDollarAmount(receipt.Total)
	if errorString != "" {
		return -1, errorString
	}
	if isRoundedDollarAmount { // Checking if the amount is rounded and adding points accordingly
		totalPoints += 50 
	}
	isMultipleOfQuarter, errorString := IsMultipleOfQuarter(receipt.Total)
	if errorString != "" {
		return -1, errorString
	}
	if isMultipleOfQuarter{
		totalPoints += 25
	}
	day, errorString := GetDayFromDate(receipt.PurchaseDate) //Getting Day from YYYY-MM-DD format
	if(errorString !=  ""){
		return -1, errorString
	}
	if(day % 2 ==1){
		totalPoints += 6
	}
	isTimeBetween2And4PM, errorString := IsTimeBetween2And4PM(receipt.PurchaseTime)
	if errorString != "" {
		return -1, errorString
	}
	if isTimeBetween2And4PM {
		totalPoints += 10
	}
	totalPoints += GetItemPoints(receipt.Items)
	return totalPoints, ""
}


func UpdateReceipt(updatedReceipt models.Receipt) (models.Receipt, string) {
	receipt, index, err := GetReceiptById(updatedReceipt.ID)
	var emptyReceipt models.Receipt
	if err != "" {
		return emptyReceipt, "Receipt not found" 
	}
	receipt.Retailer = updatedReceipt.Retailer
	receipt.PurchaseDate = updatedReceipt.PurchaseDate
	receipt.PurchaseTime = updatedReceipt.PurchaseTime
	receipt.Items = updatedReceipt.Items
	receipt.Total = updatedReceipt.Total
	newReceiptPoints, err := CalculateReceiptPoints(receipt)
	if err != "" {
		return emptyReceipt, "invalid receipt"
	}
	receipt.Points = newReceiptPoints
	models.Receipts[index] = receipt
	return receipt, ""
}