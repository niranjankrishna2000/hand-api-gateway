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

type UpdateIdBody struct {
	UpdateId int `json:"updateId" validate:"required,min=1,max=999,number"`
}
type UpdateBody struct {
	Id    int    `json:"Id" validate:"required,min=1,max=999,number"`
	Title string `json:"title" validate:"required,min=1,max=20,ascii"`
	Text  string `json:"text" validate:"required,min=1,max=50,ascii"`
}
type EditUpdateBody struct {
	Id    int    `json:"Id" validate:"required,min=1,max=999,number"`
	Title string `json:"title,omitempty" validate:"min=1,max=20,ascii"`
	Text  string `json:"text,omitempty" validate:"min=1,max=50,ascii"`//test
}

// Updates godoc
//
//	@Summary		Post Updates
//	@Description	User can see updates about a campaign
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			postId	query		string	true	"Post ID"
//	@Success		200		{object}	pb.GetUpdatesResponse
//	@Failure		400		{object}	pb.GetUpdatesResponse
//	@Failure		403		{string}	string	"You have not logged in"
//	@Failure		502		{object}	pb.GetUpdatesResponse
//	@Router			/user/post/updates  [get]
func GetUpdates(ctx *gin.Context, c pb.UserServiceClient) {
	postId, err := strconv.Atoi(ctx.Query("postId"))
	if err != nil {
		log.Println("Error with postId :", err)
		ctx.JSON(http.StatusBadRequest, pb.GetUpdatesResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid post Id",
			Updates:  []*pb.Update{},
		})
		return
	}
	postIdBody := PostIdBody{PostId: postId}
	validator := validator.New()
	if err := validator.Struct(postIdBody); err != nil {
		ctx.JSON(http.StatusBadRequest, pb.GetUpdatesResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid post id",
			Updates:  []*pb.Update{},
		})
		return
	}
	res, err := c.GetUpdates(context.Background(), &pb.GetUpdatesRequest{
		Postid: int32(postId),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.GetUpdatesResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server:" + err.Error(),
			Updates:  []*pb.Update{},
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)

}

// Add Update godoc
//
//	@Summary		Add Update
//	@Description	User can add an updates about a campaign
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			updateBody	body		UpdateBody	true	"Update Data"
//	@Success		200			{object}	pb.AddUpdatesResponse
//	@Failure		400			{object}	pb.AddUpdatesResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.AddUpdatesResponse
//	@Router			/user/post/updates  [post]
func AddUpdate(ctx *gin.Context, c pb.UserServiceClient) {
	updateBody := UpdateBody{}
	if err := ctx.BindJSON(&updateBody); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.AddUpdatesResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			Updates:     []*pb.Update{},
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(updateBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.AddUpdatesResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data"+err.Error(),
			Updates:  []*pb.Update{},
		})
		return
	}
	res, err := c.AddUpdates(context.Background(), &pb.AddUpdatesRequest{
		Userid: int32(ctx.GetInt64("userId")),
		Postid: int32(updateBody.Id),
		Title: updateBody.Title,
		Text: updateBody.Text,
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.AddUpdatesResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server:" + err.Error(),
			Updates:  []*pb.Update{},
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)

}
// Edit Update godoc
//
//	@Summary		Edit Update
//	@Description	User can edit an updates about a campaign
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			editUpdateBody	body		EditUpdateBody	true	"Update Data"
//	@Success		200				{object}	pb.EditUpdatesResponse
//	@Failure		400				{object}	pb.EditUpdatesResponse
//	@Failure		403				{string}	string	"You have not logged in"
//	@Failure		502				{object}	pb.EditUpdatesResponse
//	@Router			/user/post/updates  [patch]
func EditUpdate(ctx *gin.Context, c pb.UserServiceClient) {
	editUpdateBody := EditUpdateBody{}
	if err := ctx.BindJSON(&editUpdateBody); err != nil {
		ctx.JSON(http.StatusBadGateway, pb.EditUpdatesResponse{
			Status:   http.StatusBadGateway,
			Response: "Couldnt fetch data from client",
			Updates:     []*pb.Update{},
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(editUpdateBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.EditUpdatesResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid id",
			Updates:  []*pb.Update{},
		})
		return
	}
	res, err := c.EditUpdates(context.Background(), &pb.EditUpdatesRequest{
		Userid:   int32(ctx.GetInt64("userId")),
		Updateid: int32(editUpdateBody.Id),
		Text: editUpdateBody.Text,
		Title: editUpdateBody.Title,
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.EditUpdatesResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server:" + err.Error(),
			Updates:  []*pb.Update{},
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)

}

// Delete Update godoc
//
//	@Summary		Delete Update
//	@Description	User can delete an updates about a campaign
//	@Tags			User Post
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Param			updateId	query		string	true	"Update ID"
//	@Success		200			{object}	pb.DeleteUpdatesResponse
//	@Failure		400			{object}	pb.DeleteUpdatesResponse
//	@Failure		403			{string}	string	"You have not logged in"
//	@Failure		502			{object}	pb.DeleteUpdatesResponse
//	@Router			/user/post/updates  [delete]
func DeleteUpdate(ctx *gin.Context, c pb.UserServiceClient) {
	updateId, err := strconv.Atoi(ctx.Query("updateId"))
	if err != nil {
		log.Println("Error with postId :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeleteUpdatesResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid Id",
			Updates:  []*pb.Update{},
		})
		return
	}
	updateIdBody := UpdateIdBody{UpdateId: updateId}
	validator := validator.New()
	if err := validator.Struct(updateIdBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteUpdatesResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid id",
			Updates:  []*pb.Update{},
		})
		return
	}
	res, err := c.DeleteUpdates(context.Background(), &pb.DeleteUpdatesRequest{
		Userid:   int32(ctx.GetInt64("userId")),
		Updateid: int32(updateId),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeleteUpdatesResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server:" + err.Error(),
			Updates:  []*pb.Update{},
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)

}
