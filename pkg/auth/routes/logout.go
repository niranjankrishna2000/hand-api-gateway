package routes

import (
	"net/http"

	"hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

// User Logout godoc
//
//	@Summary		User Logout
//	@Description	User can logout here
//	@Tags			User Auth
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/logout [post]
func Logout(ctx *gin.Context, c pb.AuthServiceClient) {
	ctx.Header("Authorization", "")
	ctx.JSON(http.StatusOK, nil)
}
