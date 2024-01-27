package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

// User Donate Payment godoc
//
// @Summary Download Invoice PDF
// @Description Download the invoice PDF file
// @Tags			User Invoice
// @Security		Bearer
// @Param   invoiceID query  string true "invoice id"
// @Produce octet-stream
// @Success 200 {file} application/pdf
// @Router /user/post/donate/download-invoice  [get]
func DownloadInvoice(ctx *gin.Context, c pb.UserServiceClient) {
	// Set the appropriate headers for the file download
	invoiceID := strings.ToLower(ctx.Query("invoiceID"))
	if invoiceID == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("provide valid Id"))
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read PDF file"})
		return
	}
	ctx.Data(http.StatusOK, "application/pdf", pdfData)
}
