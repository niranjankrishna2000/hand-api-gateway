package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminLoginRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Admin Login godoc
//
//	@Summary		Admin Login
//	@Description	Admin can login here
//	@Tags			Admin Auth
//	@Accept			json
//	@Produce		json
//	@Param			AdminLoginRequestBody	body		AdminLoginRequestBody	true	"Admin Login Data"
//	@Success		200						{object}	pb.AdminLoginResponse
//	@Failure		400						{object}	pb.AdminLoginResponse
//	@Failure		502						{object}	pb.AdminLoginResponse
//	@Router			/adminlogin [post]
func AdminLogin(ctx *gin.Context, c pb.AuthServiceClient) {
	AdminLoginRequestBody := AdminLoginRequestBody{}

	if err := ctx.BindJSON(&AdminLoginRequestBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.AdminLoginResponse{
			Status: http.StatusBadRequest,
			Error:  "Error with request",
			Token:  "",
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(AdminLoginRequestBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.AdminLoginResponse{
			Status: http.StatusBadRequest,
			Error:  "Invalid data" + err.Error(),
			Token:  "",
		})
		return
	}
	res, err := c.AdminLogin(context.Background(), &pb.AdminLoginRequest{
		Email:    AdminLoginRequestBody.Email,
		Password: AdminLoginRequestBody.Password,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.AdminLoginResponse{
			Status: http.StatusBadRequest,
			Error:  "Error with internal server",
			Token:  "",
		})
		return
	}

	ctx.Header("Authorization", "Bearer "+res.Token)
	//ctx.SetCookie("Authorization", res.Token, 3600, "", "", false, true)
	ctx.JSON(http.StatusOK, &res)
}
