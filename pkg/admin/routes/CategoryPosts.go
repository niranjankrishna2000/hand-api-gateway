package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CategoryPostsBody struct {
	Limit      int `json:"limit" validate:"max=99,number"`
	Page       int `json:"page" validate:"max=99,number"`
	CategoryID int `json:"categoryId" validate:"required,max=50,number"`
}

// Admin Category Posts godoc
//
//	@Summary		Admin can see Category posts
//	@Description	Admin can see Category posts
//	@Tags			Admin Categories
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			CategoryPostsBody	body		CategoryPostsBody	true	"Page Details and Category ID "
//	@Success		200					{object}	pb.CategoryPostsResponse
//	@Failure		400					{object}	pb.CategoryPostsResponse
//	@Failure		502					{object}	pb.CategoryPostsResponse
//	@Router			/admin/categories/categorylist/posts  [get]
func CategoryPosts(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	categoryPostsBody := CategoryPostsBody{}

	if err := ctx.BindJSON(&categoryPostsBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.CategoryPostsResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Category: nil,
			Posts:    nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(categoryPostsBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.CategoryPostsResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Category: nil,
			Posts:    nil,
		})
		return
	}
	res, err := c.CategoryPosts(context.Background(), &pb.CategoryPostsRequest{
		Page:  int32(categoryPostsBody.Page),
		Limit: int32(categoryPostsBody.Limit),
		Categoryid: int32(categoryPostsBody.CategoryID),
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.CategoryPostsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Category:     nil,
			Posts: nil,
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
