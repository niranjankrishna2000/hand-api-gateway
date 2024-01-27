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

// Admin Category List godoc
//
//	@Summary		Admin can see Categories
//	@Description	Admin can see Categories
//	@Tags			Admin Categories
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Param			limit		query		string	false	"limit"
//	@Param			page		query		string	false	"Page number"
//	@Param			searchkey	query		string	false	"searchkey"
//	@Success		200			{object}	pb.CategoryListResponse
//	@Router			/admin/categories/categorylist  [get]
func CategoryList(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	pageStr := ctx.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("page format invalid"))
		return
	}

	limitStr := ctx.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("limit format invalid"))
		return
	}

	searchkey := ctx.Query("searchkey")
	log.Println("Collected data : ", page, limit, searchkey)
	log.Println("Fetching Data...")

	res, err := c.CategoryList(context.Background(), &pb.CategoryListRequest{Page: int32(page), Limit: int32(limit), Searchkey: searchkey})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
