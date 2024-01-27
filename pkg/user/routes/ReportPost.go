package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

type ReportPostRequestBody struct {
	Text      string `json:"text"`
	PostId	  int `json:"postId"`
}

// User report Post godoc
//
// @Summary  User can report a post
// @Description User can Report post
// @Tags   User Post
// @Accept   json
// @Produce  json
// @Param   body body  ReportPostRequestBody true "Report Post Data"
// @Success  200   {object} pb.ReportPostResponse
// @Router   /user/post/details/report  [post]
func ReportPost(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Reporting post started...")
	body := ReportPostRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := ctx.GetInt64("userId")
	log.Println("Collected Data: ",body,userId)
	res, err := c.ReportPost(context.Background(), &pb.ReportPostRequest{
		Text:    body.Text,
		Userid:  int32(userId),
		Postid: int32(body.PostId),
		
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved Data: ",res)

	ctx.JSON(http.StatusCreated, &res)
}
