package main

import (
	"log"

	"hand/pkg/config"

	"hand/pkg/admin"
	"hand/pkg/auth"
	"hand/pkg/user"

	docs "hand/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@SecurityDefinition				BearerAuth
//	@TokenUrl						/auth/token
//	@securityDefinitions.Bearer		type apiKey
//	@securityDefinitions.Bearer		name Authorization
//	@securityDefinitions.Bearer		in header
//	@securityDefinitions.BasicAuth	type basic

//	@title			Hand crowdfunding API
//	@version		1.0
//	@description	API for crowdfunding website

//	@securityDefinitions.apiKey	api_key
//	@in							header
//	@name						Authorization

//	@contact.name	Contact Developer
//	@contact.email	niranjankrishna.thundathil@gmail.com

//	@host		www.handcrowdfunding.online
//	@BasePath	/

// @schemes	http,https
func main() {

	docs.SwaggerInfo.Title = "Hand CrowdFunding API"
	docs.SwaggerInfo.Description = "Helping hand for all"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "handcrowdfunding.online"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.LoadHTMLGlob("pkg/templates/*.html")
	authSvc := *auth.RegisterRoutes(r, &c)
	user.RegisterRoutes(r, &c, &authSvc)
	admin.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
