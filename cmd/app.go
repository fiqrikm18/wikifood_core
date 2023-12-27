package cmd

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/fiqrikm18/wikifood/core/internal/presenter/router/api"

	_ "github.com/fiqrikm18/wikifood/core/docs"
)

// @title           Wikifood Core
// @version         0.1
// @description

// @contact.name   Fiqri Khoirul Muttaqin
// @contact.email  fiqrikm18@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func RunServer() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	api.NewRouter(v1)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
