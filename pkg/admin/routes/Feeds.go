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

type FeedsBody struct {
	Limit     int    `json:"limit" validate:"min=1,max=50,number"`
	Page      int    `json:"page" validate:"min=1,max=99,number"`
	Searchkey string `json:"searchkey"`
}

// Admin feeds godoc
//
//	@Summary		Feeds
//	@Description	Admin can see feeds
//	@Tags			Admin Feeds
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		string	false	"limit"
//	@Param			page		query		string	false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.FeedsResponse
//	@Failure		400			{object}	pb.FeedsResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.FeedsResponse
//	@Router			/admin/feeds  [get]
func Feeds(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating Feeds...")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	searchkey := ctx.Query("searchkey")
	feedsBody := FeedsBody{Page: page, Limit: limit, Searchkey: searchkey}

	validator := validator.New()
	if err := validator.Struct(feedsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.FeedsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Posts:    nil,
		})
		return
	}
	res, err := c.Feeds(context.Background(), &pb.FeedsRequest{
		Page:      int32(feedsBody.Page),
		Limit:     int32(feedsBody.Limit),
		Searchkey: feedsBody.Searchkey,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.FeedsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Posts:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
