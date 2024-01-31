package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ReportPostRequestBody struct {
	Text   string `json:"text" validator:"required,max=20,ascii"`
	PostId int    `json:"postId" validator:"required,min=1,max=999,number"`
}

// Report Post godoc
//
//	@Summary		Report post
//	@Description	User can Report post
//	@Tags			User Post
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			body	body		ReportPostRequestBody	true	"Report Post Data"
//	@Success		200		{object}	pb.ReportPostResponse
//	@Failure		400		{object}	pb.ReportPostResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.ReportPostResponse
//	@Router			/user/post/details/report  [post]
func ReportPost(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Reporting post started...")
	body := ReportPostRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println("Couldn't fetch data :", err)
		ctx.JSON(http.StatusBadRequest, pb.ReportPostResponse{
			Status:   http.StatusBadRequest,
			Response: "Couldn't fetch data",
			Post:     nil,
		})
		return
	}
	//validator

	userId := ctx.GetInt64("userId")
	log.Println("Collected Data: ", body, userId)
	validator := validator.New()
	if err := validator.Struct(body); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.ReportPostResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     nil,
		})
		return
	}
	res, err := c.ReportPost(context.Background(), &pb.ReportPostRequest{
		Text:   body.Text,
		Userid: int32(userId),
		Postid: int32(body.PostId),
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.ReportPostResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Recieved Data: ", res)

	ctx.JSON(http.StatusOK, &res)
}
