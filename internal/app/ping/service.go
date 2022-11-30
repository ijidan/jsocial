package ping

import (
	"context"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/funct"
	"github.com/ijidan/jsocial/internal/service"
)

// Service Service服务
type Service struct {
	service.BasicService
	proto_build.UnimplementedPingServiceServer
}

func (s *Service) Ping(c context.Context, req *proto_build.PingRequest) (*proto_build.PingResponse, error) {
	rsp := &proto_build.PingResponse{
		Message: "pong",
	}
	defer s.AddSpan(c, funct.GetFuncName(), req, rsp.String())
	return rsp, nil
}

// NewService  获取实例
func NewService(cf config.Rpc) *Service {
	instance := &Service{BasicService: service.BasicService{
		Name: "service_ping",
		Host: cf.Host,
		Port: cf.Port,
		Ttl:  cf.Ttl,
	}}
	return instance
}