// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: api/mailer.proto

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
	Mailer_SendEmailConfirmation_FullMethodName = "/easypwn.Mailer/SendEmailConfirmation"
)

// MailerClient is the client API for Mailer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailerClient interface {
	SendEmailConfirmation(ctx context.Context, in *SendEmailConfirmationRequest, opts ...grpc.CallOption) (*SendEmailConfirmationResponse, error)
}

type mailerClient struct {
	cc grpc.ClientConnInterface
}

func NewMailerClient(cc grpc.ClientConnInterface) MailerClient {
	return &mailerClient{cc}
}

func (c *mailerClient) SendEmailConfirmation(ctx context.Context, in *SendEmailConfirmationRequest, opts ...grpc.CallOption) (*SendEmailConfirmationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendEmailConfirmationResponse)
	err := c.cc.Invoke(ctx, Mailer_SendEmailConfirmation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailerServer is the server API for Mailer service.
// All implementations must embed UnimplementedMailerServer
// for forward compatibility.
type MailerServer interface {
	SendEmailConfirmation(context.Context, *SendEmailConfirmationRequest) (*SendEmailConfirmationResponse, error)
	mustEmbedUnimplementedMailerServer()
}

// UnimplementedMailerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMailerServer struct{}

func (UnimplementedMailerServer) SendEmailConfirmation(context.Context, *SendEmailConfirmationRequest) (*SendEmailConfirmationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailConfirmation not implemented")
}
func (UnimplementedMailerServer) mustEmbedUnimplementedMailerServer() {}
func (UnimplementedMailerServer) testEmbeddedByValue()                {}

// UnsafeMailerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailerServer will
// result in compilation errors.
type UnsafeMailerServer interface {
	mustEmbedUnimplementedMailerServer()
}

func RegisterMailerServer(s grpc.ServiceRegistrar, srv MailerServer) {
	// If the following call pancis, it indicates UnimplementedMailerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Mailer_ServiceDesc, srv)
}

func _Mailer_SendEmailConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendEmailConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mailer_SendEmailConfirmation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendEmailConfirmation(ctx, req.(*SendEmailConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mailer_ServiceDesc is the grpc.ServiceDesc for Mailer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mailer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "easypwn.Mailer",
	HandlerType: (*MailerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmailConfirmation",
			Handler:    _Mailer_SendEmailConfirmation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mailer.proto",
}
