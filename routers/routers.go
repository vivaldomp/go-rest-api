package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vivaldomp/go-rest-api/configuration"
	v1 "github.com/vivaldomp/go-rest-api/routers/api/v1"
)

func Init(config *configuration.Config) (*gin.Engine, error) {
	r := gin.Default()
	homeRouter(r.Group("/"))
	apiRouter := r.Group("/api")
	apiV1Router := apiRouter.Group("/v1")
	v1.RegisterUserRoutes(apiV1Router)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r, nil
}
