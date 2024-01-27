package routes

import (
	"context"
	"errors"
	"hand/pkg/admin/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Admin Details reported post godoc
//
//	@Summary		Admin can see details reported post
//	@Description	Admin can see details reported post
//	@Tags			Admin Reported
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			postId	query		string	true	"Post ID "
//	@Success		200		{object}	pb.ReportDetailsResponse
//	@Router			/admin/campaigns/reported/details  [get]
func ReportDetails(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	postIDstr := ctx.Query("postId")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil || postID < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid Format"))
		return
	}
	log.Println("Data collected: ", postID)
	log.Println("Fetching Data...")

	res, err := c.ReportDetails(context.Background(), &pb.ReportDetailsRequest{Postid: int32(postID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Data recieved: ", res)

	ctx.JSON(http.StatusOK, &res)
}
