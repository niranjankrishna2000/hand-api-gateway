package routes

import (
	"context"
	"log"
	"net/http"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
)

type EditPostRequestBody struct {
	PostId    int    `json:"postid"`
	Text      string `json:"text"`
	Place     string `json:"place"`
	Amount    int    `json:"amount"`
	AccountNo string `json:"accno"`
	Address   string `json:"address"`
	Image     string `json:"image"`
	Date      string `json:"date"`
}

// User Edit Post godoc
//
//	@Summary		User can edit post
//	@Description	User can edit post
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			body	body		EditPostRequestBody	true	"Edit post Data"
//	@Success		200		{object}	pb.EditPostResponse
//	@Router			/user/post/edit  [patch]
func EditPost(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Edit post started...")
	body := EditPostRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	res, err := c.EditPost(context.Background(), &pb.EditPostRequest{
		Text:    body.Text,
		Place:   body.Place,
		Amount:  int64(body.Amount),
		Image:   body.Image,
		Date:    body.Date,
		Postid:  int32(body.PostId),
		Accno:   body.AccountNo,
		Address: body.Address,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, &gin.H{
			"Status": http.StatusBadRequest,
			"Response":"Failed to update post"+err.Error(),
		})
		//ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
