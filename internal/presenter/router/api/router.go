package api

import (
	"github.com/fiqrikm18/wikifood/core/internal/presenter/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.RouterGroup) {
	router.GET("/healthcheck", controller.HealthCheckController)
}
