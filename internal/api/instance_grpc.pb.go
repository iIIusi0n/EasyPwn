// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: api/instance.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Instance_CreateInstance_FullMethodName  = "/easypwn.Instance/CreateInstance"
	Instance_GetInstance_FullMethodName     = "/easypwn.Instance/GetInstance"
	Instance_DeleteInstance_FullMethodName  = "/easypwn.Instance/DeleteInstance"
	Instance_GetInstanceLogs_FullMethodName = "/easypwn.Instance/GetInstanceLogs"
)

// InstanceClient is the client API for Instance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstanceClient interface {
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error)
	GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error)
	DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceResponse, error)
	GetInstanceLogs(ctx context.Context, in *GetInstanceLogsRequest, opts ...grpc.CallOption) (*GetInstanceLogsResponse, error)
}

type instanceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstanceClient(cc grpc.ClientConnInterface) InstanceClient {
	return &instanceClient{cc}
}

func (c *instanceClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInstanceResponse)
	err := c.cc.Invoke(ctx, Instance_CreateInstance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInstanceResponse)
	err := c.cc.Invoke(ctx, Instance_GetInstance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteInstanceResponse)
	err := c.cc.Invoke(ctx, Instance_DeleteInstance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) GetInstanceLogs(ctx context.Context, in *GetInstanceLogsRequest, opts ...grpc.CallOption) (*GetInstanceLogsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInstanceLogsResponse)
	err := c.cc.Invoke(ctx, Instance_GetInstanceLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceServer is the server API for Instance service.
// All implementations must embed UnimplementedInstanceServer
// for forward compatibility.
type InstanceServer interface {
	CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)
	GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error)
	DeleteInstance(context.Context, *DeleteInstanceRequest) (*DeleteInstanceResponse, error)
	GetInstanceLogs(context.Context, *GetInstanceLogsRequest) (*GetInstanceLogsResponse, error)
	mustEmbedUnimplementedInstanceServer()
}

// UnimplementedInstanceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInstanceServer struct{}

func (UnimplementedInstanceServer) CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstance not implemented")
}
func (UnimplementedInstanceServer) GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstance not implemented")
}
func (UnimplementedInstanceServer) DeleteInstance(context.Context, *DeleteInstanceRequest) (*DeleteInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstance not implemented")
}
func (UnimplementedInstanceServer) GetInstanceLogs(context.Context, *GetInstanceLogsRequest) (*GetInstanceLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstanceLogs not implemented")
}
func (UnimplementedInstanceServer) mustEmbedUnimplementedInstanceServer() {}
func (UnimplementedInstanceServer) testEmbeddedByValue()                  {}

// UnsafeInstanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstanceServer will
// result in compilation errors.
type UnsafeInstanceServer interface {
	mustEmbedUnimplementedInstanceServer()
}

func RegisterInstanceServer(s grpc.ServiceRegistrar, srv InstanceServer) {
	// If the following call pancis, it indicates UnimplementedInstanceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Instance_ServiceDesc, srv)
}

func _Instance_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Instance_CreateInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_GetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).GetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Instance_GetInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).GetInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_DeleteInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).DeleteInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Instance_DeleteInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).DeleteInstance(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_GetInstanceLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).GetInstanceLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Instance_GetInstanceLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).GetInstanceLogs(ctx, req.(*GetInstanceLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Instance_ServiceDesc is the grpc.ServiceDesc for Instance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Instance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "easypwn.Instance",
	HandlerType: (*InstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInstance",
			Handler:    _Instance_CreateInstance_Handler,
		},
		{
			MethodName: "GetInstance",
			Handler:    _Instance_GetInstance_Handler,
		},
		{
			MethodName: "DeleteInstance",
			Handler:    _Instance_DeleteInstance_Handler,
		},
		{
			MethodName: "GetInstanceLogs",
			Handler:    _Instance_GetInstanceLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/instance.proto",
}