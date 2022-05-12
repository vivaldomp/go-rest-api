package v1

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/vivaldomp/go-rest-api/handlers/v1"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	routerGroup := r.Group("/users")
	routerGroup.GET("/:id", handlers.GetUserByID())
}
