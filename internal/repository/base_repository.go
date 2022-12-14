package repository

import (
	"errors"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/pkg/pager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func ConvertPagerToProtoPager(pager *pager.Pager) *proto_build.Pager {
	protoPager := &proto_build.Pager{
		Page:       pager.Page,
		PageZie:    pager.PageSize,
		TotalRows:  pager.TotalRows,
		TotalPages: pager.TotalPages,
	}
	return protoPager
}

func ConvertModelQueryErrorToGrpcStatusError(err error) error {
	if err == nil {
		return nil
	}
	var code codes.Code
	if errors.Is(err, gorm.ErrRecordNotFound) {
		code = codes.NotFound
	} else {
		code = codes.Internal
	}
	return status.Error(code, err.Error())

}
