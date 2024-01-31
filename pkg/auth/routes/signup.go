package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SignUpRequestBody struct {
	Name            string `json:"name" validate:"required,alpha,min=3,max=20"` //check what is alphaunicode
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required,len=10,number"`
	Password        string `json:"password" validate:"required,min=6,max=20,alphanum"`
	ConfirmPassword string `json:"confirmpassword" validate:"required,eqfield=Password"`
}

// Signup  godoc
//
//	@Summary		signup
//	@Description	Adding new user to the database
//	@Tags			User Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		SignUpRequestBody	true	"User Data"
//	@Success		200		{object}	pb.SignUpResponse
//	@Failure		400		{object}	pb.SignUpResponse
//	@Success		502		{object}	pb.SignUpResponse
//	@Router			/signup [post]
func SignUp(ctx *gin.Context, c pb.AuthServiceClient) {
	signupBody := SignUpRequestBody{}

	if err := ctx.BindJSON(&signupBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.SignUpResponse{
			Status: http.StatusBadRequest,
			Error:  "Error with request",
			User:   nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(signupBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.SignUpResponse{
			Status: http.StatusBadRequest,
			Error:  "Invalid data" + err.Error(),
			User:   nil,
		})
		return
	}
	res, err := c.SignUp(context.Background(), &pb.SignUpRequest{
		Name:            signupBody.Name,
		Email:           signupBody.Email,
		Phone:           signupBody.Phone,
		Password:        signupBody.Password,
		Confirmpassword: signupBody.ConfirmPassword,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.SignUpResponse{
			Status: http.StatusBadRequest,
			Error:  "Error with internal server",
			User:   nil,
		})
		return
	}

	ctx.JSON(int(res.Status), &res)
}
