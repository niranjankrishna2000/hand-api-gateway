package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

// User get Post detail godoc
//
//	@Summary		User can get post detail
//	@Description	User can get post detail
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			id	query		string	true	"post id Data"
//	@Success		200	{object}	pb.UserPostDetailsResponse
//	@Router			/user/post/details  [get]
func UserPostDetails(ctx *gin.Context, c pb.UserServiceClient) {
	postIDstr := ctx.Query("id")
	postID, err := strconv.Atoi(postIDstr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.UserPostDetails(context.Background(), &pb.UserPostDetailsRequest{PostID: int32(postID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, errors.New("couldn't get post details:"+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
