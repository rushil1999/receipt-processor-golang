package helpers

import (
	"receipt-processor-module/models"
	"math"
	"strconv"
	"strings"
)

func GetReceiptById(id string) (models.Receipt, int, error) {
	receipts := models.Receipts
	for i, receipt := range(receipts) {
		if receiptId := receipt.ID; receiptId == id {
			return receipt, i, nil
		}
	}
	var emptyReceipt models.Receipt
	var errorInstance = models.CustomError {
		Message: "Resource not found",
		DebugMessage: "Receipt not found",
		HttpCode: 404,
	}
	return emptyReceipt, -1, errorInstance
}

func GetItemPoints(items []models.Item) (int, error) { 
	totalPoints := 0
	for _, item := range items { // Looping through all the items to get the points
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription) % 3 ==0{
			itemPrice, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				customError := models.CustomError {
					Message: "Invalid price",
					DebugMessage: "cannot parse invalid number to float",
					HttpCode: 400,
				}
				return -1, customError
			}
			points := math.Ceil(itemPrice*0.2)
			totalPoints += int(points) 
		}
	}
	return totalPoints, nil
}

func CalculateReceiptPoints(receipt models.Receipt) (int, error) {
	totalPoints := 5*(len(receipt.Items)/2) // Adding points for receipts pairs
	totalPoints += CountAlphanumeric(receipt.Retailer) // Counting alphanumeric characters
	isRoundedDollarAmount, err := IsRoundedDollarAmount(receipt.Total)
	if err != nil {
		return -1, err
	}
	if isRoundedDollarAmount { // Checking if the amount is rounded and adding points accordingly
		totalPoints += 50 
	}
	isMultipleOfQuarter, err := IsMultipleOfQuarter(receipt.Total)
	if err != nil {
		return -1, err
	}
	if isMultipleOfQuarter{
		totalPoints += 25
	}
	day, err := GetDayFromDate(receipt.PurchaseDate) //Getting Day from YYYY-MM-DD format
	if(err != nil){
		return -1, err
	}
	if(day % 2 ==1){
		totalPoints += 6
	}
	isTimeBetween2And4PM, err := IsTimeBetween2And4PM(receipt.PurchaseTime)
	if err != nil {
		return -1, err
	}
	if isTimeBetween2And4PM {
		totalPoints += 10
	}
	itemPoints, err := GetItemPoints(receipt.Items)
	if err != nil {
		return -1, err
	}
	totalPoints += itemPoints
	return totalPoints, nil
}


func UpdateReceipt(updatedReceipt models.Receipt) (models.Receipt, error) {
	receipt, index, err := GetReceiptById(updatedReceipt.ID)
	var emptyReceipt models.Receipt
	if err != nil {
		return emptyReceipt, err
	}
	receipt.Retailer = updatedReceipt.Retailer
	receipt.PurchaseDate = updatedReceipt.PurchaseDate
	receipt.PurchaseTime = updatedReceipt.PurchaseTime
	receipt.Items = updatedReceipt.Items
	receipt.Total = updatedReceipt.Total
	newReceiptPoints, err := CalculateReceiptPoints(receipt)
	if err != nil {
		return emptyReceipt, err
	}
	receipt.Points = newReceiptPoints
	models.Receipts[index] = receipt
	return receipt, nil
}