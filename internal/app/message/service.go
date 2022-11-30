package message

import (
	"context"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/service"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	service.BasicService
	proto_build.UnimplementedMessageServiceServer
}

func (s *Service) SendUserTextMessage(c context.Context, req *proto_build.SendUserTextMessageRequest) (*proto_build.SendUserTextMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateUserTextMessage(global.GR.Db, userId, req.ToUserId, req.Text)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendUserTextMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendUserLocationMessage(c context.Context, req *proto_build.SendUserLocationMessageRequest) (*proto_build.SendUserLocationMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateUserLocationMessage(global.GR.Db, userId, req.ToUserId, req.Location)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendUserLocationMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendUserFaceMessage(c context.Context, req *proto_build.SendUserFaceMessageRequest) (*proto_build.SendUserFaceMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateUserFaceMessage(global.GR.Db, userId, req.ToUserId, req.Face)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendUserFaceMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendUserSoundMessage(c context.Context, req *proto_build.SendUserSoundMessageRequest) (*proto_build.SendUserSoundMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateUserSoundMessage(global.GR.Db, userId, req.ToUserId, req.Sound)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendUserSoundMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendUserVideoMessage(c context.Context, req *proto_build.SendUserVideoMessageRequest) (*proto_build.SendUserVideoMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateUserVideoMessage(global.GR.Db, userId, req.ToUserId, req.Video)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendUserVideoMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendUserImageMessage(c context.Context, req *proto_build.SendUserImageMessageRequest) (*proto_build.SendUserImageMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateUserImageMessage(global.GR.Db, userId, req.ToUserId, req.Image)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendUserImageMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendUserFileMessage(c context.Context, req *proto_build.SendUserFileMessageRequest) (*proto_build.SendUserFileMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateUserFileMessage(global.GR.Db, userId, req.ToUserId, req.File)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendUserFileMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendGroupTextMessage(c context.Context, req *proto_build.SendGroupTextMessageRequest) (*proto_build.SendGroupTextMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateGroupTextMessage(global.GR.Db, userId, req.ToGroupId, req.AtUserId, req.Text)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendGroupTextMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendGroupLocationMessage(c context.Context, req *proto_build.SendGroupLocationMessageRequest) (*proto_build.SendGroupLocationMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateGroupLocationMessage(global.GR.Db, userId, req.ToGroupId, req.AtUserId, req.Location)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendGroupLocationMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendGroupFceMessage(c context.Context, req *proto_build.SendGroupFaceMessageRequest) (*proto_build.SendGroupFaceMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateGroupFaceMessage(global.GR.Db, userId, req.ToGroupId, req.AtUserId, req.Face)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendGroupFaceMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendGroupSoundMessage(c context.Context, req *proto_build.SendGroupSoundMessageRequest) (*proto_build.SendGroupSoundMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateGroupSoundMessage(global.GR.Db, userId, req.ToGroupId, req.AtUserId, req.Sound)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendGroupSoundMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendGroupVideoMessage(c context.Context, req *proto_build.SendGroupVideoMessageRequest) (*proto_build.SendGroupVideoMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateGroupVideoMessage(global.GR.Db, userId, req.ToGroupId, req.AtUserId, req.Video)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendGroupVideoMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendGroupImageMessage(c context.Context, req *proto_build.SendGroupImageMessageRequest) (*proto_build.SendGroupImageMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateGroupImageMessage(global.GR.Db, userId, req.ToGroupId, req.AtUserId, req.Image)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendGroupImageMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func (s *Service) SendGroupFileMessage(c context.Context, req *proto_build.SendGroupFileMessageRequest) (*proto_build.SendGroupFileMessageResponse, error) {
	userId := s.GetLoginUserId()
	message, err := CreateGroupFileMessage(global.GR.Db, userId, req.ToGroupId, req.AtUserId, req.File)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.SendGroupFileMessageResponse{Id: cast.ToUint64(message.ID)}
	return rsp, nil
}

func NewService(cf config.Rpc) *Service {
	instance := &Service{BasicService: service.BasicService{
		Name: "service_message",
		Host: cf.Host,
		Port: cf.Port,
		Ttl:  cf.Ttl,
	}}
	return instance
}
