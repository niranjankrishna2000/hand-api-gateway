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

// User Notification Detail godoc
//
// @Summary  User can see Notification Details
// @Description User can see Notification Details
// @Tags   User Notifications
// @Accept   json
// @Produce  json
// @Param			notificationId	query		string	false	"Notification Id"
// @Success  200   {object} pb.CommentPostResponse
// @Router   /user/notifications/details  [get]
func NotificationDetail(ctx *gin.Context, c pb.UserServiceClient) {
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
	res, err := c.NotificationDetail(context.Background(), &pb.NotificationDetailsRequest{
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
