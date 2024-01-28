package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/admin/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type DeletePostBody struct {
	PostId int `json:"postId" validate:"required,min=1,max=999,number"`
}

// Delete post godoc
//
//	@Summary		Delete post
//	@Description	Admin can delete post here
//	@Tags			Admin Feeds
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			DeletePostBody	body		DeletePostBody	true	"Post ID "
//	@Success		200				{object}	pb.DeletePostResponse
//	@Failure		400				{object}	pb.DeletePostResponse
//	@Failure		502				{object}	pb.DeletePostResponse
//	@Router			/admin/post/delete  [delete]
func DeletePost(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating DeletePost...")

	deletePostBody := DeletePostBody{}

	if err := ctx.BindJSON(&deletePostBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.DeletePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(deletePostBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DeletePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid Post ID",
		})
		return
	}
	res, err := c.DeletePost(context.Background(), &pb.DeletePostRequest{Id: int32(deletePostBody.PostId)})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeletePostResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
		})
		return
	}
	log.Println("Recieved data : ", res)
	ctx.JSON(http.StatusOK, &res)
}
