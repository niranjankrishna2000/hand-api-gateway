package routes

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
)

// Admin Delete post godoc
//
// @Summary  Admin can delete post
// @Description Admin can delete post
// @Tags   Admin Feeds
// @Accept   json
// @Produce  json
// @Param			id	query		string	true	"id"
// @Success  200   {object} pb.DeletePostResponse
// @Router   /admin/post/delete  [delete]
func DeletePost(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("id format invalid"))
		return
	}
	log.Println("Collected data : ", id)

	res, err := c.DeletePost(context.Background(), &pb.DeletePostRequest{Id: int32(id)})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)
	ctx.JSON(http.StatusOK, &res)
}
