package controller

import (
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	BaseController
}

func (c *HealthController) Ping(ctx *gin.Context) {
	data := map[string]interface{}{"message": "pong"}
	c.JsonSuccess(ctx, data, "")
}
