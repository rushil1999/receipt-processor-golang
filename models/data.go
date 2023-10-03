package models

import ("github.com/google/uuid")

var Receipts = []Receipt {
	{
		ID: uuid.New().String(),
		Retailer: "ABC1",
		PurchaseDate: "2023-02-10",
		PurchaseTime: "3:29",
		Items: []Item{
			{
				ShortDescription: "Bananans",
				Price: "12",
			},
			{
				ShortDescription: "chocolate",
				Price: "10",
			},
		},
		Total: "22",
	},
	{
		ID: uuid.New().String(),
		Retailer: "ABC2",
		PurchaseDate: "2023-02-10",
		PurchaseTime: "3:29",
		Items: []Item{
			{
				ShortDescription: "Bananans",
				Price: "12",
			},
			{
				ShortDescription: "Pencil",
				Price: "1",
			},
			{
				ShortDescription: "Box",
				Price: "5",
			},
		},
		Total: "18",
	},
}
