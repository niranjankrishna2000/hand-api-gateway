package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
)

// Admin Reject Campaign godoc
//
//	@Summary		Admin can Reject Campaign
//	@Description	Admin can Reject Campaign
//	@Tags			Admin Campaign
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"user id "
//	@Success		200	{object}	pb.RejectCampaignResponse
//	@Router			/admin/campaigns/reject  [patch]
func RejectCampaign(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	userIDstr := ctx.Query("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.RejectCampaign(context.Background(), &pb.RejectCampaignRequest{Id: int32(userID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
