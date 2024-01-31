package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/admin/pb"
	user "hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ChangeUserPermissionBody struct {
	UserId int `json:"userId" validate:"required,min=1,max=999,number"`
}

// Admin Change User Permission godoc
//
//	@Summary		Change User Permission
//	@Description	Admin can Change User Permission
//	@Tags			Admin Users
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			ChangeUserPermissionBody	body		ChangeUserPermissionBody	true	"User ID "
//	@Success		200							{object}	pb.ChangeUserPermissionResponse
//	@Failure		400							{object}	pb.ChangeUserPermissionResponse
//	@Failure		403							{string}	string	"You have not logged in"
//	@Failure		502							{object}	pb.ChangeUserPermissionResponse
//	@Router			/admin/users/changepermission  [patch]
func ChangeUserPermission(ctx *gin.Context, c pb.AdminServiceClient, usvc user.AuthServiceClient) {
	log.Println("Initiating AdminDashboard...")

	changeUserPermissionBody := ChangeUserPermissionBody{}

	if err := ctx.BindJSON(&changeUserPermissionBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, user.ChangeUserPermissionResponse{
			Status: http.StatusBadRequest,
			Error:  "Error with request",
			User:   nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(changeUserPermissionBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, user.ChangeUserPermissionResponse{
			Status: http.StatusBadRequest,
			Error:  "Invalid User ID",
			User:   nil,
		})
		return
	}
	res, err := usvc.ChangeUserPermission(context.Background(), &user.ChangeUserPermissionRequest{
		Id: int32(changeUserPermissionBody.UserId),
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, user.ChangeUserPermissionResponse{
			Status: http.StatusBadGateway,
			Error:  "Error in internal server",
			User:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
