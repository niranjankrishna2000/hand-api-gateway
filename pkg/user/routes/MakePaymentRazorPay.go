package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type PayIdBody struct {
	PayId int `json:"payId" validate:"required,min=1,number"`
}

// Donate Payment godoc
//
//	@Summary		Donate
//	@Description	User can pay for donation
//	@Tags			User Donations
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			payid	query		string	true	"pay id Data"
//	@Success		200		{object}	pb.MakePaymentRazorPayResponse
//	@Failure		400		{object}	pb.MakePaymentRazorPayResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.MakePaymentRazorPayResponse
//	@Router			/user/post/donate/razorpay  [get]
func MakePaymentRazorPay(ctx *gin.Context, c pb.UserServiceClient) {
	payID, err := strconv.Atoi(ctx.Query("payid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pb.MakePaymentRazorPayResponse{
			Status:     http.StatusBadRequest,
			Response:   "Error with request",
			RazorId:    "",
			PaymentID:  int32(payID),
			FinalPrice: 0,
		})
		return
	}
	payIdBody := PayIdBody{PayId: payID}
	validator := validator.New()
	if err := validator.Struct(payIdBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.MakePaymentRazorPayResponse{
			Status:     http.StatusBadRequest,
			Response:   "Invalid data" + err.Error(),
			RazorId:    "",
			PaymentID:  int32(payID),
			FinalPrice: 0,
		})
		return
	}
	res, err := c.MakePaymentRazorPay(context.Background(), &pb.MakePaymentRazorPayRequest{Payid: int32(payID)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.MakePaymentRazorPayResponse{
			Status:     http.StatusBadGateway,
			Response:   "Error in internal server",
			RazorId:    "",
			PaymentID:  int32(payID),
			FinalPrice: 0,
		})
		return
	}
	// orderDetail, err := p.paymentUseCase.MakePaymentRazorPay(orderID, userID)
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
	// 	c.JSON(http.StatusInternalServerError, errorRes)
	// 	return
	// }
	log.Println("collected data ", res)
	ctx.HTML(http.StatusOK, "razorpay.html", res)
}
