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
//	@Param			UserListBody	body		UserListBody	true	"Page Details and Searchkey"
//	@Success		200				{object}	pb.UserListResponse
//	@Failure		400				{object}	pb.UserListResponse
//	@Failure		502				{object}	pb.UserListResponse
//	@Router			/admin/users/list  [get]
func UserList(ctx *gin.Context, c pb.AdminServiceClient, usvc user.AuthServiceClient) {
	log.Println("Initiating AdminDashboard...")

	userListBody := UserListBody{}

	if err := ctx.BindJSON(&userListBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.UserStatsResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Users:    nil,
		})
		return
	}
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
