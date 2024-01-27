package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportCommentRequestBody struct {
	Text      string `json:"text"`
	CommentId int    `json:"commentId"`
}

// User report Comment godoc
//
// @Summary  User can report a Comment
// @Description User can Report Comment
// @Tags   User Post
// @Accept   json
// @Produce  json
// @Param   reportCommentBody body  ReportCommentRequestBody true "Report Comment Data"
// @Success  200   {object} pb.ReportCommentResponse
// @Router   /user/post/comment/report  [post]
func ReportComment(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Reporting comment started...")
	reportCommentBody := ReportCommentRequestBody{}

	if err := ctx.BindJSON(&reportCommentBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := ctx.GetInt64("userId")
	log.Println("Collected Data: ",reportCommentBody,userId)
	res, err := c.ReportComment(context.Background(), &pb.ReportCommentRequest{
		Text:   reportCommentBody.Text,
		Userid: int32(userId),
		Commentid: int32(reportCommentBody.CommentId),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved Data: ",res)
	ctx.JSON(http.StatusCreated, &res)
}
