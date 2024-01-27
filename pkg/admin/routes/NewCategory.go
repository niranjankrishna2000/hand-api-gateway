package routes

import (
	"context"
	"errors"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin New Category godoc
//
// @Summary  Admin can create New Category
// @Description Admin can create New Category
// @Tags   Admin Categories
// @Accept   json
// @Produce  json
// @Param   category query  string true "category name"
// @Success  200   {object} pb.NewCategoryResponse
// @Router   /admin/categories/new       [post]
func NewCategory(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	category := ctx.Query("category")
	if category == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid Format"))
		return
	}
	log.Println("Collected Data: ", category)
	
	res, err := c.NewCategory(context.Background(), &pb.NewCategoryRequest{Category: category})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved Data: ", res)

	ctx.JSON(http.StatusOK, &res)
}
