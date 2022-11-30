package service

import (
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/global"
	"google.golang.org/grpc/resolver"
)

func GetServiceCommonClient() (proto_build.CommonServiceClient, error) {
	clientV3 := NewClientV3(global.GH.Conf.Etcd.Host, global.GH.Conf.Etcd.Timeout)
	builder := NewJResolverBuilder(clientV3)
	resolver.Register(builder)
	conn, err := GetRpcConn("service_common")
	if err != nil {
		return nil, err
	}
	client := proto_build.NewCommonServiceClient(conn)
	return client, nil
}

func GetServiceGroupClient() (proto_build.GroupServiceClient, error) {
	clientV3 := NewClientV3(global.GH.Conf.Etcd.Host, global.GH.Conf.Etcd.Timeout)
	builder := NewJResolverBuilder(clientV3)
	resolver.Register(builder)
	conn, err := GetRpcConn("service_group")
	if err != nil {
		return nil, err
	}
	client := proto_build.NewGroupServiceClient(conn)
	return client, nil
}

func GetServiceMessageClient() (proto_build.MessageServiceClient, error) {
	clientV3 := NewClientV3(global.GH.Conf.Etcd.Host, global.GH.Conf.Etcd.Timeout)
	builder := NewJResolverBuilder(clientV3)
	resolver.Register(builder)
	conn, err := GetRpcConn("service_message")
	if err != nil {
		return nil, err
	}
	client := proto_build.NewMessageServiceClient(conn)
	return client, nil
}

func GetServicePingClient() (proto_build.PingServiceClient, error) {
	clientV3 := NewClientV3(global.GH.Conf.Etcd.Host, global.GH.Conf.Etcd.Timeout)
	builder := NewJResolverBuilder(clientV3)
	resolver.Register(builder)
	conn, err := GetRpcConn("service_ping")
	if err != nil {
		return nil, err
	}
	client := proto_build.NewPingServiceClient(conn)
	return client, nil
}

func GetServiceUserClient() (proto_build.UserServiceClient, error) {
	clientV3 := NewClientV3(global.GH.Conf.Etcd.Host, global.GH.Conf.Etcd.Timeout)
	builder := NewJResolverBuilder(clientV3)
	resolver.Register(builder)
	conn, err := GetRpcConn("service_user")
	if err != nil {
		return nil, err
	}
	client := proto_build.NewUserServiceClient(conn)
	return client, nil
}
