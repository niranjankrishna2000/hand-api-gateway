package routes

import (
	"context"
	"errors"
	"log"
	"net/http"

	unidoc "hand/pkg/pdf"
	"hand/pkg/user/pb"
	user "hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

// User Generate Invoice godoc
//
//	@Summary		generate Invoice PDF
//	@Description	generate the invoice PDF file
//	@Tags			User Invoice
//	@Security		api_key
//	@Param			invoiceID	query	string	true	"invoice id"
//	@Produce		octet-stream
//	@Success		200
//	@Router			/user/post/donate/generate-invoice  [post]
func GenerateInvoice(ctx *gin.Context, c pb.UserServiceClient, usvc user.AuthServiceClient) {
	// Set the appropriate headers for the file download

	//paymentID is Invoice ID
	invoiceID := ctx.Query("invoiceID")
	if invoiceID == "" || invoiceID == " " {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("provide valid Id"))
		return
	}

	res, err := c.GenerateInvoice(context.Background(), &pb.GenerateInvoiceRequest{InvoiceId: invoiceID})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	usr, err := usvc.GetUserDetails(context.Background(), &user.GetUserDetailsRequest{Userid: res.UserID})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("collected data: user",usr)
	//update
	//collect invoice data
	var invoiceItems []*unidoc.InvoiceData
	//data:=unidoc.InvoiceData{Title: "Donation",Quantity: 1,Price: 100}
	log.Println("Adding invoice data")
	data, err := unidoc.NewInvoiceData("Donation",res.FinalPrice)
	if err != nil {
		panic(err)
	}
	invoiceItems = append(invoiceItems, data)
	data, err = unidoc.NewInvoiceData("Transaction fee", 2)
	if err != nil {
		panic(err)
	}
	log.Println("Added invoice data")

	invoiceItems = append(invoiceItems, data)
	// Create single invoice
	log.Println("creating invoice with ",usr.User.Name, invoiceID, res.Address, invoiceItems)
	invoice := unidoc.CreateInvoice(usr.User.Name, invoiceID, res.Address, invoiceItems)
	unidoc.GenerateInvoicePdf(*invoice)
	//log.Printf("The Total Invoice Amount is: %f", invoice.CalculateInvoiceTotalAmount())

	ctx.JSON(http.StatusOK, "Successfully Generated")

}
