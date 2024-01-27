package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentPostRequestBody struct {
	Text   string `json:"text"`
	PostId int    `json:"postid"`
}

// User Comment Post godoc
//
// @Summary  User can Comment new post
// @Description User can Comment new post
// @Tags   User Post
// @Accept   json
// @Produce  json
// @Param   commentBody body  CommentPostRequestBody true "Comment post Data"
// @Success  200   {object} pb.CommentPostResponse
// @Router   /user/post/comment  [post]
func CommentPost(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Comment post intiated...")
	commentBody := CommentPostRequestBody{}

	if err := ctx.BindJSON(&commentBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := ctx.GetInt64("userId")
	log.Println("Data collected :", commentBody, userId)
	res, err := c.CommentPost(context.Background(), &pb.CommentPostRequest{Postid: int32(commentBody.PostId),
		Userid:  int32(userId),
		Comment: commentBody.Text,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Data recieved :", res)

	ctx.JSON(http.StatusCreated, &res)

}
