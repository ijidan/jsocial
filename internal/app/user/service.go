package user

import (
	"context"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/constant"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/funct"
	"github.com/ijidan/jsocial/internal/pkg/jwt"
	"github.com/ijidan/jsocial/internal/service"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	service.BasicService
	proto_build.UnimplementedUserServiceServer
}

func (s *Service) CreateUser(c context.Context, req *proto_build.UserCreateRequest) (*proto_build.UserCreateResponse, error) {
	if req.Password != req.PasswordRpt {
		return nil, status.Error(codes.Internal, "password ,password_rpt not the same")
	}
	gender := GenderProtoToDb[req.Gender]
	user, err := CreateUser(global.GR.Db, req.Nickname, req.Password, gender, req.AvatarUrl, "")
	if err != nil {
		return nil, err
	}
	//go repository.PublishNewUser(user.ID)
	protoUser, err1 := GetProtoUserByUserId(global.GR.Db, cast.ToUint64(user.ID))
	if err != nil {
		return nil, err1
	}
	rsp := &proto_build.UserCreateResponse{User: protoUser}
	defer s.AddSpan(c, funct.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *Service) GetUser(c context.Context, req *proto_build.UserGetRequest) (*proto_build.UserGetResponse, error) {
	id := req.GetId()
	protoUser, err := GetProtoUserByUserId(global.GR.Db, id)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.UserGetResponse{
		User: protoUser,
	}
	defer s.AddSpan(c, funct.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *Service) QueryUser(c context.Context, req *proto_build.UserQueryRequest) (*proto_build.UserQueryResponse, error) {
	protoUserList, pager, err := QueryProtoUser(global.GR.Db, req.GetKeyword(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.UserQueryResponse{
		Pager: pager,
		User:  protoUserList,
	}
	return rsp, nil
}

func (s *Service) UserLogin(c context.Context, req *proto_build.UserLoginRequest) (*proto_build.UserLoginResponse, error) {
	protoUser, err := GetProtoUser(global.GR.Db, req.GetNickname(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	userId := protoUser.Id
	token, _ := jwt.GenJwtToken(userId, global.GR.Conf.Jwt.Secret)
	rsp := &proto_build.UserLoginResponse{Token: token}
	return rsp, nil
}

func (s *Service) UpdatePassword(c context.Context, req *proto_build.UpdatePasswordRequest) (*proto_build.UpdatePasswordResponse, error) {
	if req.Password != req.PasswordRpt {
		return nil, status.Error(codes.Internal, "password ,password_rpt not the same")
	}
	userId := s.GetLoginUserId()
	err := UpdateUserPassword(global.GR.Db, userId, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.UpdatePasswordResponse{}
	return rsp, nil
}

func (s *Service) UpdateAvatar(c context.Context, req *proto_build.UpdateAvatarRequest) (*proto_build.UpdateAvatarResponse, error) {
	url := req.GetUrl()
	userId := s.GetLoginUserId()
	err := UpdateUserAvatar(global.GR.Db, userId, url)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.UpdateAvatarResponse{}
	return rsp, nil
}

func (s *Service) Ping(c context.Context, req *proto_build.PingRequest) (*proto_build.PingResponse, error) {
	rsp := &proto_build.PingResponse{
		Content: constant.ServiceUser + ":pong",
	}
	return rsp, nil
}

func NewService(cf config.Rpc) *Service {
	instance := &Service{BasicService: service.BasicService{
		Name: constant.ServiceUser,
		Host: cf.Host,
		Port: cf.Port,
		Ttl:  cf.Ttl,
	}}
	return instance
}
