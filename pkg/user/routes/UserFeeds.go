package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

// User feeds godoc
//
//	@Summary		User can see feeds
//	@Description	User can see feeds
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			limit		query		string	false	"limit"
//	@Param			page		query		string	false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.UserFeedsResponse
//	@Router			/user/feeds  [get]
func UserFeeds(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("starting User Feeds")
	pageStr := ctx.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limitStr := ctx.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	searchkey := ctx.Query("searchkey")
	log.Println("Collected data : ", page, limit, searchkey)
	//note: ** need userid?
	res, err := c.UserFeeds(context.Background(), &pb.UserFeedsRequest{Page: int32(page), Limit: int32(limit), Searchkey: searchkey})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
