package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/admin/pb"
	user "hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

// Admin Change User Permission godoc
//
//	@Summary		Admin can Change User Permission
//	@Description	Admin can Change User Permission
//	@Tags			Admin Users
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"user id "
//	@Success		200	{object}	pb.ChangeUserPermissionResponse
//	@Router			/admin/users/changepermission  [patch]
func ChangeUserPermission(ctx *gin.Context, c pb.AdminServiceClient, usvc user.AuthServiceClient) {
	log.Println("Initiating AdminDashboard...")

	userIDstr := ctx.Query("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := usvc.ChangeUserPermission(context.Background(), &user.ChangeUserPermissionRequest{Id: int32(userID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
