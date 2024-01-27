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

// Admin User List godoc
//
//	@Summary		Admin can see User List
//	@Description	Admin can see User List
//	@Tags			Admin Users
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			page		query		string	false	"page"
//	@Param			limit		query		string	false	"limit"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.UserListResponse
//	@Router			/admin/users/list  [get]
func UserList(ctx *gin.Context, c pb.AdminServiceClient, usvc user.AuthServiceClient) {
	log.Println("Initiating AdminDashboard...")

	pageStr := ctx.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	limitStr := ctx.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	searchkey := ctx.Query("searchkey")
	log.Println("Collected data : ", page, limit, searchkey)
	log.Println("Fetching Data...")
	res, err := usvc.UserList(context.Background(), &user.UserListRequest{Page: int32(page),Limit: int32(limit),Searchkey: searchkey})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
