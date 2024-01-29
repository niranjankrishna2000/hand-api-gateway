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

type DeleteCommentBody struct {
	CommentId int `json:"commentId" validate:"required,min=1,max=99,number"`
}

// Delete Post godoc
//
//	@Summary		Delete a comment
//	@Description	User can Delete post
//	@Security		api_key
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Param			commentId	query		string	true	"Comment ID"
//	@Success		200			{object}	pb.DeleteCommentResponse
//	@Failure		400			{object}	pb.DeleteCommentResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.DeleteCommentResponse
//	@Router			/user/post/comment/delete  [delete]
func DeleteComment(ctx *gin.Context, c pb.UserServiceClient) {
	commentID, err := strconv.Atoi(ctx.Query("commentId"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteCommentResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Post:     nil,
		})
		return
	}
	commentIdBody := DeleteCommentBody{CommentId: commentID}
	userId := ctx.GetInt64("userId")
	log.Println("Sending Data: ", commentID, userId)
	validator := validator.New()
	if err := validator.Struct(commentIdBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteCommentResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     nil,
		})
		return
	}
	res, err := c.DeleteComment(context.Background(), &pb.DeleteCommentRequest{
		Userid:    int32(userId),
		Commentid: int32(commentID),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeleteCommentResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusCreated, &res)
}
