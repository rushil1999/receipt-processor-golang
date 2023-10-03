package helpers

import (
	"fmt"
	"receipt-processor-module/models"
)

func GetReceiptById(id string, receipts []models.Receipt) models.Receipt {
	var receiptWithId models.Receipt
	for _, receipt := range(receipts){
		fmt.Println(receipt)
		if id == receipt.ID {
			fmt.Println("Found")
			receiptWithId = receipt
			break
		}
	}
	return receiptWithId
}