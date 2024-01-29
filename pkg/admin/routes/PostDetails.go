package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type PostDetailsBody struct {
	PostId int `json:"postId" validate:"required,min=1,max=999,number"`
}

// Post detail godoc
//
//	@Summary		Post detail
//	@Description	Admin can get post details
//	@Tags			Admin Feeds
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			postid	query		string	true	"Post ID "
//	@Success		200		{object}	pb.PostDetailsResponse
//	@Failure		400		{object}	pb.PostDetailsResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.PostDetailsResponse
//	@Router			/admin/post/details  [get]
func PostDetails(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating PostDetails...")
	postId, err := strconv.Atoi(ctx.Query("postid"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.PostDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with post Id",
			Post:     nil,
		})
		return
	}
	postDetailsBody := PostDetailsBody{PostId: postId}

	if err := ctx.BindJSON(&postDetailsBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.PostDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Post:     nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(postDetailsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.PostDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid Post ID",
			Post:     nil,
		})
		return
	}
	res, err := c.PostDetails(context.Background(), &pb.PostDetailsRequest{PostID: int32(postDetailsBody.PostId)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.PostDetailsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Data recieved: ", res)

	ctx.JSON(http.StatusOK, &res)
}
