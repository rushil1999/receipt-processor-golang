package models

type Receipt struct {
	ID string `json:"receiptId"`
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items []Item `json:"items"`
	Total string `json:"total"`
	Points int `json:"-"`
}