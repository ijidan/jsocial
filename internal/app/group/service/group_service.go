package service

import (
	"context"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/app/group/repository"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/funct"
	"github.com/ijidan/jsocial/internal/pkg/global"
	"github.com/ijidan/jsocial/internal/pkg/service"
)

type GroupService struct {
	service.BasicService
	proto_build.UnimplementedGroupServiceServer
}

func (s *GroupService) CreateGroup(c context.Context, req *proto_build.CreateGroupRequest) (*proto_build.CreateGroupResponse, error) {
	userId := s.GetLoginUserId()
	group, err := repository.CreateGroup(global.Db, userId, req.Name, req.Introduction, req.AvatarUrl)
	if err != nil {
		return nil, err
	}
	protoGroup, err1 := repository.GetProtoGroupByGroupId(global.Db, group.ID)
	if err != nil {
		return nil, err1
	}
	rsp := &proto_build.CreateGroupResponse{Group: protoGroup}
	defer s.AddSpan(c, funct.GetFuncName(), req, rsp.String())
	return rsp, nil
}

func (s *GroupService) UpdateGroup(c context.Context, req *proto_build.UpdateGroupRequest) (*proto_build.UpdateGroupResponse, error) {
	err := repository.UpdateGroup(global.Db, req.Id, req.Name, req.Introduction, req.AvatarUrl)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.UpdateGroupResponse{}
	return rsp, nil
}

func (s *GroupService) GetGroup(c context.Context, req *proto_build.GetGroupRequest) (*proto_build.GetGroupResponse, error) {
	group, err := repository.GetProtoGroup(global.Db, req.Id)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.GetGroupResponse{Group: group}
	return rsp, nil
}

func (s *GroupService) QueryGroup(c context.Context, req *proto_build.QueryGroupRequest) (*proto_build.QueryGroupResponse, error) {
	protoGroupList, pager, err := repository.QueryProtoGroup(global.Db, req.GetKeyword(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.QueryGroupResponse{
		Pager: pager,
		Group: protoGroupList,
	}
	return rsp, nil
}

func (s *GroupService) DeleteGroup(c context.Context, req *proto_build.DeleteGroupRequest) (*proto_build.DeleteGroupResponse, error) {
	err := repository.DeleteGroup(global.Db, req.Id)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.DeleteGroupResponse{}
	return rsp, nil
}

func (s *GroupService) JoinGroup(c context.Context, req *proto_build.JoinGroupRequest) (*proto_build.JoinGroupResponse, error) {
	err := repository.JoinGroup(global.Db, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.JoinGroupResponse{}
	return rsp, nil
}

func (s *GroupService) QuitGroup(c context.Context, req *proto_build.QuitGroupRequest) (*proto_build.QuitGroupResponse, error) {
	err := repository.QuitGroup(global.Db, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}
	rsp := &proto_build.QuitGroupResponse{}
	return rsp, nil
}

func NewGroupService(cf config.Rpc) *GroupService {
	instance := &GroupService{BasicService: service.BasicService{
		Name: "service_group",
		Host: cf.Host,
		Port: cf.Port,
		Ttl:  cf.Ttl,
	}}
	return instance
}
