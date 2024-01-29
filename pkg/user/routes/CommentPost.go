package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CommentPostRequestBody struct {
	Text   string `json:"text" validate:"required,alpha,ascii"`//test
	PostId int    `json:"postid" validate:"required,min=1,max=999,number"`
}

// Comment Post godoc
//
//	@Summary		Comment
//	@Description	User can Comment new post
//	@Security		api_key
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Param			commentBody	body		CommentPostRequestBody	true	"Comment post Data"
//	@Success		200			{object}	pb.CommentPostResponse
//	@Failure		400			{object}	pb.CommentPostResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.CommentPostResponse
//	@Router			/user/post/comment  [post]
func CommentPost(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Comment post intiated...")
	commentBody := CommentPostRequestBody{}

	if err := ctx.BindJSON(&commentBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.CommentPostResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Post:     nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(commentBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.CommentPostResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     nil,
		})
		return
	}
	userId := ctx.GetInt64("userId")
	log.Println("Data collected :", commentBody, userId)
	res, err := c.CommentPost(context.Background(), &pb.CommentPostRequest{Postid: int32(commentBody.PostId),
		Userid:  int32(userId),
		Comment: commentBody.Text,
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.CommentPostResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Data recieved :", res)

	ctx.JSON(http.StatusCreated, &res)

}

