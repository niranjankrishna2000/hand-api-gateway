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

// Admin get Campaign Details godoc
//
//	@Summary		Admin can get Campaign Details
//	@Description	Admin can get Campaign Details
//	@Tags			Admin Campaign
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"post id"
//	@Success		200	{object}	pb.CampaignDetailsResponse
//	@Router			/admin/campaigns/details  [get]
func CampaignDetails(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	IDstr := ctx.Query("id")
	ID, err := strconv.Atoi(IDstr)
	if err != nil || ID < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("post Id format invalid"))
		return
	}
	log.Println("Collected Data: ", ID)
	log.Println("Fetching Data...")

	res, err := c.CampaignDetails(context.Background(), &pb.CampaignDetailsRequest{Id: int32(ID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusOK, &res)
}
