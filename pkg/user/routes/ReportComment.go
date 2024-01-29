package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ReportCommentRequestBody struct {
	Text      string `json:"text" validate:"required,max=20,ascii"`
	CommentId int    `json:"commentId" validate:"required,min=1,max=999,number"`
}

// Report Comment godoc
//
//	@Summary		Report Comment
//	@Description	User can Report Comment
//	@Tags			User Post
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			reportCommentBody	body		ReportCommentRequestBody	true	"Report Comment Data"
//	@Success		200					{object}	pb.ReportCommentResponse
//	@Failure		400					{object}	pb.ReportCommentResponse
//	@Failure		403					{string}	string	"You have not logged in"
//	@Failure		502					{object}	pb.ReportCommentResponse
//	@Router			/user/post/comment/report  [post]
func ReportComment(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Reporting comment started...")
	reportCommentBody := ReportCommentRequestBody{}

	if err := ctx.BindJSON(&reportCommentBody); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.ReportCommentResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			Post:     nil,
		})
	}
	//validator
	userId := ctx.GetInt64("userId")
	log.Println("Collected Data: ", reportCommentBody, userId)
	validator := validator.New()
	if err := validator.Struct(reportCommentBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.ReportCommentResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:    nil,
		})
		return
	}
	res, err := c.ReportComment(context.Background(), &pb.ReportCommentRequest{
		Text:      reportCommentBody.Text,
		Userid:    int32(userId),
		Commentid: int32(reportCommentBody.CommentId),
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.ReportCommentResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusCreated, &res)
}
