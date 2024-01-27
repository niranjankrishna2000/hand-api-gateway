package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginWithOtpRequestBody struct {
	Phone string `json:"phone" validate:"required,len=10,number"`
}

// User Login with Otp godoc
//
// @Summary  User Login with otp
// @Description Forgot Password?User can login with otp here
// @Tags   User Auth
// @Accept   json
// @Produce  json
// @Param   b body  LoginWithOtpRequestBody true "User Login Data"
// @Success  200   {object} pb.LoginWithOtpResponse
// @Router   /forgot-password [post]
func LoginWithOtp(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginWithOtpRequestBody{}

	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid input",
		})
		return
	}

	validator := validator.New()
	if err := validator.Struct(b); err != nil {
		log.Println("Error while validating :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid Input",
		})
		return
	}

	res, err := c.LoginWithOtp(context.Background(), &pb.LoginWithOtpRequest{
		Phone: b.Phone,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error":  "Error with internal server",
		})
		return	}

	ctx.JSON(http.StatusOK, &res)
}
