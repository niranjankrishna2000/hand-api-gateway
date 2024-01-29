package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserIdBody struct {
	UserId int64 `json:"userId" validate:"required,min=1,max=999,number"`
}

// Clear Notification godoc
//
//	@Summary		Clear Notification
//	@Description	User can Clear Notifications
//	@Tags			User Notifications
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	pb.ClearNotificationResponse
//	@Failure		400	{object}	pb.ClearNotificationResponse
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		502	{object}	pb.ClearNotificationResponse
//	@Router			/user/notifications/clear  [delete]
func ClearNotification(ctx *gin.Context, c pb.UserServiceClient) {
	userId := UserIdBody{UserId: ctx.GetInt64("userId")}
	validator := validator.New()
	if err := validator.Struct(userId); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.ClearNotificationResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid user ID",
		})
		return
	}
	log.Println("Collected data : ", userId)
	res, err := c.ClearNotification(context.Background(), &pb.ClearNotificationRequest{

		Userid: int32(userId.UserId),
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.ClearNotificationResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
