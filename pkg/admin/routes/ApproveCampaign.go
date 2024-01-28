package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ApproveCampaignBody struct {
	PostId int `json:"postId" validate:"required,min=1,max=999,number"`
}

// Admin Approve Campaign godoc
//
//	@Summary		Approve Campaigns
//	@Description	Admin can Approve Campaign after verifying it
//	@Tags			Admin Campaign
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			ApproveCampaignBody	body		ApproveCampaignBody	true	"Post ID "
//	@Success		200					{object}	pb.ApproveCampaignResponse
//	@Failure		400					{object}	pb.ApproveCampaignResponse
//	@Failure		502					{object}	pb.ApproveCampaignResponse
//	@Router			/admin/campaigns/approve  [patch]
func ApproveCampaign(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating ApproveCampaign...")

	approveCampaignBody := ApproveCampaignBody{}

	if err := ctx.BindJSON(&approveCampaignBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.ApproveCampaignResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Post:     nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(approveCampaignBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest,pb.ApproveCampaignResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid Post ID",
			Post:     nil,
		})
		return
	}
	log.Println("Collected Data: ", approveCampaignBody)
	res, err := c.ApproveCampaign(context.Background(), &pb.ApproveCampaignRequest{
		Id: int32(approveCampaignBody.PostId),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.ApproveCampaignResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusOK, &res)
}
