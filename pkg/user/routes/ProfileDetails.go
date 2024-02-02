package routes

import (
	"context"
	"hand/pkg/user/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profile Detail godoc
//
//	@Summary		Profile Details
//	@Description	User can see Profile Details
//	@Tags			User Profile
//	@Security		api_key
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	pb.ProfileDetailsRequest
//	@Failure		403	{string}	string	"You have not logged in"
//	@Failure		502	{object}	pb.ProfileDetailsRequest
//	@Router			/user/profile/details  [get]
func ProfileDetails(ctx *gin.Context, c pb.UserServiceClient) {
	res, err := c.ProfileDetails(context.Background(), &pb.ProfileDetailsRequest{
		Userid: int32(ctx.GetInt64("userId")),
	})
	if err != nil {
		log.Println("Error with internal server :", err)
		ctx.JSON(http.StatusBadGateway, pb.ProfileDetailsResponse{
			Status:       http.StatusBadGateway,
			Response:     "Error in internal server",
			User: &pb.UserProfile{},
		})
		return
	}
	log.Println("Recieved Data: ", res)
	ctx.JSON(http.StatusOK, &res)
}
