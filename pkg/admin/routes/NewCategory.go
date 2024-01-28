package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type NewCategoryBody struct {
	Category string `json:"category" validate:"required,min=2,max=20,alpha"`
}

// New Category godoc
//
//	@Summary		Create Category
//	@Description	Admin can create New Category
//	@Tags			Admin Categories
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			NewCategoryBody	body		NewCategoryBody	true	"Category Name"
//	@Success		200				{object}	pb.NewCategoryResponse
//	@Failure		400				{object}	pb.NewCategoryResponse
//	@Failure		502				{object}	pb.NewCategoryResponse
//	@Router			/admin/categories/new       [post]
func NewCategory(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating NewCategory...")

	NewCategoryBody := NewCategoryBody{}

	if err := ctx.BindJSON(&NewCategoryBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.NewCategoryResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
			Category:    nil,
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(NewCategoryBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.NewCategoryResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid data" + err.Error(),
			Category:    nil,
		})
		return
	}
	res, err := c.NewCategory(context.Background(), &pb.NewCategoryRequest{Category: NewCategoryBody.Category})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.NewCategoryResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
			Category:    nil,
		})
		return
	}
	log.Println("Recieved Data: ", res)

	ctx.JSON(http.StatusOK, &res)
}
