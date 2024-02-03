package user

import (
	"hand/pkg/auth"
	"hand/pkg/config"
	"hand/pkg/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
		Auth:   authSvc.Client,
	}

	routes := r.Group("/user")
	routes.Use(a.AuthRequired)
	routes.GET("/feeds", svc.UserFeeds)
	///////////////////////////////////////////////////////
	post := routes.Group("/post")
	post.POST("/new", svc.CreatePost)
	post.GET("/new", svc.GetCreatePost)
	post.POST("/upload-image", svc.UploadImage)
	post.POST("/upload-multiple-image") //beta
	post.GET("/details", svc.PostDetails)
	post.PATCH("/edit", svc.EditPost)
	post.PATCH("/expire",svc.ExpirePost)
	post.DELETE("/delete",svc.DeletePost)
	post.POST("/details/report", svc.ReportPost)
	post.POST("/like", svc.LikePost)
	post.POST("/comment", svc.CommentPost)
	post.POST("/comment/report", svc.ReportComment)
	post.DELETE("/comment/delete", svc.DeleteComment)
	post.GET("/updates",svc.GetUpdates)
	post.POST("/updates",svc.AddUpdate)
	post.PATCH("/updates",svc.EditUpdate)
	post.DELETE("/updates",svc.DeleteUpdate)
	///////////////////////////////////////////////////////
	donate := post.Group("/donate")
	donate.POST("/", svc.Donate)
	r.GET("/user/post/donate/razorpay", svc.MakePaymentRazorPay)
	donate.POST("/generate-invoice", svc.GenerateInvoice)
	donate.GET("/download-invoice", svc.DownloadInvoice)
	donate.GET("/history", svc.DonationHistory)
	///////////////////////////////////////////////////////
	profile := routes.Group("/profile")
	profile.GET("/details", svc.ProfileDetails) // add full details
	profile.PATCH("/edit", svc.EditProfile)
	profile.GET("/monthly-goal",svc.GetMonthlyGoal)
	profile.POST("/monthly-goal",svc.AddMonthlyGoal)
	profile.PUT("/monthly-goal",svc.EditMonthlyGoal)
	profile.GET("/my-campaigns",svc.GetMyCampaigns)
	profile.GET("/my-impact",svc.GetmyImpact)
	///////////////////////////////////////////////////////
	notification := routes.Group("/notifications")
	notification.GET("", svc.Notifications)
	notification.DELETE("/delete", svc.DeleteNotification)
	notification.DELETE("/clear", svc.ClearNotification)
	///////////////////////////////////////////////////////
	success := routes.Group("/success-stories")
	success.GET("",svc.GetSuccessStory)
	success.POST("",svc.AddSuccessStory)
	success.PATCH("",svc.EditSuccessStory)
	success.DELETE("",svc.DeleteSuccessStory)
	///////////////////////////////////////////////////////
}

func (svc *ServiceClient) UserFeeds(ctx *gin.Context) {
	routes.UserFeeds(ctx, svc.Client)
}

func (svc *ServiceClient) UploadImage(ctx *gin.Context) {
	routes.UploadImage(ctx, svc.Client)
}

func (svc *ServiceClient) CreatePost(ctx *gin.Context) {
	routes.CreatePost(ctx, svc.Client)
}
func (svc *ServiceClient) ExpirePost(ctx *gin.Context) {
	routes.ExpirePost(ctx, svc.Client)
}
func (svc *ServiceClient) DeletePost(ctx *gin.Context) {
	routes.DeletePost(ctx, svc.Client)
}
func (svc *ServiceClient) GetCreatePost(ctx *gin.Context) {
	routes.GetCreatePost(ctx, svc.Client)
}
func (svc *ServiceClient) PostDetails(ctx *gin.Context) {
	routes.UserPostDetails(ctx, svc.Client)
}

