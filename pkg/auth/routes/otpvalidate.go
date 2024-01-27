package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OtpValidateRequestBody struct {
	Phone    string `json:"phone" validate:"required,len=10,number"`
	Otp      string `json:"otp" validate:"required,len=6,number"`
	Password string `json:"password" validate:"required,min=6,alphanum"`
	Confirm  string `json:"confirm" validate:"required,eqfield=Password"`
}

// User otp validation godoc
//
// @Summary  User Otp Validation
// @Description User can validate otp here
// @Tags   User Auth
// @Accept   json
// @Produce  json
// @Param   b body  OtpValidateRequestBody true "Validate otp Data"
// @Success  200   {object} pb.LoginResponse
// @Router   /otp-validate [patch]
func OtpValidate(ctx *gin.Context, c pb.AuthServiceClient) {
	otpBody := OtpValidateRequestBody{}

	if err := ctx.BindJSON(&otpBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid input",
		})
		return
	}

	validator := validator.New()
	if err := validator.Struct(otpBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid Input",
		})
		return
	}
	res, err := c.OtpValidate(context.Background(), &pb.OtpValidationRequest{
		Phone:    otpBody.Phone,
		Otp:      otpBody.Otp,
		Password: otpBody.Password,
		Confirm:  otpBody.Confirm,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Error with internal server",
		})
		return
	}

	ctx.Header("Authorization", "Bearer "+res.Token)
	//ctx.SetCookie("Authorization",res.Token,3600,"","",false,true)
	ctx.JSON(http.StatusOK, &res)
}
