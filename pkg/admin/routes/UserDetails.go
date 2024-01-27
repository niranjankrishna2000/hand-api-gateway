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

// Admin get User detail godoc
//
//	@Summary		Admin can get user detail
//	@Description	Admin can get user details
//	@Tags			Admin Users
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"user id "
//	@Success		200	{object}	pb.GetUserDetailsResponse
//	@Router			/admin/users/details  [get]
func UserDetails(ctx *gin.Context, c pb.AdminServiceClient, usvc user.AuthServiceClient) {
	log.Println("Initiating AdminDashboard...")

	userIDstr := ctx.Query("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println("Fetching Data...")

	res, err := usvc.GetUserDetails(context.Background(), &user.GetUserDetailsRequest{Userid: int32(userID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
