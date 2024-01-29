package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// User Notifications godoc
//
//	@Summary		Notifications
//	@Description	User can see Notifications
//	@Tags			User Notifications
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		string	false	"limit"
//	@Param			page	query		string	false	"Page number"
//	@Success		200		{object}	pb.NotificationResponse
//	@Failure		400		{object}	pb.NotificationResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.NotificationResponse
//	@Router			/user/notifications  [get]
func Notifications(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("starting User Donation history")

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page=1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit=10
	}
	pageBody := PageBody{Page: page, Limit: limit}

	userId := ctx.GetInt64("userId")
	log.Println("Collected data : ", page, limit, userId)
	validator := validator.New()
	if err := validator.Struct(pageBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.NotificationResponse{
			Status:        http.StatusBadRequest,
			Response:      "Invalid data" + err.Error(),
			Notifications: nil,
		})
		return
	}
	res, err := c.Notifications(context.Background(), &pb.NotificationRequest{
		Page:   int32(page),
		Limit:  int32(limit),
		Userid: int32(userId),
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.NotificationResponse{
			Status:        http.StatusBadGateway,
			Response:      "Error in internal server",
			Notifications: nil,
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
