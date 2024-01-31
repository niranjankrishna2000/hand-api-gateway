package auth

import (
	"github.com/gin-gonic/gin"
	"hand/pkg/auth/routes"
	"hand/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	//routes := r.Group("/auth")
	r.POST("/signup", svc.SignUp)
	r.POST("/login", svc.Login)
	r.POST("/adminlogin", svc.AdminLogin)
	r.POST("/forgot-password", svc.LoginWithOtp)
	r.PATCH("/otp-validate", svc.OtpValidate)
	r.POST("/logout", svc.Logout)

	return svc
}

func (svc *ServiceClient) SignUp(ctx *gin.Context) {
	routes.SignUp(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
func (svc *ServiceClient) AdminLogin(ctx *gin.Context) {
	routes.AdminLogin(ctx, svc.Client)
}
func (svc *ServiceClient) LoginWithOtp(ctx *gin.Context) {
	routes.LoginWithOtp(ctx, svc.Client)
}

func (svc *ServiceClient) OtpValidate(ctx *gin.Context) {
	routes.OtpValidate(ctx, svc.Client)
}
func (svc *ServiceClient) Logout(ctx *gin.Context) {
	routes.Logout(ctx, svc.Client)
}
