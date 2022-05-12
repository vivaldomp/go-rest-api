package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vivaldomp/go-rest-api/dto"
	"github.com/vivaldomp/go-rest-api/services"
)

// CreateUserInput ...
type CreateUserInput struct {
	Name string `json:"name"`
	MCI  int    `json:"mci"`
}

// HTTPError ...
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// CreateUser godoc
// @Summary Create user
// @Description Create user operation
// @Tags user
// @Accept  json
// @Produce  json
// @Param requestData body models.User true "Request operation body"
// @Success 200 {object} models.User
// @Header 200 {string} Authorization "token access"
// @Failure 400,404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Failure default {object} HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/users [post]
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()
		var input CreateUserInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser := dto.User{
			Name: input.Name,
			MCI:  input.MCI,
		}

		err := services.CreateUser(ctx, &newUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, newUser)
	}
}

// GetUserByID godoc
// @Summary Get User by ID
// @Description Get User by ID operation
// @Tags user
// @Accept  json
// @Produce  json
// @Param		id path string false "User ID"
// @Success 200 {object} models.User
// @Header 200 {string} Authorization "token access"
// @Failure 400,404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Failure default {object} HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [get]
func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id := c.Param("id")
		aid, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := services.GetUserByID(ctx, aid)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}

}
