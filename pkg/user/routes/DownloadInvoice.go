package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type InvoiceIdBody struct {
	InvoiceId string `json:"invoiceId" validate:"required,ascii"` //test
}

// Download Invoice godoc
//
//	@Summary		Download Invoice
//	@Description	Download the invoice PDF file
//	@Tags			User Invoice
//	@Security		api_key
//	@Param			invoiceID	query	string	true	"invoice id"
//	@Produce		octet-stream
//	@Success		200	{file}		application/pdf
//	@Failure		400	{string}	string	"Couldn't fetch data from client "
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		502	{string}	string	"Error in internal server"
//	@Router			/user/post/donate/download-invoice  [get]
func DownloadInvoice(ctx *gin.Context, c pb.UserServiceClient) {

	invoiceID := strings.ToLower(ctx.Query("invoiceID"))
	if invoiceID == "" {
		ctx.JSON(http.StatusInternalServerError, "invalid invoice id")
		return
	}
	invoiceBody := InvoiceIdBody{InvoiceId: invoiceID}
	validator := validator.New()
	if err := validator.Struct(invoiceBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, "Invalid data"+err.Error())
		return
	}
	headerfilename := fmt.Sprintf("attachment; filename=%s.pdf", invoiceID)
	ctx.Header("Content-Disposition", headerfilename)
	ctx.Header("Content-Type", "application/pdf")
	userID := ctx.GetInt64("userId")
	log.Println("User ID:", userID)
	fmt.Println("Collected User Id and Invoice ID :", userID, invoiceID)

	//edit  get userid AND check if valid

	// res, err := c.MakePaymentRazorPay(context.Background(), &pb.MakePaymentRazorPayRequest{Postid: int32(postID), Userid: int32(userID),Amount: int32(amount)})

	// if err != nil {
	// 	ctx.AbortWithError(http.StatusBadGateway, err)
	// 	return
	// }

	// Read the PDF file and write it to the response
	filename := fmt.Sprintf("src/invoice/%s.pdf", invoiceID)
	pdfData, err := os.ReadFile(filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to read PDF file")
		return
	}
	// if err != nil {
	// 	log.Println("Error with internal server :", err)
	// 	ctx.JSON(http.StatusBadGateway, pb.ClearNotificationResponse{
	// 		Status:   http.StatusBadGateway,
	// 		Response: "Error in internal server",
	// 	})
	// 	return
	// }
	ctx.Data(http.StatusOK, "application/pdf", pdfData)
}
