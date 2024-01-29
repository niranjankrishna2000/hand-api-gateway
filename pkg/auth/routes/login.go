package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// User Login godoc
//
//	@Summary		User Login
//	@Description	User can login here
//	@Tags			User Auth
//	@Accept			json
//	@Produce		json
//	@Param			b	body		LoginRequestBody	true	"User Login Data"
//	@Success		200	{object}	pb.LoginResponse
//	@Router			/login [post]
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	LoginRequestBody := LoginRequestBody{}

	if err := ctx.BindJSON(&LoginRequestBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error: "Error with request",
			User: nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(LoginRequestBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest,pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error: "Invalid data"+err.Error(),
			User: nil,
		})
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    LoginRequestBody.Email,
		Password: LoginRequestBody.Password,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error: "Error with internal server",
			User: nil,
		})
		return
	}

	ctx.Header("Authorization", "Bearer "+res.Token)
	//ctx.SetCookie("Authorization", res.Token, 3600, "", "", false, true)
	ctx.JSON(http.StatusOK, &res)
}
