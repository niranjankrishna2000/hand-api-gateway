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
	post.POST("/details/report", svc.ReportPost)
	post.POST("/like", svc.LikePost)
	post.POST("/comment", svc.CommentPost)
	post.POST("/comment/report", svc.ReportComment)
	post.DELETE("/comment/delete", svc.DeleteComment)
	post.GET("/updates")
	post.POST("/updates")
	post.PATCH("/updates")
	post.DELETE("/updates")
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
	profile.GET("/monthly-goal")
	profile.POST("/monthly-goal")
	profile.PATCH("/monthly-goal")
	profile.GET("/my-campaigns")
	profile.GET("/my-impact")
	///////////////////////////////////////////////////////
	notification := routes.Group("/notifications")
	notification.GET("", svc.Notifications)
	notification.GET("/details", svc.NotificationDetail)
	notification.DELETE("/delete", svc.DeleteNotification)
	notification.DELETE("/clear", svc.ClearNotification)
	///////////////////////////////////////////////////////
	success := routes.Group("/success-stories")
	success.GET("")
	success.POST("")
	success.PATCH("")
	success.DELETE("")
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
	//routes.ProfileDetails(ctx, svc.Client)
}
func (svc *ServiceClient) EditProfile(ctx *gin.Context) {
	//routes.ReportComment(ctx, svc.Client)
}
func (svc *ServiceClient) Notifications(ctx *gin.Context) {
	routes.Notifications(ctx, svc.Client)
}
func (svc *ServiceClient) NotificationDetail(ctx *gin.Context) {
	routes.NotificationDetail(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteNotification(ctx *gin.Context) {
	routes.DeleteNotification(ctx, svc.Client)
}
func (svc *ServiceClient) ClearNotification(ctx *gin.Context) {
	routes.ClearNotification(ctx, svc.Client)
}

/////////////////////////////////////////////////////////////////////////////////////////
//note:
// ** try chats
// ** consider the omitempty, probably in posts and feeds
// ** use of actual location
// ** feeds filter  with trending-successfull-taxbenefit-urgent, category, location
// ** tax benefit for campaign
// ** user profile- my donations, my campaigns, your impact and complete profile
// ** show donations on campaign details
// ** campaign updates and supporters
