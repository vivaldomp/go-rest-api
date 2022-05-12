package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Rest API!")
	}
}
