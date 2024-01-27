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
	loginBody := LoginRequestBody{}

	if err := ctx.BindJSON(&loginBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid input",
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(loginBody); err != nil {
		log.Println("Error while validating :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid Input",
		})
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    loginBody.Email,
		Password: loginBody.Password,
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
	//ctx.SetCookie("Authorization", res.Token, 3600, "", "", false, true)
	ctx.JSON(http.StatusOK, &res)
}
