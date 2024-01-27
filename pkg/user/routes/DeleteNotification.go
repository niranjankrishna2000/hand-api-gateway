package routes

import (
	"context"
	"errors"
	"hand/pkg/user/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// User Notification Delete godoc
//
//	@Summary		User can Delete Notification 
//	@Description	User can delete Notification 
//	@Security		api_key
//	@Tags			User Notifications
//	@Accept			json
//	@Produce		json
//	@Param			notificationId	query		string	false	"Notification Id"
//	@Success		200				{object}	pb.CommentPostResponse
//	@Router			/user/notifications/delete  [delete]
func DeleteNotification(ctx *gin.Context, c pb.UserServiceClient) {
	notificationIDstr := ctx.Query("notificationId")
	notificationID, err := strconv.Atoi(notificationIDstr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid NotificationId format"))
		return
	}
	userId := ctx.GetInt64("userId")
	if userId < 1 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid UserID format"))
		return
	}
	log.Println("Sending Data: ", notificationID, userId)
	res, err := c.DeleteNotification(context.Background(), &pb.DeleteNotificationRequest{
		Userid:         int32(userId),
		Notificationid: int32(notificationID),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusOK, &res)
}
