package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMyImpact godoc
//
//	@Summary		My Impact
//	@Description	User can see their impact on society
//	@Tags			User Profile
//	@Accept			json
//	@Produce		json
//	@Security		api_key
//	@Success		200	{object}	pb.GetmyImpactResponse
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		502	{object}	pb.GetmyImpactResponse
//	@Router			/user/profile/my-impact  [get]
func GetmyImpact(ctx *gin.Context, c pb.UserServiceClient) {
	res, err := c.GetmyImpact(context.Background(), &pb.GetmyImpactRequest{
		UserId: int32(ctx.GetInt64("userid")),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.GetmyImpactResponse{
			Status:   http.StatusBadGateway,
			Response: "Error in internal server:" + err.Error(),
			Likes: 0,
			Views: 0,
			Collected: 0,
			Donated: 0,
			LifesChanged: 0,
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(int(res.Status), &res)
}
