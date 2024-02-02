package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CampaignRequestListBody struct {
	Limit     int    `json:"limit" validate:"min=1,max=50,number"`
	Page      int    `json:"page" validate:"min=1,max=99,number"`
	Searchkey string `json:"searchkey"`
}

// Admin campaign list godoc
//
//	@Summary		Campaign Requests
//	@Description	Admin can see campaogn requests
//	@Tags			Admin Campaign
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		int		false	"limit"
//	@Param			page		query		int		false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.CampaignRequestListResponse
//	@Failure		400			{object}	pb.CampaignRequestListResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.CampaignRequestListResponse
//	@Router			/admin/campaigns/requestlist  [get]
func CampaignRequestList(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating CampaignRequestList...")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	searchkey := ctx.Query("searchkey")

	campaignRequestListBody := CampaignRequestListBody{Page: page, Limit: limit, Searchkey: searchkey}

	validator := validator.New()
	if err := validator.Struct(campaignRequestListBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.CampaignRequestListResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     nil,
		})
		return
	}
	res, err := c.CampaignRequestList(context.Background(), &pb.CampaignRequestListRequest{
		Page:      int32(campaignRequestListBody.Page),
		Limit:     int32(campaignRequestListBody.Limit),
		Searchkey: campaignRequestListBody.Searchkey,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.CampaignRequestListResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
