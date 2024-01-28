package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CategoryStatsBody struct {
	Limit int `json:"limit" validate:"min=1,max=99,number"`
	Page  int `json:"page" validate:"min=1,max=99,number"`
}

// Admin Category Stats godoc
//
//	@Summary		Top Categories
//	@Description	Admin can see top Categories
//	@Tags			Admin Dashboard
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			CategoryStatsBody	body		CategoryStatsBody	true	"Page Details"
//	@Success		200					{object}	pb.CategoryStatsResponse
//	@Failure		200					{object}	pb.CategoryStatsResponse
//	@Failure		200					{object}	pb.CategoryStatsResponse
//	@Router			/admin/dashboard/category  [get]
func CategoryStats(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating CategoryStats...")

	categoryStatsBody := CategoryStatsBody{}

	if err := ctx.BindJSON(&categoryStatsBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.CategoryStatsResponse{
			Status:     http.StatusBadRequest,
			Response:   "Error with request",
			Categories: nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(categoryStatsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.CategoryStatsResponse{
			Status:     http.StatusBadRequest,
			Response:   "Invalid data" + err.Error(),
			Categories: nil,
		})
		return
	}
	res, err := c.CategoryStats(context.Background(), &pb.CategoryStatsRequest{
		Page: int32(categoryStatsBody.Page), 
		Limit: int32(categoryStatsBody.Limit)})

		if err != nil {
			log.Println("Error with internal server :", err)
			ctx.JSON(http.StatusBadGateway, pb.CategoryStatsResponse{
				Status:   http.StatusBadGateway,
				Response: "Error in internal server",
				Categories:     nil,
			})
			return
		}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
