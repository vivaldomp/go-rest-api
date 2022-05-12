package main

import (
	"github.com/vivaldomp/go-rest-api/cmd"
	_ "github.com/vivaldomp/go-rest-api/docs"
)

// @title REST API
// @version 1.0
// @description Restful API description
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cmd.Execute()
}
