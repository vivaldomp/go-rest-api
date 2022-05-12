package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vivaldomp/go-rest-api/handlers"
)

func homeRouter(router *gin.RouterGroup) {
	router.GET("/", handlers.HomeHandler())
}
