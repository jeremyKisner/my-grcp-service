// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: roller/rollerService/service.proto

package rollerService

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

const (
	Roller_Roll_FullMethodName = "/rollerService.Roller/Roll"
)

// RollerClient is the client API for Roller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RollerClient interface {
	Roll(ctx context.Context, in *RollerRequest, opts ...grpc.CallOption) (*RollerReply, error)
}

type rollerClient struct {
	cc grpc.ClientConnInterface
}

func NewRollerClient(cc grpc.ClientConnInterface) RollerClient {
	return &rollerClient{cc}
}

func (c *rollerClient) Roll(ctx context.Context, in *RollerRequest, opts ...grpc.CallOption) (*RollerReply, error) {
	out := new(RollerReply)
	err := c.cc.Invoke(ctx, Roller_Roll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RollerServer is the server API for Roller service.
// All implementations must embed UnimplementedRollerServer
// for forward compatibility
type RollerServer interface {
	Roll(context.Context, *RollerRequest) (*RollerReply, error)
	mustEmbedUnimplementedRollerServer()
}

// UnimplementedRollerServer must be embedded to have forward compatible implementations.
type UnimplementedRollerServer struct {
}

func (UnimplementedRollerServer) Roll(context.Context, *RollerRequest) (*RollerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Roll not implemented")
}
func (UnimplementedRollerServer) mustEmbedUnimplementedRollerServer() {}

// UnsafeRollerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RollerServer will
// result in compilation errors.
type UnsafeRollerServer interface {
	mustEmbedUnimplementedRollerServer()
}

func RegisterRollerServer(s grpc.ServiceRegistrar, srv RollerServer) {
	s.RegisterService(&Roller_ServiceDesc, srv)
}

func _Roller_Roll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RollerServer).Roll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roller_Roll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RollerServer).Roll(ctx, req.(*RollerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Roller_ServiceDesc is the grpc.ServiceDesc for Roller service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Roller_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rollerService.Roller",
	HandlerType: (*RollerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Roll",
			Handler:    _Roller_Roll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "roller/rollerService/service.proto",
}
