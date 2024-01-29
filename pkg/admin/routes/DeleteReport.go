package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type DeleteReportBody struct {
	PostId int `json:"postId" validate:"required,min=1,max=999,number"`
}

// Delete reported post godoc
//
//	@Summary		Delete reported post
//	@Description	Admin can Delete reported post
//	@Tags			Admin Reported
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			DeleteReportBody	body		DeleteReportBody	true	"Post ID "
//	@Success		200					{object}	pb.DeleteReportResponse
//	@Failure		400					{object}	pb.DeleteReportResponse
//	@Failure		403					{string}	string	"You have not logged in"
//	@Failure		502					{object}	pb.DeleteReportResponse
//	@Router			/admin/campaigns/reported/delete  [delete]
func DeleteReport(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating DeleteReport...")

	deleteReportBody := DeleteReportBody{}

	if err := ctx.BindJSON(&deleteReportBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteReportResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(deleteReportBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteReportResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid Post ID",
		})
		return
	}
	res, err := c.DeleteReport(context.Background(), &pb.DeleteReportRequest{Postid: int32(deleteReportBody.PostId)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeleteReportResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
		})
		return
	}
	log.Println("Recieved data : ", res)
	ctx.JSON(http.StatusOK, &res)
}
