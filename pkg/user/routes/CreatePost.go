package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreatePostRequestBody struct {
	Text      string `json:"text" validate:"required,max=50,ascii"` //test
	Place     string `json:"place" validate:"required,max=10,ascii"`
	Amount    int    `json:"amount" validate:"min=100,number"`
	AccountNo string `json:"accno" validate:"max=17,min=9,alphanum"`
	Address   string `json:"address" validate:"required,max=50,ascii"`
	Image     string `json:"image"`
	Date      string `json:"date" validate:"required"`
}

// create Post godoc
//
//	@Summary		Create Post
//	@Description	User can create new post
//	@Tags			User Post
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreatePostRequestBody	true	"Create post Data"
//	@Success		200		{object}	pb.CreatePostResponse
//	@Failure		400		{object}	pb.CreatePostResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.CreatePostResponse
//	@Router			/user/post/new  [post]
func CreatePost(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Create post started...")
	body := CreatePostRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.CreatePostResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			Post:     nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(body); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.CreatePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:    nil,
		})
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
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.CreatePostResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
