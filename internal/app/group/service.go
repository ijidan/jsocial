package group

import (
	"context"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/funct"
	"github.com/ijidan/jsocial/internal/service"
	"github.com/spf13/cast"
)

type Service struct {
	service.BasicService
	proto_build.UnimplementedGroupServiceServer
}

func (s *Service) CreateGroup(c context.Context, req *proto_build.CreateGroupRequest) (*proto_build.CreateGroupResponse, error) {
	userId := s.GetLoginUserId()
	group, err := CreateGroup(global.GR.Db, userId, req.Name, req.Introduction, req.AvatarUrl)
	if err != nil {
		return nil, err
	}
	protoGroup, err1 := GetProtoGroupByGroupId(global.GR.Db, cast.ToUint64(group.ID))
	if err != nil {
		return nil, err1
	}
	rsp := &proto_build.CreateGroupResponse{Group: protoGroup}
	defer s.AddSpan(c, funct.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *Service) UpdateGroup(c context.Context, req *proto_build.UpdateGroupRequest) (*proto_build.UpdateGroupResponse, error) {
	err := UpdateGroup(global.GR.Db, req.Id, req.Name, req.Introduction, req.AvatarUrl)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.UpdateGroupResponse{}
	return rsp, nil
}

func (s *Service) GetGroup(c context.Context, req *proto_build.GetGroupRequest) (*proto_build.GetGroupResponse, error) {
	group, err := GetProtoGroup(global.GR.Db, req.Id)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.GetGroupResponse{Group: group}
	return rsp, nil
}

func (s *Service) QueryGroup(c context.Context, req *proto_build.QueryGroupRequest) (*proto_build.QueryGroupResponse, error) {
	protoGroupList, pager, err := QueryProtoGroup(global.GR.Db, req.GetKeyword(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.QueryGroupResponse{
		Pager: pager,
		Group: protoGroupList,
	}
	return rsp, nil
}

func (s *Service) DeleteGroup(c context.Context, req *proto_build.DeleteGroupRequest) (*proto_build.DeleteGroupResponse, error) {
	err := DeleteGroup(global.GR.Db, req.Id)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.DeleteGroupResponse{}
	return rsp, nil
}

func (s *Service) JoinGroup(c context.Context, req *proto_build.JoinGroupRequest) (*proto_build.JoinGroupResponse, error) {
	err := JoinGroup(global.GR.Db, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.JoinGroupResponse{}
	return rsp, nil
}

func (s *Service) QuitGroup(c context.Context, req *proto_build.QuitGroupRequest) (*proto_build.QuitGroupResponse, error) {
	err := QuitGroup(global.GR.Db, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.QuitGroupResponse{}
	return rsp, nil
}

func NewService(cf config.Rpc) *Service {
	instance := &Service{BasicService: service.BasicService{
		Name: "service_group",
		Host: cf.Host,
		Port: cf.Port,
		Ttl:  cf.Ttl,
	}}
	return instance
}
