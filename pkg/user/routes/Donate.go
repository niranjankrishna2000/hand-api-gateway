package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

type DonateRequestBody struct {
	PostID int `json:"postID"`
	Amount int `json:"amount"`
}

// User Donate godoc
//
//	@Summary		User can donate for campaign
//	@Description	User can donate for campaigns
//	@Tags			User Donations
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			body	body		DonateRequestBody	true	"Donate Data"
//	@Success		200		{object}	pb.DonateResponse
//	@Router			/user/post/donate/  [post]
func Donate(ctx *gin.Context, c pb.UserServiceClient) {
	body := DonateRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println(body)

	id := ctx.GetInt64("userId")
	log.Println("User ID:", id)

	res, err := c.Donate(context.Background(), &pb.DonateRequest{Postid: int32(body.PostID), Amount: int32(body.Amount), Userid: int32(id)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}



//note: Add option to give name so that the user can have anonymity 