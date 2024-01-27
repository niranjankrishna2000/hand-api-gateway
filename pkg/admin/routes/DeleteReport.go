package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Admin Delete reported post godoc
//
// @Summary  Admin can Delete reported post
// @Description Admin can Delete reported post
// @Tags   Admin Reported
// @Accept   json
// @Produce  json
// @Param   postId query  string true "Post ID "
// @Success  200   {object} pb.DeleteReportResponse
// @Router   /admin/campaigns/reported/delete  [delete]
func DeleteReport(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	postIDstr := ctx.Query("postid")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil || postID < 0 {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"Status":   http.StatusBadRequest,
			"Response": "Id format invalid" + err.Error(),
			"Id":       postID,
		})
		return
	}

	log.Println("Collected data : ", postID)

	res, err := c.DeleteReport(context.Background(), &pb.DeleteReportRequest{Postid: int32(postID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)
	ctx.JSON(http.StatusOK, &res)
}
