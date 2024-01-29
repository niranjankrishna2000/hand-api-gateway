package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type EditPostRequestBody struct {
	PostId    int    `json:"postid" validate:"required,max=999,min=1,number"`
	Text      string `json:"text" validate:"required,max=50,ascii"` //test
	Place     string `json:"place" validate:"required,max=10,ascii"`
	Amount    int    `json:"amount" validate:"min=100,number"`
	AccountNo string `json:"accno" validate:"max=17,min=9,alphanum"`
	Address   string `json:"address" validate:"required,max=50,ascii"`
	Image     string `json:"image"`
	Date      string `json:"date" validate:"required"`
}

// Edit Post godoc
//
//	@Summary		Edit post
//	@Description	User can edit post
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			body	body		EditPostRequestBody	true	"Edit post Data"
//	@Success		200		{object}	pb.EditPostResponse
//	@Failure		400		{object}	pb.EditPostResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.EditPostResponse
//	@Router			/user/post/edit  [patch]
func EditPost(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Edit post started...")
	body := EditPostRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.EditPostResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			Post:     nil,
		})
		return
	}
	//validator
	validator := validator.New()
	if err := validator.Struct(body); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.EditPostResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:    nil,
		})
		return
	}
	res, err := c.EditPost(context.Background(), &pb.EditPostRequest{
		Text:    body.Text,
		Place:   body.Place,
		Amount:  int64(body.Amount),
		Image:   body.Image,
		Date:    body.Date,
		Postid:  int32(body.PostId),
		Accno:   body.AccountNo,
		Address: body.Address,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.EditPostResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
