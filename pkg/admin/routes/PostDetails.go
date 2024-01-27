package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
)

// Admin get Post detail godoc
//
// @Summary  Admin can get post detail
// @Description Admin can get post detail
// @Tags   Admin Feeds
// @Accept   json
// @Produce  json
// @Param   id query  string true "post id Data"
// @Success  200   {object} pb.PostDetailsResponse
// @Router   /admin/post/details  [get]
func PostDetails(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	postIDstr := ctx.Query("id")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println("Data collected: ",postID)
	log.Println("Fetching Data...")

	res, err := c.PostDetails(context.Background(), &pb.PostDetailsRequest{PostID: int32(postID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Data recieved: ",res)

	ctx.JSON(http.StatusOK, &res)
}
