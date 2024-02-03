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

type FeedsPageBody struct {
	Limit     int    `json:"limit" validate:"min=0,max=50,number"`
	Page      int    `json:"page" validate:"min=0,max=99,number"`
	Searchkey string `json:"searchkey" validate:"max=10,ascii"`
	PostType  int    `json:"type" validate:"min=0,max=3,number"`
	Category  int    `json:"category" validate:"min=0,max=5,number"`
}

// Feeds godoc
//
//	@Summary		Feeds
//	@Description	User can see feeds
//	@Description	Select Type: 1. Trending 2. Expired 3. Tax benefit 
//	@Description	select category: 0.All category 1. medical,2. child care,3. animal care, 4. Education,5. Memorial
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			limit		query		int		false	"limit"
//	@Param			page		query		int		false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Param			type		query		int		false	"Page number"
//	@Param			category	query		int		false	"category"
//	@Success		200			{object}	pb.UserFeedsResponse
//	@Failure		400			{object}	pb.UserFeedsResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.UserFeedsResponse
//	@Router			/user/feeds  [get]
func UserFeeds(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("starting User Feeds")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	posttype, err := strconv.Atoi(ctx.Query("type"))
	if err != nil {
		posttype = 1
	}
	category, err := strconv.Atoi(ctx.Query("category"))
	if err != nil {
		category = 0
	}
	searchkey := ctx.Query("searchkey")

	pageBody := FeedsPageBody{Page: page, Limit: limit, Searchkey: searchkey,PostType: posttype,Category: category}
	log.Println("Collected data : ", page, limit, searchkey, posttype, category)
	//note: ** need userid?
	validator := validator.New()
	if err := validator.Struct(pageBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.UserFeedsResponse{
			Status:         http.StatusBadRequest,
			Response:       "Invalid data :" + err.Error(),
			Posts:          []*pb.Post{},
			Categories:     []*pb.Category{},
			Successstories: []*pb.SuccesStory{},
		})
		return
	}
	userid:=ctx.GetInt64("userid")
	res, err := c.UserFeeds(context.Background(), &pb.UserFeedsRequest{Page: int32(page), Limit: int32(limit), Searchkey: searchkey, Category: int32(category), Type: int32(posttype), Userid: int32(userid)})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.UserFeedsResponse{
			Status:         http.StatusBadGateway,
			Response:       err.Error(),
			Posts:          []*pb.Post{},
			Categories:     []*pb.Category{},
			Successstories: []*pb.SuccesStory{},
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

//note:
//**add categories list
//**success stories
