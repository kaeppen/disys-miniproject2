// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package disysminiproject2

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

// ChittyChatClient is the client API for ChittyChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatClient interface {
	Publish(ctx context.Context, in *MessageWithLamport, opts ...grpc.CallOption) (*Empty, error)
	Broadcast(ctx context.Context, in *MessageWithLamport, opts ...grpc.CallOption) (*Empty, error)
	EstablishConnection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (ChittyChat_EstablishConnectionClient, error)
	RecieveBroadcast(ctx context.Context, in *MessageWithLamport, opts ...grpc.CallOption) (*Empty, error)
	Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*Empty, error)
}

type chittyChatClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatClient(cc grpc.ClientConnInterface) ChittyChatClient {
	return &chittyChatClient{cc}
}

func (c *chittyChatClient) Publish(ctx context.Context, in *MessageWithLamport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chittychat.ChittyChat/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatClient) Broadcast(ctx context.Context, in *MessageWithLamport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chittychat.ChittyChat/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatClient) EstablishConnection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (ChittyChat_EstablishConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChat_ServiceDesc.Streams[0], "/chittychat.ChittyChat/EstablishConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatEstablishConnectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChittyChat_EstablishConnectionClient interface {
	Recv() (*MessageWithLamport, error)
	grpc.ClientStream
}

type chittyChatEstablishConnectionClient struct {
	grpc.ClientStream
}

func (x *chittyChatEstablishConnectionClient) Recv() (*MessageWithLamport, error) {
	m := new(MessageWithLamport)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chittyChatClient) RecieveBroadcast(ctx context.Context, in *MessageWithLamport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chittychat.ChittyChat/RecieveBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatClient) Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chittychat.ChittyChat/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittyChatServer is the server API for ChittyChat service.
// All implementations must embed UnimplementedChittyChatServer
// for forward compatibility
type ChittyChatServer interface {
	Publish(context.Context, *MessageWithLamport) (*Empty, error)
	Broadcast(context.Context, *MessageWithLamport) (*Empty, error)
	EstablishConnection(*ConnectionRequest, ChittyChat_EstablishConnectionServer) error
	RecieveBroadcast(context.Context, *MessageWithLamport) (*Empty, error)
	Leave(context.Context, *LeaveRequest) (*Empty, error)
	mustEmbedUnimplementedChittyChatServer()
}

// UnimplementedChittyChatServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServer struct {
}

func (UnimplementedChittyChatServer) Publish(context.Context, *MessageWithLamport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedChittyChatServer) Broadcast(context.Context, *MessageWithLamport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedChittyChatServer) EstablishConnection(*ConnectionRequest, ChittyChat_EstablishConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method EstablishConnection not implemented")
}
func (UnimplementedChittyChatServer) RecieveBroadcast(context.Context, *MessageWithLamport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecieveBroadcast not implemented")
}
func (UnimplementedChittyChatServer) Leave(context.Context, *LeaveRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (UnimplementedChittyChatServer) mustEmbedUnimplementedChittyChatServer() {}

// UnsafeChittyChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServer will
// result in compilation errors.
type UnsafeChittyChatServer interface {
	mustEmbedUnimplementedChittyChatServer()
}

func RegisterChittyChatServer(s grpc.ServiceRegistrar, srv ChittyChatServer) {
	s.RegisterService(&ChittyChat_ServiceDesc, srv)
}

func _ChittyChat_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageWithLamport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chittychat.ChittyChat/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).Publish(ctx, req.(*MessageWithLamport))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChat_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageWithLamport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chittychat.ChittyChat/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).Broadcast(ctx, req.(*MessageWithLamport))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChat_EstablishConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChittyChatServer).EstablishConnection(m, &chittyChatEstablishConnectionServer{stream})
}

type ChittyChat_EstablishConnectionServer interface {
	Send(*MessageWithLamport) error
	grpc.ServerStream
}

type chittyChatEstablishConnectionServer struct {
	grpc.ServerStream
}

func (x *chittyChatEstablishConnectionServer) Send(m *MessageWithLamport) error {
	return x.ServerStream.SendMsg(m)
}

func _ChittyChat_RecieveBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageWithLamport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).RecieveBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chittychat.ChittyChat/RecieveBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).RecieveBroadcast(ctx, req.(*MessageWithLamport))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChat_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chittychat.ChittyChat/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).Leave(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChittyChat_ServiceDesc is the grpc.ServiceDesc for ChittyChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chittychat.ChittyChat",
	HandlerType: (*ChittyChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _ChittyChat_Publish_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _ChittyChat_Broadcast_Handler,
		},
		{
			MethodName: "RecieveBroadcast",
			Handler:    _ChittyChat_RecieveBroadcast_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _ChittyChat_Leave_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EstablishConnection",
			Handler:       _ChittyChat_EstablishConnection_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chittychat.proto",
}
