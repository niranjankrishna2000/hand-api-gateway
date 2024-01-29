package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CategoryPostsBody struct {
	Limit      int `json:"limit" validate:"min=1,max=99,number"`
	Page       int `json:"page" validate:"min=1,max=99,number"`
	CategoryID int `json:"categoryId" validate:"required,min=1,max=50,number"`
}

// Admin Category Posts godoc
//
//	@Summary		Admin can see Category posts
//	@Description	Admin can see Category posts
//	@Tags			Admin Categories
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		string	false	"limit"
//	@Param			page		query		string	false	"Page number"
//	@Param			categoryId	query		string	true	"Category Id"
//	@Success		200			{object}	pb.CategoryPostsResponse
//	@Failure		400			{object}	pb.CategoryPostsResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.CategoryPostsResponse
//	@Router			/admin/categories/categorylist/posts  [get]
func CategoryPosts(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating CategoryPosts...")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	categoryId, err := strconv.Atoi(ctx.Query("categoryId"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.CategoryPostsResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with category Id",
			Posts:     nil,
		})
		return
	}
	categoryPostsBody := CategoryPostsBody{Page: page,Limit: limit,CategoryID: categoryId}
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
		Page:       int32(categoryPostsBody.Page),
		Limit:      int32(categoryPostsBody.Limit),
		Categoryid: int32(categoryPostsBody.CategoryID),
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.CategoryPostsResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Category: nil,
			Posts:    nil,
		})
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
