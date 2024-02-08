package internal

type Invoice struct {
	Name         string
	InvoiceID    string
	Address      string
	InvoiceItems []*InvoiceData
}

func CreateInvoice(name string, invoiceID string, address string, invoiceItems []*InvoiceData) *Invoice {
	return &Invoice{
		Name:         name,
		InvoiceID:    invoiceID,
		Address:      address,
		InvoiceItems: invoiceItems,
	}
}

func (i *Invoice) CalculateInvoiceTotalAmount() float64 {
	var invoiceTotalAmount int = 0
	for _, data := range i.InvoiceItems {
		//amount := data.CalculateTotalAmount()
		invoiceTotalAmount += data.Price
	}

	totalAmount := float64(invoiceTotalAmount) 

	return totalAmount
}
