package routes

import (
	"context"
	"errors"
	"hand/pkg/admin/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Admin Delete Category godoc
//
// @Summary  Admin can Delete Category
// @Description Admin can Delete Category
// @Tags   Admin Categories
// @Accept   json
// @Produce  json
// @Param   categoryId query  string true "Category ID "
// @Success  200   {object} pb.DeleteCategoryResponse
// @Router   /admin/categories/delete  [delete]
func DeleteCategory(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	catIDstr := ctx.Query("categoryId")
	catID, err := strconv.Atoi(catIDstr)
	if err != nil || catID < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("id format invalid"))
		return
	}
	log.Println("Collected data : ", catID)

	res, err := c.DeleteCategory(context.Background(), &pb.DeleteCategoryRequest{Categoryid: int32(catID)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)
	ctx.JSON(http.StatusOK, &res)
}
