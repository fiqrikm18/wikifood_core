package controller

import "github.com/gin-gonic/gin"

// Healt Check
// @Summary      Check if server is working
// @Description  Check if server is working
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {string} Hello
// @Router       /healthcheck [get]
func HealthCheckController(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}
