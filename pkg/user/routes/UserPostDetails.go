package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserPostDetailsBody struct {
	PostID int `json:"postId" validate:"min=1,max=999,number"`
}

// Post detail godoc
//
//	@Summary		Post detail
//	@Description	User can get post detail
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			id	query		string	true	"Post ID"
//	@Success		200	{object}	pb.UserPostDetailsResponse
//	@Failure		400	{object}	pb.UserPostDetailsResponse
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		502	{object}	pb.UserPostDetailsResponse
//	@Router			/user/post/details  [get]
func UserPostDetails(ctx *gin.Context, c pb.UserServiceClient) {
	postID, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pb.UserPostDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "Couldn't fetch data from user" + err.Error(),
			Post:     nil,
		})
		return
	}

	postId := UserPostDetailsBody{PostID: postID}
	validator := validator.New()
	if err := validator.Struct(postId); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.UserPostDetailsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     nil,
		})
		return
	}
	res, err := c.UserPostDetails(context.Background(), &pb.UserPostDetailsRequest{PostID: int32(postID)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.UserPostDetailsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
