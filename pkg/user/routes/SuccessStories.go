package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type SuccesStoryBody struct {
	Title string `json:"title" validate:"required,max=15,ascii"`
	Text  string `json:"text" validate:"required,max=50,ascii"`
	Place string `json:"place" validate:"required,max=10,ascii"`
	Image string `json:"image"`
}
type EditSuccesStoryBody struct {
	Title string `json:"title,omitempty" validate:"max=15,ascii"`
	Text  string `json:"text,omitempty" validate:"max=50,ascii"`
	Place string `json:"place,omitempty" validate:"max=10,ascii"`
	Image string `json:"image,omitempty" `
}

// Get Success Story godoc
//
//	@Summary		Success Story
//	@Description	User can see Success Stories
//	@Tags			User Success Story
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			limit		query		int		false	"limit"
//	@Param			page		query		int		false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.GetSuccessStoryResponse
//	@Failure		400			{object}	pb.GetSuccessStoryResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.GetSuccessStoryResponse
//	@Router			/user/success-stories  [get]
func GetSuccessStory(ctx *gin.Context, c pb.UserServiceClient) {
	log.Println("starting GetSuccessStory")
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	searchkey := ctx.Query("searchkey")
	pageBody := PageBody{Page: page, Limit: limit, Searchkey: searchkey}

	log.Println("Collected data : ", page, limit, searchkey)
	//note: ** need userid?
	validator := validator.New()
	if err := validator.Struct(pageBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.GetSuccessStoryResponse{
			Status:         http.StatusBadRequest,
			Response:       "Invalid data :" + err.Error(),
			SuccessStories: []*pb.SuccesStory{},
		})
		return
	}
	userid:=ctx.GetInt64("userid")
	res, err := c.GetSuccessStory(context.Background(), &pb.GetSuccessStoryRequest{Page: int32(page), Limit: int32(limit), Searchkey: searchkey, Userid: int32(userid)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.GetSuccessStoryResponse{
			Status:         http.StatusBadGateway,
			Response:       err.Error(),
			SuccessStories: []*pb.SuccesStory{},
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

// Add Success Story godoc
//
//	@Summary		Create Success Story
//	@Description	User can create Success Stories
//	@Tags			User Success Story
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			succesStoryBody	body		SuccesStoryBody	true	"Success Story Data"
//	@Success		200				{object}	pb.AddSuccessStoryResponse
//	@Failure		400				{object}	pb.AddSuccessStoryResponse
//	@Failure		403				{string}	string	"You have not logged in"
//	@Failure		502				{object}	pb.AddSuccessStoryResponse
//	@Router			/user/success-stories  [post]
func AddSuccessStory(ctx *gin.Context, c pb.UserServiceClient) {
	succesStoryBody := SuccesStoryBody{}
	if err := ctx.BindJSON(&succesStoryBody); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.AddSuccessStoryResponse{
			Status:       http.StatusBadGateway,
			Response:     "Couldnt fetch data from client",
			SuccessStory: &pb.SuccesStory{},
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(succesStoryBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.AddSuccessStoryResponse{
			Status:       http.StatusBadRequest,
			Response:     "Invalid data" + err.Error(),
			SuccessStory: &pb.SuccesStory{},
		})
		return
	}
	userid:=ctx.GetInt64("userid")
	res, err := c.AddSuccessStory(context.Background(), &pb.AddSuccessStoryRequest{
		UserId: int32(userid),
		Title:  succesStoryBody.Title,
		Text:   succesStoryBody.Text,
		Image:  succesStoryBody.Image,
		Place:  succesStoryBody.Place,
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.AddSuccessStoryResponse{
			Status:       http.StatusBadGateway,
			Response:     "Error in internal server:" + err.Error(),
			SuccessStory: &pb.SuccesStory{},
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)
}

// Edit Success Story godoc
//
//	@Summary		Edit Success Story
//	@Description	User can edit Success Stories
//	@Tags			User Success Story
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			postId			query		int					true	"Post ID"
//	@Param			successStory	body		EditSuccesStoryBody	true	"Story Data"
//	@Success		200				{object}	pb.EditSuccessStoryResponse
//	@Failure		400				{object}	pb.EditSuccessStoryResponse
//	@Failure		403				{string}	string	"You have not logged in"
//	@Failure		502				{object}	pb.EditSuccessStoryResponse
//	@Router			/user/success-stories  [patch]
func EditSuccessStory(ctx *gin.Context, c pb.UserServiceClient) {
	successStory := EditSuccesStoryBody{}
	if err := ctx.BindJSON(&successStory); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.EditSuccessStoryResponse{
			Status:       http.StatusBadGateway,
			Response:     "Couldnt fetch data from client",
			SuccessStory: &pb.SuccesStory{},
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(successStory); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.EditSuccessStoryResponse{
			Status:       http.StatusBadRequest,
			Response:     "Invalid data",
			SuccessStory: &pb.SuccesStory{},
		})
		return
	}
	userid:=ctx.GetInt64("userid")
	res, err := c.EditSuccessStory(context.Background(), &pb.EditSuccessStoryRequest{
		UserId: int32(userid),
		Title:  successStory.Title,
		Text:   successStory.Text,
		Image:  successStory.Image,
		Place:  successStory.Place,
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.EditSuccessStoryResponse{
			Status:       http.StatusBadGateway,
			Response:     "Error in internal server:" + err.Error(),
			SuccessStory: &pb.SuccesStory{},
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)
}

// Delete Success Story godoc
//
//	@Summary		Success Story
//	@Description	User can Delete Success Stories
//	@Tags			User Success Story
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			storyId	query		int	true	"Success Story ID"
//	@Success		200		{object}	pb.DeleteSuccessStoryResponse
//	@Failure		400		{object}	pb.DeleteSuccessStoryResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.DeleteSuccessStoryResponse
//	@Router			/user/success-stories  [delete]
func DeleteSuccessStory(ctx *gin.Context, c pb.UserServiceClient) {
	storyId, err := strconv.Atoi(ctx.Query("storyId"))
	if err != nil {
		log.Println("Error with storyId :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeleteSuccessStoryResponse{
			Status:         http.StatusBadRequest,
			Response:       "Invalid Id",
			SuccessStories: []*pb.SuccesStory{},
		})
		return
	}
	storyIdBody := PostIdBody{PostId: storyId}
	validator := validator.New()
	if err := validator.Struct(storyIdBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteSuccessStoryResponse{
			Status:         http.StatusBadRequest,
			Response:       "Invalid id",
			SuccessStories: []*pb.SuccesStory{},
		})
		return
	}
	userid:=ctx.GetInt64("userid")
	res, err := c.DeleteSuccessStory(context.Background(), &pb.DeleteSuccessStoryRequest{
		Userid:  int32(userid),
		Storyid: int32(storyIdBody.PostId),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeleteSuccessStoryResponse{
			Status:         http.StatusBadGateway,
			Response:       "Error in internal server:" + err.Error(),
			SuccessStories: []*pb.SuccesStory{},
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)
}
