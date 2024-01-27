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

// User Notifications godoc
//
//	@Summary		User can see Notifications
//	@Description	User can see Notifications
//	@Tags			User Notifications
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		string	false	"limit"
//	@Param			page	query		string	false	"Page number"
//	@Success		200		{object}	pb.NotificationResponse
//	@Router			/user/notifications  [get]
func Notifications(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("starting User Donation history")
	pageStr := ctx.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid Page format"))
		return
	}

	limitStr := ctx.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid limit format"))
		return
	}
	userId := ctx.GetInt64("userId")
	log.Println("Collected data : ", page, limit, userId)

	res, err := c.Notifications(context.Background(), &pb.NotificationRequest{
		Page:      int32(page),
		Limit:     int32(limit),
		Userid:    int32(userId),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
