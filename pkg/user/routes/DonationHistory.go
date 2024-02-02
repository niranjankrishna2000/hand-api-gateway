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

type PageBody struct {
	Limit     int    `json:"limit" validate:"min=0,max=50,number"`
	Page      int    `json:"page" validate:"min=0,max=99,number"`
	Searchkey string `json:"searchkey" validate:"max=10,ascii"`
}

// Donation History godoc
//
//	@Summary		Donation History
//	@Description	User can see Donation History
//	@Tags			User Donations
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			limit		query		int		false	"limit"
//	@Param			page		query		int		false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.DonationHistoryResponse
//	@Failure		400			{object}	pb.DonationHistoryResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.DonationHistoryResponse
//	@Router			/user/post/donate/history  [get]
func DonationHistory(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("starting User Donation history")
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

	userId := ctx.GetInt64("userId")
	log.Println("Collected data : ", page, limit, searchkey, userId)
	validator := validator.New()
	if err := validator.Struct(pageBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DonationHistoryResponse{
			Status:    http.StatusBadRequest,
			Response:  "Invalid data" + err.Error(),
			Donations: nil,
		})
		return
	}
	res, err := c.DonationHistory(context.Background(), &pb.DonationHistoryRequest{
		Page:      int32(page),
		Limit:     int32(limit),
		Searchkey: searchkey,
		Userid:    int32(userId),
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DonationHistoryResponse{
			Status:    http.StatusBadGateway,
			Response:  "Error in internal server",
			Donations: nil,
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)

}
