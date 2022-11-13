package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jsocial/internal/pkg/response"
	"net/http"
)

type BaseController struct {
}

func (c *BaseController) JsonSuccess(ctx *gin.Context, data map[string]interface{}, jumpUrl string) {
	ins, _ := response.GetResponseInstance()
	rsp := ins.JsonSuccess(data, jumpUrl)
	ctx.JSON(http.StatusOK, rsp)
}

func (c *BaseController) JsonFail(ctx *gin.Context, code int32, businessCode int32, message string, data map[string]interface{}, jumpUrl string) {
	ins, _ := response.GetResponseInstance()
	rsp := ins.JsonFail(code, businessCode, message, data, jumpUrl)
	ctx.JSON(http.StatusOK, rsp)
}

func (c *BaseController) GetLoginUserId() uint64 {
	return 1
}
