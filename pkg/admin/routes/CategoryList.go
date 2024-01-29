package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CategoryListBody struct {
	Limit     int    `json:"limit" validate:"min=1,max=99,number"`
	Page      int    `json:"page" validate:"min=1,max=99,number"`
	Searchkey string `json:"searchkey"`
}

// Admin Category List godoc
//
//	@Summary		Categories
//	@Description	Admin can see Categories
//	@Tags			Admin Categories
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			CategoryListBody	body		CategoryListBody	true	"Page Details and Searchkey "
//	@Success		200					{object}	pb.CategoryListResponse
//	@Failure		400					{object}	pb.CategoryListResponse
//	@Failure		403					{string}	string	"You have not logged in"
//	@Failure		502					{object}	pb.CategoryListResponse
//	@Router			/admin/categories/categorylist  [get]
func CategoryList(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	categoryListBody := CategoryListBody{}

	if err := ctx.BindJSON(&categoryListBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.CategoryListResponse{
			Status:     http.StatusBadRequest,
			Response:   "Error with request",
			Categories: nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(categoryListBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.CategoryListResponse{
			Status:     http.StatusBadRequest,
			Response:   "Invalid data" + err.Error(),
			Categories: nil,
		})
		return
	}
	res, err := c.CategoryList(context.Background(), &pb.CategoryListRequest{
		Page:      int32(categoryListBody.Page),
		Limit:     int32(categoryListBody.Limit),
		Searchkey: categoryListBody.Searchkey})

		if err != nil {
			log.Println("Error with internal server :", err)
			ctx.JSON(http.StatusBadGateway, pb.CategoryListResponse{
				Status:   http.StatusBadGateway,
				Response: "Error in internal server",
				Categories:     nil,
			})
			return
		}
		log.Println("Recieved data : ", res)
	
	ctx.JSON(http.StatusOK, &res)
}
