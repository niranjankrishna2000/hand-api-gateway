package routes

import (
	"context"
	"errors"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User Clear History godoc
//
// @Summary  User can Clear History
// @Description User can Clear History
// @Tags   User Donations
// @Accept   json
// @Produce  json
// @Success  200   {object} pb.ClearHistoryResponse
// @Router   /user/post/donate/history/clear  [delete]
func ClearHistory(ctx *gin.Context, c pb.UserServiceClient) {
	userId := ctx.GetInt64("userId")
	if userId <= 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("user Id invalid"))
		return
	}
	log.Println("Collected data : ", userId)
	res, err := c.ClearHistory(context.Background(), &pb.ClearHistoryRequest{

		Userid: int32(userId),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)

}
