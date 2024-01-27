package routes

import (
	"context"
	"errors"
	"hand/pkg/user/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User Donation History godoc
//
//	@Summary		User can see Donation History
//	@Description	User can see Donation History
//	@Tags			User Donations
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			limit		query		string	false	"limit"
//	@Param			page		query		string	false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.DonationHistoryResponse
//	@Router			/user/post/donate/history  [get]
func DonationHistory(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("starting User Donation history")
	pageStr := ctx.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid Page format"))
		return
	}

	limitStr := ctx.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid limit format"))
		return
	}

	searchkey := ctx.Query("searchkey")
	userId := ctx.GetInt64("userId")
	log.Println("Collected data : ", page, limit, searchkey, userId)
	res, err := c.DonationHistory(context.Background(), &pb.DonationHistoryRequest{
		Page:      int32(page),
		Limit:     int32(limit),
		Searchkey: searchkey,
		Userid:    int32(userId),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)

}
