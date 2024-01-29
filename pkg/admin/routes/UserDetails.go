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

type UserDetailsBody struct {
	UserId int `json:"userId" validate:"required,min=1,max=999,number"`
}

// User detail godoc
//
//	@Summary		User detail
//	@Description	Admin can get user details
//	@Tags			Admin Users
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			userid	query		string	true	"User ID "
//	@Success		200		{object}	pb.GetUserDetailsResponse
//	@Failure		400		{object}	pb.GetUserDetailsResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.GetUserDetailsResponse
//	@Router			/admin/users/details  [get]
func UserDetails(ctx *gin.Context, c pb.AdminServiceClient, usvc user.AuthServiceClient) {
	log.Println("Initiating UserDetails...")
	userId, err := strconv.Atoi(ctx.Query("userid"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, user.GetUserDetailsResponse{
			Status:   http.StatusBadRequest,
			Error: "Error with post Id",
			User:     nil,
		})
		return
	}
	userDetailsBody := UserDetailsBody{UserId: userId}

	validator := validator.New()
	if err := validator.Struct(userDetailsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, user.GetUserDetailsResponse{
			Status: http.StatusBadRequest,
			Error:  "Invalid user ID",
			User:   nil,
		})
		return
	}
	res, err := usvc.GetUserDetails(context.Background(), &user.GetUserDetailsRequest{Userid: int32(userDetailsBody.UserId)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, user.GetUserDetailsResponse{
			Status: http.StatusBadGateway,
			Error:  "Error in internal server",
			User:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
