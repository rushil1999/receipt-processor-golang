package receiptService

import (
	"receipt-processor-module/pkg/models"
	"receipt-processor-module/app/services"
)

func isReceiptValid(receipt models.Receipt) error {
	retailer := receipt.Retailer
	// purchaseDate := receipt.PurchaseDate
	// PurchaseTime := receipt.PurchaseTime

	if services.IsEmpty(retailer) {
		err := models.CustomError{HttpCode: 400, Message: "Retailer is invalid", DebugMessage: "Retailer is invalid"}
		return err
	}

	return nil

}