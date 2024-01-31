package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/admin/pb"
	user "hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserListBody struct {
	Limit     int    `json:"limit" validate:"min=1,max=99,number"`
	Page      int    `json:"page" validate:"min=1,max=99,number"`
	Searchkey string `json:"searchkey"`
}

// User List godoc
//
//	@Summary		User List
//	@Description	Admin can see User List
//	@Tags			Admin Users
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		string	false	"limit"
//	@Param			page		query		string	false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.UserListResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		400			{object}	pb.UserListResponse
//	@Failure		502			{object}	pb.UserListResponse
//	@Router			/admin/users/list  [get]
func UserList(ctx *gin.Context, c pb.AdminServiceClient, usvc user.AuthServiceClient) {
	log.Println("Initiating UserList...")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	searchkey := ctx.Query("searchkey")
	userListBody := UserListBody{Page: page, Limit: limit, Searchkey: searchkey}

	validator := validator.New()
	if err := validator.Struct(userListBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.UserStatsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Users:    nil,
		})
		return
	}
	res, err := usvc.UserList(context.Background(), &user.UserListRequest{Page: int32(userListBody.Page), Limit: int32(userListBody.Limit), Searchkey: userListBody.Searchkey})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.UserStatsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Users:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
