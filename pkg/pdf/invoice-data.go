package internal

import (
	"errors"
	"log"
)

type InvoiceData struct {
	Title string
	//Quantity    int
	Price int
	//TotalAmount int
}

// func (d *InvoiceData) CalculateTotalAmount() int {
// 	totalAmount := d.Quantity * d.Price
// 	return totalAmount
// }

// func (d *InvoiceData) ReturnItemTotalAmount() float64 {
// 	totalAmount := d.CalculateTotalAmount()
// 	converted := float64(totalAmount) / 100
// 	return converted
// }

func (d *InvoiceData) ReturnItemPrice() float64 {
	converted := float64(d.Price) * 100
	return converted
}

func NewInvoiceData(title string, price interface{}) (*InvoiceData, error) {
	log.Println("Starting Unidoc New invoice data")
	var convertedPrice int

	switch priceValue := price.(type) {
	case int:
		convertedPrice = priceValue
	case int64:
		convertedPrice = int(priceValue)
	case float32:
		convertedPrice = int(priceValue)
	case float64:
		convertedPrice = int(priceValue)
	default:
		return nil, errors.New("type not permitted")
	}

	return &InvoiceData{
		Title: title,
		//Quantity: qty,
		Price: convertedPrice,
	}, nil
}
