package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// My Campaigns godoc
//
//	@Summary		My Campaigns
//	@Description	User can see their campaigns
//	@Tags			User Profile
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			limit		query		int		false	"limit"
//	@Param			page		query		int		false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.GetMyCampaignsResponse
//	@Failure		400			{object}	pb.GetMyCampaignsResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.GetMyCampaignsResponse
//	@Router			/user/profile/my-campaigns  [get]
func GetMyCampaigns(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("starting User Feeds")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	searchkey := ctx.Query("searchkey")
	pageBody := PageBody{Page: page, Limit: limit, Searchkey: searchkey}

	log.Println("Collected data : ", page, limit, searchkey)
	//note: ** need userid?
	validator := validator.New()
	if err := validator.Struct(pageBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.GetMyCampaignsResponse{
			Status:         http.StatusBadRequest,
			Response:       "Invalid data :" + err.Error(),
			Posts:          []*pb.Post{},
		})
		return
	}
	log.Println(int32(ctx.GetInt64("userid")))
	res, err := c.GetMyCampaigns(context.Background(), &pb.GetMyCampaignsRequest{Page: int32(page), Limit: int32(limit), Searchkey: searchkey, UserId: int32(ctx.GetInt64("userid"))})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.GetMyCampaignsResponse{
			Status:         http.StatusBadGateway,
			Response:       err.Error(),
			Posts:          []*pb.Post{},
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
