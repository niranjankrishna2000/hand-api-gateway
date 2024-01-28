package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CampaignRequestListBody struct {
	Limit     int    `json:"limit" validate:"min=1,max=99,number"`
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
//	@Param			CampaignRequestListBody	body		CampaignRequestListBody	true	"Page Details and Searchkey "
//	@Success		200						{object}	pb.CampaignRequestListResponse
//	@Failure		400						{object}	pb.CampaignRequestListResponse
//	@Success		502						{object}	pb.CampaignRequestListResponse
//	@Router			/admin/campaigns/requestlist  [get]
func CampaignRequestList(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	campaignRequestListBody := CampaignRequestListBody{}

	if err := ctx.BindJSON(&campaignRequestListBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.CampaignRequestListResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Post:     nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(campaignRequestListBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.CampaignRequestListResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data"+err.Error(),
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
