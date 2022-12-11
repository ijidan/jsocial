package feed

import (
	"context"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/constant"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	service.BasicService
	proto_build.UnimplementedFeedServiceServer
}

func (s *Service) FeedCreate(c context.Context, req *proto_build.FeedCreateRequest) (*proto_build.FeedCreateResponse, error) {
	userId := s.GetLoginUserId()
	var protoFeed *proto_build.Feed
	var err error
	switch req.Type {
	case proto_build.FeedType_Txt:
		protoFeed, err = CreateTextFeed(global.GR.Db, userId, req.Content)
	case proto_build.FeedType_Image:
		protoFeed, err = CreateImageFeed(global.GR.Db, userId, req.Content, req.Resource)
	case proto_build.FeedType_Video:
		protoFeed, err = CreateVideoFeed(global.GR.Db, userId, req.Content, req.Resource)
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.FeedCreateResponse{Feed: protoFeed}
	return rsp, nil
}

func (s *Service) FeedEdit(c context.Context, req *proto_build.FeedEditRequest) (*proto_build.FeedEditResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method FeedEdit not implemented")
}

func (s *Service) FeedLike(c context.Context, req *proto_build.FeedLikeRequest) (*proto_build.FeedLikeResponse, error) {
	userId := s.GetLoginUserId()
	err := LikeFeed(global.GR.Db, req.Id, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.FeedLikeResponse{}
	return rsp, nil
}

func (s *Service) FeedUnLike(c context.Context, req *proto_build.FeedUnLikeRequest) (*proto_build.FeedUnLikeResponse, error) {
	userId := s.GetLoginUserId()
	err := UnLikeFeed(global.GR.Db, req.Id, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.FeedUnLikeResponse{}
	return rsp, nil
}

func (s *Service) FeedGet(c context.Context, req *proto_build.FeedGetRequest) (*proto_build.FeedGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedGet not implemented")
}

func (s *Service) FeedDelete(c context.Context, req *proto_build.FeedDeleteRequest) (*proto_build.FeedDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedDelete not implemented")
}

func (s *Service) FeedOwn(c context.Context, req *proto_build.FeedOwnRequest) (*proto_build.FeedOwnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedOwn not implemented")
}

func (s *Service) FeedQuery(c context.Context, req *proto_build.FeedQueryRequest) (*proto_build.FeedQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedSearch not implemented")
}

func (s *Service) FeedFollow(c context.Context, req *proto_build.FeedFollowRequest) (*proto_build.FeedFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedFollow not implemented")
}

func (s *Service) Ping(c context.Context, req *proto_build.PingRequest) (*proto_build.PingResponse, error) {
	rsp := &proto_build.PingResponse{
		Content: constant.ServiceFeed + ":pong",
	}
	return rsp, nil
}

func NewService(cf config.Rpc) *Service {
	instance := &Service{BasicService: service.BasicService{
		Name: constant.ServiceFeed,
		Host: cf.Host,
		Port: cf.Port,
		Ttl:  cf.Ttl,
	}}
	return instance
}
