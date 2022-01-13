// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworldpb

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a pure grpc greeting.
	AAAASayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	// Sends a greeting.
	SayHelloREST(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	// Sends a greeting.
	SayHiREST(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) AAAASayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/AAAASayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloREST(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHelloREST", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHiREST(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHiREST", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a pure grpc greeting.
	AAAASayHello(context.Context, *Request) (*Reply, error)
	// Sends a greeting.
	SayHelloREST(context.Context, *Request) (*Reply, error)
	// Sends a greeting.
	SayHiREST(context.Context, *Request) (*Reply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) AAAASayHello(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AAAASayHello not implemented")
}
func (UnimplementedGreeterServer) SayHelloREST(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloREST not implemented")
}
func (UnimplementedGreeterServer) SayHiREST(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHiREST not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_AAAASayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).AAAASayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/AAAASayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).AAAASayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloREST_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloREST(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHelloREST",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloREST(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHiREST_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHiREST(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHiREST",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHiREST(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AAAASayHello",
			Handler:    _Greeter_AAAASayHello_Handler,
		},
		{
			MethodName: "SayHelloREST",
			Handler:    _Greeter_SayHelloREST_Handler,
		},
		{
			MethodName: "SayHiREST",
			Handler:    _Greeter_SayHiREST_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworldservice.proto",
}
