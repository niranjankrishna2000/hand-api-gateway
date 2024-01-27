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

// Admin Category Stats godoc
//
// @Summary  Admin can see Category toplist
// @Description Admin can see Category toplist
// @Tags   Admin Dashboard
// @Accept   json
// @Produce  json
// @Param			limit	query		string	false	"limit"
// @Param			page	query		string	false	"Page number"
// @Success  200   {object} pb.CategoryStatsResponse
// @Router   /admin/dashboard/category  [get]
func CategoryStats(ctx *gin.Context, c pb.AdminServiceClient) {
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
	log.Println("Collected data : ", page, limit)
	log.Println("Fetching Data...")

	res, err := c.CategoryStats(context.Background(), &pb.CategoryStatsRequest{Page: int32(page),Limit: int32(limit)})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
