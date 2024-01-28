package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type DeleteCategoryBody struct {
	CategoryId int `json:"categoryId" validate:"required,min=1,max=99,number"`
}

// Delete Category godoc
//
//	@Summary		Delete Category
//	@Description	Admin can Delete Category
//	@Tags			Admin Categories
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			DeleteCategoryBody	body		DeleteCategoryBody	true	"Category ID "
//	@Success		200					{object}	pb.DeleteCategoryResponse
//	@Failure		400					{object}	pb.DeleteCategoryResponse
//	@Failure		502					{object}	pb.DeleteCategoryResponse
//	@Router			/admin/categories/delete  [delete]
func DeleteCategory(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating DeleteCategory...")

	deleteCategoryBody := DeleteCategoryBody{}

	if err := ctx.BindJSON(&deleteCategoryBody); err != nil {
		log.Println("Error while fetching data :", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteCategoryResponse{
			Status:   http.StatusBadRequest,
			Response: "Error with request",
		})
		return
	}
	validator := validator.New()
	if err := validator.Struct(deleteCategoryBody); err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusBadRequest, pb.DeleteCategoryResponse{
			Status:   http.StatusBadRequest,
			Response: "Invalid Category ID",
		})
		return
	}
	res, err := c.DeleteCategory(context.Background(), &pb.DeleteCategoryRequest{Categoryid: int32(deleteCategoryBody.CategoryId)})

	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.DeleteCategoryResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server",
		})
		return
	}
	log.Println("Recieved data : ", res)
	ctx.JSON(http.StatusOK, &res)
}
