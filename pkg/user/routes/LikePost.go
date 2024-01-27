package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User Like Post godoc
//
// @Summary  User can Like a post
// @Description User can Like post
// @Tags   User Post
// @Accept   json
// @Produce  json
// @Param   postid query  string true "PostID"
// @Success  200   {object} pb.LikePostResponse
// @Router   /user/post/like  [post]
func LikePost(ctx *gin.Context, c pb.UserServiceClient) {
	postIDstr := ctx.Query("postid")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := ctx.GetInt64("userId")
	log.Println("Sending Data: ", postID, userId)
	res, err := c.LikePost(context.Background(), &pb.LikePostRequest{
		Userid: int32(userId),
		Postid: int32(postID),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusCreated, &res)
}
