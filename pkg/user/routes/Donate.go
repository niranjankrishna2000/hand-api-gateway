package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type DonateRequestBody struct {
	PostID int `json:"postID" validate:"required,number,min=1,max=999"`
	Amount int `json:"amount" validate:"required,number,min=100"`
}

// Donate godoc
//
//	@Summary		Donate
//	@Description	User can donate for campaigns
//	@Tags			User Donations
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			body	body		DonateRequestBody	true	"Donation Data"
//	@Success		200		{object}	pb.DonateResponse
//	@Failure		400		{object}	pb.DonateResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.DonateResponse
//	@Router			/user/post/donate/  [post]
func Donate(ctx *gin.Context, c pb.UserServiceClient) {
	body := DonateRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.DonateResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			Post:     nil,
		})
		return
	}
	log.Println(body)

	id := ctx.GetInt64("userId")
	log.Println("User ID:", id)
	validator := validator.New()
	if err := validator.Struct(body); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DonateResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     nil,
		})
		return
	}
	res, err := c.Donate(context.Background(), &pb.DonateRequest{Postid: int32(body.PostID), Amount: int32(body.Amount), Userid: int32(id)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DonateResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

//note: Add option to give name so that the user can have anonymity
