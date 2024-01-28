package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ReportedListBody struct {
	Limit     int    `json:"limit" validate:"min=1,max=99,number"`
	Page      int    `json:"page" validate:"min=1,max=99,number"`
	Searchkey string `json:"searchkey"`
}

// Admin Reported List godoc
//
//	@Summary		Reported posts
//	@Description	Admin can see reported posts
//	@Tags			Admin Reported
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			ReportedListBody	body		ReportedListBody	true	"limit"
//	@Success		200					{object}	pb.ReportedListResponse
//	@Failure		400					{object}	pb.ReportedListResponse
//	@Failure		502					{object}	pb.ReportedListResponse
//	@Router			/admin/campaigns/reported  [get]
func ReportedList(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	reportedListBody := ReportedListBody{}

	if err := ctx.BindJSON(&reportedListBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.ReportedListResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Post:     nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(reportedListBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.ReportedListResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     nil,
		})
		return
	}
	res, err := c.ReportedList(context.Background(), &pb.ReportedListRequest{
		Page: int32(reportedListBody.Page),
		Limit: int32(reportedListBody.Limit), 
		Searchkey: reportedListBody.Searchkey})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.ReportedListResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
