package routes

import (
	"context"
	"errors"
	"hand/pkg/admin/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin Dashboard godoc
//
//	@Summary		Admin can see website statistics
//	@Description	Admin can see website statistics
//	@Tags			Admin Dashboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	pb.AdminDashboardResponse
//	@Router			/admin/dashboard  [get]
func AdminDashboard(ctx *gin.Context, c pb.AdminServiceClient) {
	log.Println("Initiating AdminDashboard...")

	log.Println("Fetching Data...")
	res, err := c.AdminDashboard(context.Background(), &pb.AdminDashboardRequest{})
	if err != nil {
		log.Println("Error occured while fetching data...", err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": http.StatusBadGateway,
			"error":  errors.New("failed to fetch admin dashboard"),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, &res)
}
