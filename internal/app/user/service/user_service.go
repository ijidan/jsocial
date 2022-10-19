package service

import (
	"context"
	"github.com/ijidan/jsocial/api/proto_build"
	repository2 "github.com/ijidan/jsocial/internal/app/user/repository"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/funct"
	"github.com/ijidan/jsocial/internal/pkg/global"
	"github.com/ijidan/jsocial/internal/pkg/jwt"
	"github.com/ijidan/jsocial/internal/pkg/repository"
	"github.com/ijidan/jsocial/internal/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	service.BasicService
	proto_build.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(c context.Context, req *proto_build.UserCreateRequest) (*proto_build.UserCreateResponse, error) {
	if req.Password != req.PasswordRpt {
		return nil, status.Error(codes.Internal, "password ,password_rpt not the same")
	}
	gender := repository2.GenderProtoToDb[req.Gender]
	user, err := repository2.CreateUser(global.Db, req.Nickname, req.Password, gender, req.AvatarUrl, "")
	if err != nil {
		return nil, err
	}
	go repository.PublishNewUser(user.ID)
	protoUser, err1 := repository2.GetProtoUserByUserId(global.Db, user.ID)
	if err != nil {
		return nil, err1
	}
	rsp := &proto_build.UserCreateResponse{User: protoUser}
	defer s.AddSpan(c, funct.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *UserService) GetUser(c context.Context, req *proto_build.UserGetRequest) (*proto_build.UserGetResponse, error) {
	id := req.GetId()
	protoUser, err := repository2.GetProtoUserByUserId(global.Db, id)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.UserGetResponse{
		User: protoUser,
	}
	defer s.AddSpan(c, funct.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *UserService) QueryUser(c context.Context, req *proto_build.UserQueryRequest) (*proto_build.UserQueryResponse, error) {
	protoUserList, pager, err := repository2.QueryProtoUser(global.Db, req.GetKeyword(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.UserQueryResponse{
		Pager: pager,
		User:  protoUserList,
	}
	return rsp, nil
}

func (s *UserService) UserLogin(c context.Context, req *proto_build.UserLoginRequest) (*proto_build.UserLoginResponse, error) {
	protoUser, err := repository2.GetProtoUser(global.Db, req.GetNickname(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	userId := protoUser.Id
	token, _ := jwt.GenJwtToken(userId, global.Conf.Jwt.Secret)
	rsp := &proto_build.UserLoginResponse{Token: token}
	return rsp, nil
}

func (s *UserService) UpdatePassword(c context.Context, req *proto_build.UpdatePasswordRequest) (*proto_build.UpdatePasswordResponse, error) {
	if req.Password != req.PasswordRpt {
		return nil, status.Error(codes.Internal, "password ,password_rpt not the same")
	}
	userId := s.GetLoginUserId()
	err := repository2.UpdateUserPassword(global.Db, userId, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.UpdatePasswordResponse{}
	return rsp, nil
}

func (s *UserService) UpdateAvatar(c context.Context, req *proto_build.UpdateAvatarRequest) (*proto_build.UpdateAvatarResponse, error) {
	url := req.GetUrl()
	userId := s.GetLoginUserId()
	err := repository2.UpdateUserAvatar(global.Db, userId, url)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &proto_build.UpdateAvatarResponse{}
	return rsp, nil
}

func NewUserService(cf config.Rpc) *UserService {
	instance := &UserService{BasicService: service.BasicService{
		Name: "service_ping",
		Host: cf.Host,
		Port: cf.Port,
		Ttl:  cf.Ttl,
	}}
	return instance
}
