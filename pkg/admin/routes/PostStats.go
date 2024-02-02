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

type PostStatsBody struct {
	Limit int `json:"limit" validate:"min=1,max=50,number"`
	Page  int `json:"page" validate:"min=1,max=99,number"`
}

// Post Stats godoc
//
//	@Summary		Top Posts
//	@Description	Admin can see Post toplist
//	@Tags			Admin Dashboard
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		int	false	"limit"
//	@Param			page	query		int	false	"Page number"
//	@Success		200		{object}	pb.PostStatsResponse
//	@Failure		400		{object}	pb.PostStatsResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.PostStatsResponse
//	@Router			/admin/dashboard/posts  [get]
func PostStats(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating PostStats...")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	postStatsBody := PostStatsBody{Page: page, Limit: limit}

	validator := validator.New()
	if err := validator.Struct(postStatsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.PostStatsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Posts:    nil,
		})
		return
	}
	res, err := c.PostStats(context.Background(), &pb.PostStatsRequest{
		Limit: int32(postStatsBody.Limit),
		Page:  int32(postStatsBody.Limit),
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.PostStatsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Posts:    nil,
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
