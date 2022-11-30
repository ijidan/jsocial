package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jsocial/internal/query"
)

type HealthController struct {
	BaseController
}

func (c *HealthController) Ping(ctx *gin.Context) {
	user, _ := query.User.WithContext(ctx).First()
	data := map[string]interface{}{"message": "pong", "user": user}
	c.JsonSuccess(ctx, data, "")
}
