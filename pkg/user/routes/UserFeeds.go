package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// Feeds godoc
//
//	@Summary		Feeds
//	@Description	User can see feeds
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			limit		query		string	false	"limit"
//	@Param			page		query		string	false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.UserFeedsResponse
//	@Failure		400			{object}	pb.UserFeedsResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.UserFeedsResponse
//	@Router			/user/feeds  [get]
func UserFeeds(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("starting User Feeds")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pb.UserFeedsResponse{
			Status:    http.StatusBadRequest,
			Response:  "Error with page",
			Posts: nil,
		})
		return
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.UserFeedsResponse{
			Status:    http.StatusBadRequest,
			Response:  "Error with limit",
			Posts: nil,
		})
		return
	}
	searchkey := ctx.Query("searchkey")
	pageBody := PageBody{Page: page, Limit: limit, Searchkey: searchkey}

	log.Println("Collected data : ", page, limit, searchkey)
	//note: ** need userid?
	validator := validator.New()
	if err := validator.Struct(pageBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.UserFeedsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Posts:    nil,
		})
		return
	}
	res, err := c.UserFeeds(context.Background(), &pb.UserFeedsRequest{Page: int32(page), Limit: int32(limit), Searchkey: searchkey})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.UserFeedsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Posts:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
