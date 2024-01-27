package admin

import (
	"hand/pkg/auth"
	"hand/pkg/config"
	"hand/pkg/admin/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
		Auth: authSvc.Client,
	}

	routes := r.Group("/admin")
	routes.Use(a.AdminAuthRequired)
	routes.GET("/feeds", svc.Feeds)

	post := routes.Group("/post")
	post.GET("/details", svc.PostDetails)
	post.DELETE("/delete",svc.DeletePost)

	users := routes.Group("/users")
	users.GET("/list",svc.UserList)
	users.GET("/details",svc.UserDetails)
	users.PATCH("/changepermission",svc.ChangeUserPermission)

	campaigns:=routes.Group("/campaigns")
	campaigns.GET("/requestlist",svc.CampaignRequestList)
	campaigns.GET("/details",svc.CampaignDetails)
	campaigns.PATCH("/approve",svc.ApproveCampaign)
	campaigns.PATCH("/reject",svc.RejectCampaign)
	
	///////////////////////////////////////////////////////////////
	
	reported:=campaigns.Group("/reported")
	reported.GET("",svc.ReportedList)
	reported.GET("/details",svc.ReportDetails)
	reported.DELETE("/delete",svc.DeleteReport)

	categories:=routes.Group("/categories")
	categories.GET("/categorylist",svc.CategoryList)
	categories.GET("/categorylist/posts",svc.CategoryPosts)
	categories.POST("/new",svc.NewCategory)
	categories.DELETE("/delete",svc.DeleteCategory)

	dashboard:=routes.Group("/dashboard")
	dashboard.GET("",svc.AdminDashboard)
	dashboard.GET("/posts",svc.PostStats)
	//dashboard.GET("/users",svc.UserStats)
	//dashboard.GET("/category",svc.CategoryStats)

}

func (svc *ServiceClient) Feeds(ctx *gin.Context) {
	routes.Feeds(ctx, svc.Client)
}

func (svc *ServiceClient) PostDetails(ctx *gin.Context) {
	routes.PostDetails(ctx, svc.Client)
}

func (svc *ServiceClient) DeletePost(ctx *gin.Context) {
	routes.DeletePost(ctx, svc.Client)
}

func (svc *ServiceClient) UserList(ctx *gin.Context) {
	routes.UserList(ctx, svc.Client,svc.Auth)
}

func (svc *ServiceClient) UserDetails(ctx *gin.Context) {
	routes.UserDetails(ctx, svc.Client,svc.Auth)
}

func (svc *ServiceClient) ChangeUserPermission(ctx *gin.Context) {
	routes.ChangeUserPermission(ctx, svc.Client,svc.Auth)

}

func (svc *ServiceClient) CampaignRequestList(ctx *gin.Context) {
	routes.CampaignRequestList(ctx, svc.Client)

}

func (svc *ServiceClient) CampaignDetails(ctx *gin.Context) {
	routes.CampaignDetails(ctx, svc.Client)

}

func (svc *ServiceClient) ApproveCampaign(ctx *gin.Context) {
	routes.ApproveCampaign(ctx, svc.Client)

}

func (svc *ServiceClient) RejectCampaign(ctx *gin.Context) {
	routes.RejectCampaign(ctx, svc.Client)

}

func (svc *ServiceClient) ReportedList(ctx *gin.Context) {
	routes.ReportedList(ctx, svc.Client)

}

func (svc *ServiceClient) ReportDetails(ctx *gin.Context) {
	routes.ReportDetails(ctx, svc.Client)

}

func (svc *ServiceClient) DeleteReport(ctx *gin.Context) {
	routes.DeleteReport(ctx, svc.Client)

}

func (svc *ServiceClient) CategoryList(ctx *gin.Context) {
	routes.CategoryList(ctx, svc.Client)

}

func (svc *ServiceClient) CategoryPosts(ctx *gin.Context) {
	routes.CategoryPosts(ctx, svc.Client)

}

func (svc *ServiceClient) NewCategory(ctx *gin.Context) {
	routes.NewCategory(ctx, svc.Client)

}
func (svc *ServiceClient) DeleteCategory(ctx *gin.Context) {
	routes.DeleteCategory(ctx, svc.Client)

}

func (svc *ServiceClient) AdminDashboard(ctx *gin.Context) {
	routes.AdminDashboard(ctx, svc.Client)

}

func (svc *ServiceClient) PostStats(ctx *gin.Context) {
	routes.PostStats(ctx, svc.Client)

}
//edit
// func (svc *ServiceClient) UserStats(ctx *gin.Context) {
// 	routes.UserStats(ctx, svc.Client)
// }
// func (svc *ServiceClient) CategoryStats(ctx *gin.Context) {
// 	routes.CategoryStats(ctx, svc.Client)

// }