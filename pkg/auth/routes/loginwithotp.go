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
//	@Summary		User Login with otp
//	@Description	Forgot Password?User can login with otp here
//	@Tags			User Auth
//	@Accept			json
//	@Produce		json
//	@Param			LoginWithOtpRequestBody	body		LoginWithOtpRequestBody	true	"User Login Data"
//	@Success		200						{object}	pb.LoginWithOtpResponse
//	@Failure		400						{object}	pb.LoginWithOtpResponse
//	@Failure		502						{object}	pb.LoginWithOtpResponse
//	@Router			/forgot-password [post]
func LoginWithOtp(ctx *gin.Context, c pb.AuthServiceClient) {
	LoginWithOtpRequestBody := LoginWithOtpRequestBody{}

	if err := ctx.BindJSON(&LoginWithOtpRequestBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.LoginWithOtpResponse{
			Status: http.StatusBadRequest,
			Error: "Error with request",
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(LoginWithOtpRequestBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest,pb.LoginWithOtpResponse{
			Status: http.StatusBadRequest,
			Error: "Invalid data"+err.Error(),
		})
		return
	}

	res, err := c.LoginWithOtp(context.Background(), &pb.LoginWithOtpRequest{
		Phone: LoginWithOtpRequestBody.Phone,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.LoginWithOtpResponse{
			Status: http.StatusBadRequest,
			Error: "Error with internal server",
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
