package routes

import (
	"context"
	"log"
	"net/http"

	user "hand/pkg/auth/pb"
	unidoc "hand/pkg/pdf"
	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type InvoiceBody struct {
	InvoiceId string `json:"invoiceId" validate:"required"`
}

// User Generate Invoice godoc
//
//	@Summary		generate Invoice PDF
//	@Description	generate the invoice PDF file
//	@Tags			User Invoice
//	@Security		api_key
//	@Param			invoiceID	query	string	true	"invoice id"
//	@Produce		octet-stream
//	@Success		201 {string} string "Successfully Generated"
//	@Failure		400	{object}	string	"Invalid Post ID"
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		500	{object}	string	"Internal server error"
//	@Failure		502	{object}	string	"Error in internal server"
//	@Router			/user/post/donate/generate-invoice  [post]
func GenerateInvoice(ctx *gin.Context, c pb.UserServiceClient, usvc user.AuthServiceClient) {
	// Set the appropriate headers for the file download

	//paymentID is Invoice ID
	id := ctx.Query("invoiceID")
	invoiceID := InvoiceBody{InvoiceId: id}

	validator := validator.New()
	if err := validator.Struct(invoiceID); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, "Invalid Post ID"+err.Error())
		return
	}
	res, err := c.GenerateInvoice(context.Background(), &pb.GenerateInvoiceRequest{InvoiceId: invoiceID.InvoiceId})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, "Error in internal server")
		return
	}

	usr, err := usvc.GetUserDetails(context.Background(), &user.GetUserDetailsRequest{Userid: res.UserID})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, "Error from user server")
		return
	}
	log.Println("collected data: user", usr)
	//update
	//collect invoice data
	var invoiceItems []*unidoc.InvoiceData
	//data:=unidoc.InvoiceData{Title: "Donation",Quantity: 1,Price: 100}
	log.Println("Adding invoice data")
	data, err := unidoc.NewInvoiceData("Donation", res.FinalPrice)
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
	log.Println("creating invoice with ", usr.User.Name, invoiceID, res.Address, invoiceItems)
	invoice := unidoc.CreateInvoice(usr.User.Name, invoiceID.InvoiceId, res.Address, invoiceItems)
	err = unidoc.GenerateInvoicePdf(*invoice)
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusInternalServerError, "Error in internal server")
		return
	}

	ctx.JSON(http.StatusCreated, "Successfully Generated")

}
