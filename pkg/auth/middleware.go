package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"hand/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	//token, _ := ctx.Cookie("Authorization")
	authorization := ctx.Request.Header.Get("Authorization")
	log.Println("Authorization checking...")
	log.Println("Authorization", authorization)
	if authorization == "" {
		log.Println("No token found")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	log.Println("res:", res, token)
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// var key models.UserKey = "userID"
	// var val models.UserKey = models.UserKey(res.UserId)

	// context := context.WithValue(ctx, key, val)
	// // Set the context to the request
	// ctx.Request = ctx.Request.WithContext(context)

	ctx.Set("userId", res.UserId)
	id, ok := ctx.Get("userId")
	log.Println("User ID:", id, ok)

	ctx.Next()
}

func (c *AuthMiddlewareConfig) AdminAuthRequired(ctx *gin.Context) {
	//token, _ := ctx.Cookie("Authorization")
	authorization := ctx.Request.Header.Get("Authorization")
	log.Println("Admin Authorization checking...")
	log.Println("Token", authorization)
	if authorization == "" {
		log.Println("No token found")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.AdminValidate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	log.Println("res:", res)
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ctx.Set("adminId", res.UserId)
	//id, ok := ctx.Get("userId")
	//log.Println("User ID:", id, ok)

	ctx.Next()
}
