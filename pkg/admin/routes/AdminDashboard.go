package routes

import (
	"context"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin Dashboard godoc
//
//	@Summary		Admin Dashboard
//	@Description	Admin can see website statistics
//	@Tags			Admin Dashboard
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	pb.AdminDashboardResponse
//
//	@Failure		502	{object}	pb.AdminDashboardResponse
//
//	@Router			/admin/dashboard  [get]
func AdminDashboard(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	log.Println("Fetching Data...")
	res, err := c.AdminDashboard(context.Background(), &pb.AdminDashboardRequest{})
	if err != nil {
		log.Println("Error occured while fetching data...", err)
		ctx.JSON(http.StatusBadGateway, &pb.AdminDashboardResponse{
			Status:   http.StatusBadGateway,
			Response: "Error from internal server",
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
