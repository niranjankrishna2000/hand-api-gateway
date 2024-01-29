package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// edit
type UserStatsBody struct {
	Limit int `json:"limit" validate:"min=1,max=99,number"`
	Page  int `json:"page" validate:"min=1,max=99,number"`
}

// User Stats godoc
//
//	@Summary		Top Users
//	@Description	Admin can see User toplist
//	@Tags			Admin Dashboard
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			UserStatsBody	body		UserStatsBody	true	"Page details"
//	@Success		200				{object}	pb.UserStatsResponse
//	@Failure		400				{object}	pb.UserStatsResponse
//	@Failure		403				{string}	string	"You have not logged in"
//	@Failure		502				{object}	pb.UserStatsResponse
//	@Router			/admin/dashboard/User  [get]
func UserStats(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating UserStats...")

	userStatsBody := UserStatsBody{}

	if err := ctx.BindJSON(&userStatsBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.UserStatsResponse{
			Status:     http.StatusBadRequest,
			Response:   "Error with request",
			Users: nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(userStatsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.UserStatsResponse{
			Status:     http.StatusBadRequest,
			Response:   "Invalid data" + err.Error(),
			Users: nil,
		})
		return
	}
	res, err := c.UserStats(context.Background(), &pb.UserStatsRequest{})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.UserStatsResponse{
			Status:     http.StatusBadGateway,
			Response:   "Error in internal server",
			Users: nil,
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
