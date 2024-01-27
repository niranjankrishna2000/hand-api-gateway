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
// @Summary  Admin Login
// @Description Admin can login here
// @Tags   Admin Auth
// @Accept   json
// @Produce  json
// @Param   b body  AdminLoginRequestBody true "Admin Login Data"
// @Success  200   {object} pb.AdminLoginResponse
// @Router   /adminlogin [post]
func AdminLogin(ctx *gin.Context, c pb.AuthServiceClient) {
	b := AdminLoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		log.Println("Error while fetching data :", err)
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
	res, err := c.AdminLogin(context.Background(), &pb.AdminLoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		log.Println("Error while validating :", err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid input",
		})
		return
	}

	ctx.Header("Authorization", "Bearer "+res.Token)
	//ctx.SetCookie("Authorization", res.Token, 3600, "", "", false, true)
	ctx.JSON(http.StatusOK, &res)
}
