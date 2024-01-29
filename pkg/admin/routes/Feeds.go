package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type FeedsBody struct {
	Limit     int    `json:"limit" validate:"min=0,max=50,number"`
	Page      int    `json:"page" validate:"min=0,max=99,number"`
	Searchkey string `json:"searchkey" validate:"max=10,ascii"`
}

// Admin feeds godoc
//
//	@Summary		Feeds
//	@Description	Admin can see feeds
//	@Tags			Admin Feeds
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			FeedsBody	body		FeedsBody	true	"Page Details and Searchkey "
//	@Success		200			{object}	pb.FeedsResponse
//	@Failure		400			{object}	pb.FeedsResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.FeedsResponse
//	@Router			/admin/feeds  [get]
func Feeds(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating Feeds...")

	feedsBody := FeedsBody{}

	if err := ctx.BindJSON(&feedsBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.FeedsResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Posts:     nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(feedsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.FeedsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Posts:     nil,
		})
		return
	}
	res, err := c.Feeds(context.Background(), &pb.FeedsRequest{
		Page: int32(feedsBody.Page), 
		Limit: int32(feedsBody.Limit), 
		Searchkey: feedsBody.Searchkey,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.FeedsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Posts:     nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
