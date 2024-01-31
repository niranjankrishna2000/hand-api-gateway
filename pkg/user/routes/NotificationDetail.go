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

type NotificationDetailBody struct {
	NotificationID int `json:"notificationId" validate:"min=1,max=999,number"`
}

// Notification Detail godoc
//
//	@Summary		Notification Details
//	@Description	User can see Notification Details
//	@Tags			User Notifications
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			notificationId	query		string	false	"Notification Id"
//	@Success		200				{object}	pb.NotificationDetailsResponse
//	@Failure		400				{object}	pb.NotificationDetailsResponse
//	@Failure		403				{string}	string	"You have not logged in"
//	@Failure		502				{object}	pb.NotificationDetailsResponse
//	@Router			/user/notifications/details  [get]
func NotificationDetail(ctx *gin.Context, c pb.UserServiceClient) {
	notificationID, err := strconv.Atoi(ctx.Query("notificationId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pb.NotificationDetailsResponse{
			Status:       http.StatusBadRequest,
			Response:     "Invalid Notification",
			Notification: nil,
		})
		return
	}
	notificationIdBody := NotificationIdBody{NotficationId: notificationID}
	userId := ctx.GetInt64("userId")
	log.Println("Sending Data: ", notificationID, userId)
	//validator
	validator := validator.New()
	if err := validator.Struct(notificationIdBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.NotificationDetailsResponse{
			Status:       http.StatusBadRequest,
			Response:     "Invalid data" + err.Error(),
			Notification: nil,
		})
		return
	}
	res, err := c.NotificationDetail(context.Background(), &pb.NotificationDetailsRequest{
		Userid:         int32(userId),
		Notificationid: int32(notificationID),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.NotificationDetailsResponse{
			Status:       http.StatusBadGateway,
			Response:     "Error in internal server",
			Notification: nil,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusOK, &res)
}
