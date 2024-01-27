package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

type CreatePostRequestBody struct {
	Text      string `json:"text"`
	Place     string `json:"place"`
	Amount    int    `json:"amount"`
	AccountNo string `json:"accno"`
	Address   string `json:"address"`
	Image     string `json:"image"`
	Date      string `json:"date"`
}

// User create Post godoc
//
//	@Summary		User can create new post
//	@Description	User can create new post
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreatePostRequestBody	true	"Create post Data"
//	@Success		200		{object}	pb.CreatePostResponse
//	@Router			/user/post/new  [post]
func CreatePost(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Create post started...")
	body := CreatePostRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if body.Amount < 0 {
		//ctx.AbortWithError(http.StatusBadRequest, errors.New("add a postitive amount"))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status":   http.StatusBadRequest,
			"Response": "add a positive number"})
		return
	}
	userId := ctx.GetInt64("userId")
	log.Println("User ID:", userId)
	res, err := c.CreatePost(context.Background(), &pb.CreatePostRequest{
		Text:    body.Text,
		Place:   body.Place,
		Amount:  int64(body.Amount),
		Image:   body.Image,
		Date:    body.Date,
		Userid:  int32(userId),
		Accno:   body.AccountNo,
		Address: body.Address,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
