package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

// User Donate Payment godoc
//
// @Summary  User can pay for donation
// @Description User can pay for donation
// @Tags   User Donations
// @Accept   json
// @Produce  json
// @Param   payid query  string true "pay id Data"
// @Success  200   {object} pb.MakePaymentRazorPayResponse
// @Router   /user/post/donate/razorpay  [get]
func MakePaymentRazorPay(ctx *gin.Context, c pb.UserServiceClient) {
	payIDstr := ctx.Query("payid")
	payID, err := strconv.Atoi(payIDstr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.MakePaymentRazorPay(context.Background(), &pb.MakePaymentRazorPayRequest{Payid: int32(payID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	// orderDetail, err := p.paymentUseCase.MakePaymentRazorPay(orderID, userID)
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
	// 	c.JSON(http.StatusInternalServerError, errorRes)
	// 	return
	// }
		log.Println("collected data ",res)
	ctx.HTML(http.StatusOK, "razorpay.html", res)
}
