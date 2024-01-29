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

type CampaignDetailsBody struct {
	PostId int `json:"postId" validate:"required,min=1,max=999,number"`
}

// Admin get Campaign Details godoc
//
//	@Summary		Campaign Details
//	@Description	Admin can see Campaign Details
//	@Tags			Admin Campaign
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			postid	query		string	true	"Post ID "
//	@Success		200		{object}	pb.CampaignDetailsResponse
//	@Failure		400		{object}	pb.CampaignDetailsResponse
//	@Failure		502		{object}	pb.CampaignDetailsResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Router			/admin/campaigns/details  [get]
func CampaignDetails(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating CampaignDetails...")

	postId, err := strconv.Atoi(ctx.Query("postid"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.CampaignDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with post Id",
			Post:     nil,
		})
		return
	}
	campaignDetailsBody := CampaignDetailsBody{PostId: postId}

	validator := validator.New()
	if err := validator.Struct(campaignDetailsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.CampaignDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid Post ID",
			Post:     nil,
		})
		return
	}

	res, err := c.CampaignDetails(context.Background(), &pb.CampaignDetailsRequest{Id: int32(campaignDetailsBody.PostId)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.CampaignDetailsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusOK, &res)
}
