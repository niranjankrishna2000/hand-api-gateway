package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ReportDetailsBody struct {
	PostId int `json:"postId" validate:"required,min=1,max=999,number"`
}

// Reported Post Details godoc
//
//	@Summary		Details of reported post
//	@Description	Admin can see details of reported post
//	@Tags			Admin Reported
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			postid	query		string	true	"Post ID "
//	@Success		200		{object}	pb.ReportDetailsResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		400		{object}	pb.ReportDetailsResponse
//	@Failure		502		{object}	pb.ReportDetailsResponse
//	@Router			/admin/campaigns/reported/details  [get]
func ReportDetails(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating ReportDetails...")
	postId, err := strconv.Atoi(ctx.Query("postid"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.PostDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with post Id",
			Post:     nil,
		})
		return
	}
	reportDetailsBody := ReportDetailsBody{PostId: postId}

	validator := validator.New()
	if err := validator.Struct(reportDetailsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.PostDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid Post ID",
			Post:     nil,
		})
		return
	}
	res, err := c.ReportDetails(context.Background(), &pb.ReportDetailsRequest{Postid: int32(reportDetailsBody.PostId)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.PostDetailsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Data recieved: ", res)

	ctx.JSON(http.StatusOK, &res)
}
