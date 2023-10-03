package helpers

import (
	"fmt"
	"receipt-processor-module/models"
	"math"
	"strconv"
	"strings"
)

func GetReceiptById(id string, receipts []models.Receipt) (models.Receipt, string) {
	errorMsg := "Invalid Receipt Id"
	for _, receipt := range(receipts){ // Looping through all the receipts to get the receipt with give ID
		fmt.Println(receipt) 
		if id == receipt.ID {
			return receipt, ""
		}
	}
	var emptyStruct models.Receipt
	return emptyStruct, errorMsg
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
