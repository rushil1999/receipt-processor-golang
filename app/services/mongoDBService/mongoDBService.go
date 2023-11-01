package mongoDBService

import (
	"receipt-processor-module/pkg/models"
	"fmt"
	"context"
	"receipt-processor-module/pkg/database"
)


func AddReceiptToMongoDB(receipt models.Receipt) error {
	coll := database.GetDatabase()
	result, err := coll.InsertOne(context.TODO(), receipt)

	if err != nil {
		return models.CustomError{HttpCode: 500, Message: "Internal Server Error", DebugMessage:"Could not insert document to database"}
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil

}