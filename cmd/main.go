package main

import (
	"log"

	"hand/pkg/config"

	"hand/pkg/admin"
	"hand/pkg/auth"
	"hand/pkg/user"

	"github.com/gin-gonic/gin"
	//docs "hand/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @SecurityDefinition BearerAuth
// @TokenUrl /auth/token
// @securityDefinitions.Bearer		type apiKey
// @securityDefinitions.Bearer		name Authorization
// @securityDefinitions.Bearer		in header
// @securityDefinitions.BasicAuth	type basic
func main() {

	// docs.SwaggerInfo.Title = "Hand CrowdFunding API"
	// docs.SwaggerInfo.Description = "Helping hand for all"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:1111"
	// docs.SwaggerInfo.BasePath = ""
	// //docs.SwaggerInfo.Schemes = []string{"http","https"}
   
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
