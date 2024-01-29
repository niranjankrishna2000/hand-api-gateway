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

type NotificationIdBody struct {
	NotficationId int `json:"notificationId" validate:"required,min=1,max=999,number"`
}

// Notification Delete godoc
//
//	@Summary		Delete Notification
//	@Description	User can delete Notification
//	@Security		api_key
//	@Tags			User Notifications
//	@Accept			json
//	@Produce		json
//	@Param			notificationId	query		string	false	"Notification Id"
//	@Success		200				{object}	pb.DeleteNotificationResponse
//	@Failure		400				{object}	pb.DeleteNotificationResponse
//	@Failure		403				{string}	string	"You have not logged in"
//	@Failure		502				{object}	pb.DeleteNotificationResponse
//	@Router			/user/notifications/delete  [delete]
func DeleteNotification(ctx *gin.Context, c pb.UserServiceClient) {
	notificationID, err := strconv.Atoi(ctx.Query("notificationId"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteNotificationResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
		})
		return
	}
	userId := ctx.GetInt64("userId")

	log.Println("Sending Data: ", notificationID, userId)
	notificationIdBody := NotificationIdBody{NotficationId: notificationID}

	validator := validator.New()
	if err := validator.Struct(notificationIdBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteNotificationResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
		})
		return
	}
	res, err := c.DeleteNotification(context.Background(), &pb.DeleteNotificationRequest{
		Userid:         int32(userId),
		Notificationid: int32(notificationID),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeleteNotificationResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusOK, &res)
}
