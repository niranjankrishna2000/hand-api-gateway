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

// Admin Reported List godoc
//
// @Summary  Admin can see reported posts
// @Description Admin can see reported posts
// @Tags   Admin Reported
// @Accept   json
// @Produce  json
// @Param			limit	query		string	false	"limit"
// @Param			page	query		string	false	"Page number"
// @Param			searchkey	query		string	false	"searchkey"
// @Success  200   {object} pb.ReportedListResponse
// @Router   /admin/campaigns/reported  [get]
func ReportedList(ctx *gin.Context, c pb.AdminServiceClient) {
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

	res, err := c.ReportedList(context.Background(), &pb.ReportedListRequest{Page: int32(page), Limit: int32(limit), Searchkey: searchkey})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	log.Println("Recieved data : ", res)

	ctx.JSON(http.StatusOK, &res)
}
