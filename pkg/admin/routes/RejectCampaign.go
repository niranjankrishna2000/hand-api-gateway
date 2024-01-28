package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type RejectCampaignBody struct {
	PostId int `json:"postId" validate:"required,min=1,max=999,number"`
}

// Reject Campaign godoc
//
//	@Summary		Reject Campaign
//	@Description	Admin can Reject Campaign
//	@Tags			Admin Campaign
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			RejectCampaignBody	body		RejectCampaignBody	true	"Post ID "
//	@Success		200					{object}	pb.RejectCampaignResponse
//	@Failure		400					{object}	pb.RejectCampaignResponse
//	@Failure		502					{object}	pb.RejectCampaignResponse
//	@Router			/admin/campaigns/reject  [patch]
func RejectCampaign(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating RejectCampaign...")

	rejectCampaignBody := RejectCampaignBody{}

	if err := ctx.BindJSON(&rejectCampaignBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.RejectCampaignResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Post:     nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(rejectCampaignBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.RejectCampaignResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid Post ID",
			Post:     nil,
		})
		return
	}
	res, err := c.RejectCampaign(context.Background(), &pb.RejectCampaignRequest{Id: int32(rejectCampaignBody.PostId)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.RejectCampaignResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