func (svc *ServiceClient) EditPost(ctx *gin.Context) {
	routes.EditPost(ctx, svc.Client)
}
func (svc *ServiceClient) ReportPost(ctx *gin.Context) {
	routes.ReportPost(ctx, svc.Client)
}
func (svc *ServiceClient) Donate(ctx *gin.Context) {
	routes.Donate(ctx, svc.Client)
}

func (svc *ServiceClient) MakePaymentRazorPay(ctx *gin.Context) {
	routes.MakePaymentRazorPay(ctx, svc.Client)
}
func (svc *ServiceClient) DownloadInvoice(ctx *gin.Context) {
	routes.DownloadInvoice(ctx, svc.Client)
}
func (svc *ServiceClient) GenerateInvoice(ctx *gin.Context) {
	routes.GenerateInvoice(ctx, svc.Client, svc.Auth)
}

func (svc *ServiceClient) LikePost(ctx *gin.Context) {
	routes.LikePost(ctx, svc.Client)
}
func (svc *ServiceClient) CommentPost(ctx *gin.Context) {
	routes.CommentPost(ctx, svc.Client)
}
func (svc *ServiceClient) ReportComment(ctx *gin.Context) {
	routes.ReportComment(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteComment(ctx *gin.Context) {
	routes.DeleteComment(ctx, svc.Client)
}

func (svc *ServiceClient) DonationHistory(ctx *gin.Context) {
	routes.DonationHistory(ctx, svc.Client)
}

// edit
func (svc *ServiceClient) ProfileDetails(ctx *gin.Context) {
	routes.ProfileDetails(ctx, svc.Client)
}
func (svc *ServiceClient) EditProfile(ctx *gin.Context) {
	routes.EditProfile(ctx, svc.Client)
}
func (svc *ServiceClient) Notifications(ctx *gin.Context) {
	routes.Notifications(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteNotification(ctx *gin.Context) {
	routes.DeleteNotification(ctx, svc.Client)
}
func (svc *ServiceClient) ClearNotification(ctx *gin.Context) {
	routes.ClearNotification(ctx, svc.Client)
}

//updates
func (svc *ServiceClient) GetUpdates(ctx *gin.Context) {
	routes.GetUpdates(ctx, svc.Client)
}
func (svc *ServiceClient) AddUpdate(ctx *gin.Context) {
	routes.AddUpdate(ctx, svc.Client)
}
func (svc *ServiceClient) EditUpdate(ctx *gin.Context) {
	routes.EditUpdate(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteUpdate(ctx *gin.Context) {
	routes.DeleteUpdate(ctx, svc.Client)
}
//monthly goal
func (svc *ServiceClient) GetMonthlyGoal(ctx *gin.Context) {
	routes.GetMonthlyGoal(ctx, svc.Client)
}
func (svc *ServiceClient) AddMonthlyGoal(ctx *gin.Context) {
	routes.AddMonthlyGoal(ctx, svc.Client)
}
func (svc *ServiceClient) EditMonthlyGoal(ctx *gin.Context) {
	routes.EditMonthlyGoal(ctx, svc.Client)
}
//my impact
func (svc *ServiceClient) GetmyImpact(ctx *gin.Context) {
	routes.GetmyImpact(ctx, svc.Client)
}
//my campaigns
func (svc *ServiceClient) GetMyCampaigns(ctx *gin.Context) {
	routes.GetMyCampaigns(ctx, svc.Client)
}
//success stories
func (svc *ServiceClient) GetSuccessStory(ctx *gin.Context) {
	routes.GetSuccessStory(ctx, svc.Client)
}
func (svc *ServiceClient) AddSuccessStory(ctx *gin.Context) {
	routes.AddSuccessStory(ctx, svc.Client)
}
func (svc *ServiceClient) EditSuccessStory(ctx *gin.Context) {
	routes.EditSuccessStory(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteSuccessStory(ctx *gin.Context) {
	routes.DeleteSuccessStory(ctx, svc.Client)
}

/////////////////////////////////////////////////////////////////////////////////////////
//note:
// ** try chats
// ** consider the omitempty, probably in posts and feeds
// ** use of actual location
 