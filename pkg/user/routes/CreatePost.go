package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"hand/pkg/user/pb"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreatePostRequestBody struct {
	Text       string `json:"text" validate:"required,max=50,ascii"`
	Title      string `json:"title" validate:"required,max=20,ascii"`
	Place      string `json:"place" validate:"required,max=10,ascii"`
	Amount     int    `json:"amount" validate:"min=100,number"`
	AccountNo  string `json:"accno" validate:"max=17,min=9,alphanum"`
	Address    string `json:"address" validate:"required,max=50,ascii"`
	Image      string `json:"image"`
	Date       string `json:"date" validate:"required"`
	CategoryId int    `json:"categoryId" validate:"required,min=1,max=10,number"`
	TaxBenefit bool   `json:"taxbenefit" validate:"required,boolean"`
}

// create Post godoc
//
//	@Summary		Create Post
//	@Description	User can create new post
//	@Tags			User Post
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreatePostRequestBody	true	"Create post Data"
//	@Success		200		{object}	pb.CreatePostResponse
//	@Failure		400		{object}	pb.CreatePostResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.CreatePostResponse
//	@Router			/user/post/new  [post]
func CreatePost(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("Create post started...")
	body := CreatePostRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.CreatePostResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			Post:     &pb.Post{},
		})
		return
	}
	_, err := time.Parse("2006-01-02 15:04:05", body.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pb.CreatePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Error Parsing time string",
			Post:     &pb.Post{},
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(body); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.CreatePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     &pb.Post{},
		})
		return
	}
	userId := ctx.GetInt64("userId")
	log.Println("User ID:", userId)
	res, err := c.CreatePost(context.Background(), &pb.CreatePostRequest{
		Text:       body.Text,
		Place:      body.Place,
		Amount:     int64(body.Amount),
		Image:      body.Image,
		Date:       body.Date,
		Userid:     int32(userId),
		Accno:      body.AccountNo,
		Address:    body.Address,
		Categoryid: int32(body.CategoryId),
		Taxbenefit: body.TaxBenefit,
	})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.CreatePostResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server : " + err.Error(),
			Post:     nil,
		})
		return
	}

	ctx.JSON(int(res.Status), &res)
}

// create Post godoc
//
//	@Summary		Create Post
//	@Description	Choose A category
//	@Tags			User Post
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	pb.GetCreatePostRequest
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		502	{object}	pb.GetCreatePostRequest
//	@Router			/user/post/new  [get]
func GetCreatePost(ctx *gin.Context, c pb.UserServiceClient) {
	res, err := c.GetCreatePost(context.Background(), &pb.GetCreatePostRequest{})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.GetCreatePostResponse{
			Status:     http.StatusBadGateway,
			Response:   "Error in internal server: " + err.Error(),
			Categories: []*pb.Category{},
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)
}

// Expire Post godoc
//
//	@Summary		Expire Post
//	@Description	Expire an active post
//	@Tags			User Post
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			postid	query		string	true	"Post Id"
//	@Success		200	{object}	pb.ExpirePostRequest
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		502	{object}	pb.ExpirePostRequest
//	@Router			/user/post/expire  [patch]
func ExpirePost(ctx *gin.Context, c pb.UserServiceClient) {
	postID, err := strconv.Atoi(ctx.Query("postId"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.ExpirePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Post:     nil,
		})
		return
	}
	postIdBody := PostIdBody{PostId: postID}
	validator := validator.New()
	if err := validator.Struct(postIdBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.ExpirePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Post:     nil,
		})
		return
	}
	res, err := c.ExpirePost(context.Background(), &pb.ExpirePostRequest{
		Userid: int32(ctx.GetInt64("userId")),
		Postid: int32(postID),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.ExpirePostResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Post:     nil,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusCreated, &res)
}

// Delete Post godoc
//
//	@Summary		Delete Post
//	@Description	Delete An expired post
//	@Tags			User Post
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			postid	query		string	true	"Post Id"
//	@Success		200	{object}	pb.DeletePostRequest
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		502	{object}	pb.DeletePostRequest
//	@Router			/user/post/delete  [delete]
func DeletePost(ctx *gin.Context, c pb.UserServiceClient) {
	postID, err := strconv.Atoi(ctx.Query("postId"))
	if err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.DeletePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
		})
		return
	}
	postIdBody := PostIdBody{PostId: postID}
	validator := validator.New()
	if err := validator.Struct(postIdBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DeletePostResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
		})
		return
	}
	res, err := c.DeletePost(context.Background(), &pb.DeletePostRequest{
		Userid: int32(ctx.GetInt64("userId")),
		Postid: int32(postID),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeletePostResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusCreated, &res)
}
