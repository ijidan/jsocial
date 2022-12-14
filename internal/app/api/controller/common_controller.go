package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/app/api/request"
	"github.com/ijidan/jsocial/internal/service"
	"io"
	"io/ioutil"
	"net/http"
)

type CommonController struct {
	BaseController
}

func (c *CommonController) UploadImage(ctx *gin.Context) {
	var req request.CommonUploadImageRequest
	if err := ctx.ShouldBind(&req); err != nil {
		c.JsonFail(ctx, http.StatusBadRequest, 0, err.Error(), nil, "")
		return
	}
	file, err1 := ctx.FormFile("file")
	if err1 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err1.Error(), nil, "")
		return
	}
	reader, err2 := file.Open()
	if err2 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err2.Error(), nil, "")
		return
	}
	fileBytes, err3 := ioutil.ReadAll(reader)
	if err3 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err3.Error(), nil, "")
		return
	}
	serviceClient, err := service.GetServiceCommonClient()
	if err != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err.Error(), nil, "")
		return
	}
	uploadImageClient, err5 := serviceClient.UploadImage(ctx)
	if err5 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err5.Error(), nil, "")
		return
	}
	serviceReq := &proto_build.UploadImageRequest{
		Content: fileBytes,
	}
	serviceErr := uploadImageClient.Send(serviceReq)
	if serviceErr != nil {
		if serviceErr == io.EOF {
			err6 := uploadImageClient.CloseSend()
			if err6 != nil {
				c.JsonFail(ctx, http.StatusInternalServerError, 0, err6.Error(), nil, "")
				return
			}
		} else {
			c.JsonFail(ctx, http.StatusInternalServerError, 0, serviceErr.Error(), nil, "")
			return
		}
	}
	serviceRsp, err7 := uploadImageClient.CloseAndRecv()
	if err7 != nil {
		c.JsonFail(ctx, http.StatusInternalServerError, 0, err7.Error(), nil, "")
		return
	}
	rsp := map[string]interface{}{"url": serviceRsp.Url}
	c.JsonSuccess(ctx, rsp, "")
	return
}
