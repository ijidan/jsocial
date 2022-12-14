// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: message.proto

package proto_build

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	SendUserTextMessage(ctx context.Context, in *SendUserTextMessageRequest, opts ...grpc.CallOption) (*SendUserTextMessageResponse, error)
	SendUserLocationMessage(ctx context.Context, in *SendUserLocationMessageRequest, opts ...grpc.CallOption) (*SendUserLocationMessageResponse, error)
	SendUserFaceMessage(ctx context.Context, in *SendUserFaceMessageRequest, opts ...grpc.CallOption) (*SendUserFaceMessageResponse, error)
	SendUserSoundMessage(ctx context.Context, in *SendUserSoundMessageRequest, opts ...grpc.CallOption) (*SendUserSoundMessageResponse, error)
	SendUserVideoMessage(ctx context.Context, in *SendUserVideoMessageRequest, opts ...grpc.CallOption) (*SendUserVideoMessageResponse, error)
	SendUserImageMessage(ctx context.Context, in *SendUserImageMessageRequest, opts ...grpc.CallOption) (*SendUserImageMessageResponse, error)
	SendUserFileMessage(ctx context.Context, in *SendUserFileMessageRequest, opts ...grpc.CallOption) (*SendUserFileMessageResponse, error)
	SendGroupTextMessage(ctx context.Context, in *SendGroupTextMessageRequest, opts ...grpc.CallOption) (*SendGroupTextMessageResponse, error)
	SendGroupLocationMessage(ctx context.Context, in *SendGroupLocationMessageRequest, opts ...grpc.CallOption) (*SendGroupLocationMessageResponse, error)
	SendGroupFceMessage(ctx context.Context, in *SendGroupFaceMessageRequest, opts ...grpc.CallOption) (*SendGroupFaceMessageResponse, error)
	SendGroupSoundMessage(ctx context.Context, in *SendGroupSoundMessageRequest, opts ...grpc.CallOption) (*SendGroupSoundMessageResponse, error)
	SendGroupVideoMessage(ctx context.Context, in *SendGroupVideoMessageRequest, opts ...grpc.CallOption) (*SendGroupVideoMessageResponse, error)
	SendGroupImageMessage(ctx context.Context, in *SendGroupImageMessageRequest, opts ...grpc.CallOption) (*SendGroupImageMessageResponse, error)
	SendGroupFileMessage(ctx context.Context, in *SendGroupFileMessageRequest, opts ...grpc.CallOption) (*SendGroupFileMessageResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) SendUserTextMessage(ctx context.Context, in *SendUserTextMessageRequest, opts ...grpc.CallOption) (*SendUserTextMessageResponse, error) {
	out := new(SendUserTextMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendUserTextMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendUserLocationMessage(ctx context.Context, in *SendUserLocationMessageRequest, opts ...grpc.CallOption) (*SendUserLocationMessageResponse, error) {
	out := new(SendUserLocationMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendUserLocationMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendUserFaceMessage(ctx context.Context, in *SendUserFaceMessageRequest, opts ...grpc.CallOption) (*SendUserFaceMessageResponse, error) {
	out := new(SendUserFaceMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendUserFaceMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendUserSoundMessage(ctx context.Context, in *SendUserSoundMessageRequest, opts ...grpc.CallOption) (*SendUserSoundMessageResponse, error) {
	out := new(SendUserSoundMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendUserSoundMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendUserVideoMessage(ctx context.Context, in *SendUserVideoMessageRequest, opts ...grpc.CallOption) (*SendUserVideoMessageResponse, error) {
	out := new(SendUserVideoMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendUserVideoMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendUserImageMessage(ctx context.Context, in *SendUserImageMessageRequest, opts ...grpc.CallOption) (*SendUserImageMessageResponse, error) {
	out := new(SendUserImageMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendUserImageMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendUserFileMessage(ctx context.Context, in *SendUserFileMessageRequest, opts ...grpc.CallOption) (*SendUserFileMessageResponse, error) {
	out := new(SendUserFileMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendUserFileMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendGroupTextMessage(ctx context.Context, in *SendGroupTextMessageRequest, opts ...grpc.CallOption) (*SendGroupTextMessageResponse, error) {
	out := new(SendGroupTextMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendGroupTextMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendGroupLocationMessage(ctx context.Context, in *SendGroupLocationMessageRequest, opts ...grpc.CallOption) (*SendGroupLocationMessageResponse, error) {
	out := new(SendGroupLocationMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendGroupLocationMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendGroupFceMessage(ctx context.Context, in *SendGroupFaceMessageRequest, opts ...grpc.CallOption) (*SendGroupFaceMessageResponse, error) {
	out := new(SendGroupFaceMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendGroupFceMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendGroupSoundMessage(ctx context.Context, in *SendGroupSoundMessageRequest, opts ...grpc.CallOption) (*SendGroupSoundMessageResponse, error) {
	out := new(SendGroupSoundMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendGroupSoundMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendGroupVideoMessage(ctx context.Context, in *SendGroupVideoMessageRequest, opts ...grpc.CallOption) (*SendGroupVideoMessageResponse, error) {
	out := new(SendGroupVideoMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendGroupVideoMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendGroupImageMessage(ctx context.Context, in *SendGroupImageMessageRequest, opts ...grpc.CallOption) (*SendGroupImageMessageResponse, error) {
	out := new(SendGroupImageMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendGroupImageMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendGroupFileMessage(ctx context.Context, in *SendGroupFileMessageRequest, opts ...grpc.CallOption) (*SendGroupFileMessageResponse, error) {
	out := new(SendGroupFileMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendGroupFileMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	SendUserTextMessage(context.Context, *SendUserTextMessageRequest) (*SendUserTextMessageResponse, error)
	SendUserLocationMessage(context.Context, *SendUserLocationMessageRequest) (*SendUserLocationMessageResponse, error)
	SendUserFaceMessage(context.Context, *SendUserFaceMessageRequest) (*SendUserFaceMessageResponse, error)
	SendUserSoundMessage(context.Context, *SendUserSoundMessageRequest) (*SendUserSoundMessageResponse, error)
	SendUserVideoMessage(context.Context, *SendUserVideoMessageRequest) (*SendUserVideoMessageResponse, error)
	SendUserImageMessage(context.Context, *SendUserImageMessageRequest) (*SendUserImageMessageResponse, error)
	SendUserFileMessage(context.Context, *SendUserFileMessageRequest) (*SendUserFileMessageResponse, error)
	SendGroupTextMessage(context.Context, *SendGroupTextMessageRequest) (*SendGroupTextMessageResponse, error)
	SendGroupLocationMessage(context.Context, *SendGroupLocationMessageRequest) (*SendGroupLocationMessageResponse, error)
	SendGroupFceMessage(context.Context, *SendGroupFaceMessageRequest) (*SendGroupFaceMessageResponse, error)
	SendGroupSoundMessage(context.Context, *SendGroupSoundMessageRequest) (*SendGroupSoundMessageResponse, error)
	SendGroupVideoMessage(context.Context, *SendGroupVideoMessageRequest) (*SendGroupVideoMessageResponse, error)
	SendGroupImageMessage(context.Context, *SendGroupImageMessageRequest) (*SendGroupImageMessageResponse, error)
	SendGroupFileMessage(context.Context, *SendGroupFileMessageRequest) (*SendGroupFileMessageResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) SendUserTextMessage(context.Context, *SendUserTextMessageRequest) (*SendUserTextMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserTextMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendUserLocationMessage(context.Context, *SendUserLocationMessageRequest) (*SendUserLocationMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserLocationMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendUserFaceMessage(context.Context, *SendUserFaceMessageRequest) (*SendUserFaceMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserFaceMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendUserSoundMessage(context.Context, *SendUserSoundMessageRequest) (*SendUserSoundMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserSoundMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendUserVideoMessage(context.Context, *SendUserVideoMessageRequest) (*SendUserVideoMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserVideoMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendUserImageMessage(context.Context, *SendUserImageMessageRequest) (*SendUserImageMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserImageMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendUserFileMessage(context.Context, *SendUserFileMessageRequest) (*SendUserFileMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserFileMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendGroupTextMessage(context.Context, *SendGroupTextMessageRequest) (*SendGroupTextMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupTextMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendGroupLocationMessage(context.Context, *SendGroupLocationMessageRequest) (*SendGroupLocationMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupLocationMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendGroupFceMessage(context.Context, *SendGroupFaceMessageRequest) (*SendGroupFaceMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupFceMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendGroupSoundMessage(context.Context, *SendGroupSoundMessageRequest) (*SendGroupSoundMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupSoundMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendGroupVideoMessage(context.Context, *SendGroupVideoMessageRequest) (*SendGroupVideoMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupVideoMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendGroupImageMessage(context.Context, *SendGroupImageMessageRequest) (*SendGroupImageMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupImageMessage not implemented")
}
func (UnimplementedMessageServiceServer) SendGroupFileMessage(context.Context, *SendGroupFileMessageRequest) (*SendGroupFileMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupFileMessage not implemented")
}
func (UnimplementedMessageServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_SendUserTextMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserTextMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendUserTextMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendUserTextMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendUserTextMessage(ctx, req.(*SendUserTextMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendUserLocationMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserLocationMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendUserLocationMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendUserLocationMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendUserLocationMessage(ctx, req.(*SendUserLocationMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendUserFaceMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserFaceMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendUserFaceMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendUserFaceMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendUserFaceMessage(ctx, req.(*SendUserFaceMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendUserSoundMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserSoundMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendUserSoundMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendUserSoundMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendUserSoundMessage(ctx, req.(*SendUserSoundMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendUserVideoMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserVideoMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendUserVideoMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendUserVideoMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendUserVideoMessage(ctx, req.(*SendUserVideoMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendUserImageMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserImageMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendUserImageMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendUserImageMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendUserImageMessage(ctx, req.(*SendUserImageMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendUserFileMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserFileMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendUserFileMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendUserFileMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendUserFileMessage(ctx, req.(*SendUserFileMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendGroupTextMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupTextMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendGroupTextMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendGroupTextMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendGroupTextMessage(ctx, req.(*SendGroupTextMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendGroupLocationMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupLocationMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendGroupLocationMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendGroupLocationMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendGroupLocationMessage(ctx, req.(*SendGroupLocationMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendGroupFceMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupFaceMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendGroupFceMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendGroupFceMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendGroupFceMessage(ctx, req.(*SendGroupFaceMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendGroupSoundMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupSoundMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendGroupSoundMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendGroupSoundMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendGroupSoundMessage(ctx, req.(*SendGroupSoundMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendGroupVideoMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupVideoMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendGroupVideoMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendGroupVideoMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendGroupVideoMessage(ctx, req.(*SendGroupVideoMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendGroupImageMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupImageMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendGroupImageMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendGroupImageMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendGroupImageMessage(ctx, req.(*SendGroupImageMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendGroupFileMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupFileMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendGroupFileMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendGroupFileMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendGroupFileMessage(ctx, req.(*SendGroupFileMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUserTextMessage",
			Handler:    _MessageService_SendUserTextMessage_Handler,
		},
		{
			MethodName: "SendUserLocationMessage",
			Handler:    _MessageService_SendUserLocationMessage_Handler,
		},
		{
			MethodName: "SendUserFaceMessage",
			Handler:    _MessageService_SendUserFaceMessage_Handler,
		},
		{
			MethodName: "SendUserSoundMessage",
			Handler:    _MessageService_SendUserSoundMessage_Handler,
		},
		{
			MethodName: "SendUserVideoMessage",
			Handler:    _MessageService_SendUserVideoMessage_Handler,
		},
		{
			MethodName: "SendUserImageMessage",
			Handler:    _MessageService_SendUserImageMessage_Handler,
		},
		{
			MethodName: "SendUserFileMessage",
			Handler:    _MessageService_SendUserFileMessage_Handler,
		},
		{
			MethodName: "SendGroupTextMessage",
			Handler:    _MessageService_SendGroupTextMessage_Handler,
		},
		{
			MethodName: "SendGroupLocationMessage",
			Handler:    _MessageService_SendGroupLocationMessage_Handler,
		},
		{
			MethodName: "SendGroupFceMessage",
			Handler:    _MessageService_SendGroupFceMessage_Handler,
		},
		{
			MethodName: "SendGroupSoundMessage",
			Handler:    _MessageService_SendGroupSoundMessage_Handler,
		},
		{
			MethodName: "SendGroupVideoMessage",
			Handler:    _MessageService_SendGroupVideoMessage_Handler,
		},
		{
			MethodName: "SendGroupImageMessage",
			Handler:    _MessageService_SendGroupImageMessage_Handler,
		},
		{
			MethodName: "SendGroupFileMessage",
			Handler:    _MessageService_SendGroupFileMessage_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _MessageService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
