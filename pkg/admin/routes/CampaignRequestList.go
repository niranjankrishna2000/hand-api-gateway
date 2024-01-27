package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
)

// Admin campaign list godoc
//
//	@Summary		Admin can see campaign request
//	@Description	Admin can see campaogn requests
//	@Tags			Admin Campaign
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		string	false	"limit"
//	@Param			page		query		string	false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.CampaignRequestListResponse
//	@Router			/admin/campaigns/requestlist  [get]
func CampaignRequestList(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	pageStr := ctx.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	limitStr := ctx.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	searchkey := ctx.Query("searchkey")
	log.Println("Collected data : ", page, limit, searchkey)
	log.Println("Fetching Data...")

	res, err := c.CampaignRequestList(context.Background(), &pb.CampaignRequestListRequest{Page: int32(page), Limit: int32(limit), Searchkey: searchkey})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
