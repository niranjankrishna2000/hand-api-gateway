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

type CategoryStatsBody struct {
	Limit int `json:"limit" validate:"min=1,max=50,number"`
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
//	@Param			limit	query		int	false	"limit"
//	@Param			page	query		int	false	"Page number"
//	@Success		200		{object}	pb.CategoryStatsResponse
//	@Failure		200		{object}	pb.CategoryStatsResponse
//	@Failure		200		{object}	pb.CategoryStatsResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Router			/admin/dashboard/category  [get]
func CategoryStats(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating CategoryStats...")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	categoryStatsBody := CategoryStatsBody{Page: page, Limit: limit}

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
		Page:  int32(categoryStatsBody.Page),
		Limit: int32(categoryStatsBody.Limit)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.CategoryStatsResponse{
			Status:     http.StatusBadGateway,
			Response:   "Error in internal server",
			Categories: nil,
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
