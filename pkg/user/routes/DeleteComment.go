package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User report Post godoc
//
// @Summary  User can report a post
// @Description User can Report post
// @Security		api_key
// @Tags   User Post
// @Accept   json
// @Produce  json
// @Param   commentId query  string true "Comment ID"
// @Success  200   {object} pb.CommentPostResponse
// @Router   /user/post/comment/delete  [delete]
func DeleteComment(ctx *gin.Context, c pb.UserServiceClient) {
	commentIDstr := ctx.Query("commentId")
	commentID, err := strconv.Atoi(commentIDstr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := ctx.GetInt64("userId")
	log.Println("Sending Data: ", commentID, userId)
	res, err := c.DeleteComment(context.Background(), &pb.DeleteCommentRequest{
		Userid: int32(userId),
		Commentid: int32(commentID),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusCreated, &res)
}
