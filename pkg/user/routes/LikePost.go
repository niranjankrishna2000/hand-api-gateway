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

type LikePostBody struct {
	PostId int `json:"postId" validate:"required,min=1,max=999,number"`
}

// User Like Post godoc
//
//	@Summary		User can Like a post
//	@Description	User can Like post
//	@Tags			User Post
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			postid	query		string	true	"Post ID"
//	@Success		200		{object}	pb.LikePostResponse
//	@Failure		400		{object}	pb.LikePostResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.LikePostResponse
//	@Router			/user/post/like  [post]
func LikePost(ctx *gin.Context, c pb.UserServiceClient) {
	postID, err := strconv.Atoi(ctx.Query("postid"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.CommentPostResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Post:     nil,
		})
		return
	}
	likePostBody := LikePostBody{PostId: postID}
	userId := ctx.GetInt64("userId")
	log.Println("Sending Data: ", postID, userId)
	validator := validator.New()
	if err := validator.Struct(likePostBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.LikePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     nil,
		})
		return
	}
	res, err := c.LikePost(context.Background(), &pb.LikePostRequest{
		Userid: int32(userId),
		Postid: int32(postID),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.LikePostResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusCreated, &res)
}
