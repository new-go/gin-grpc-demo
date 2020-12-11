// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BServiceClient is the client API for BService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BServiceClient interface {
	EchoB(ctx context.Context, in *BMsg, opts ...grpc.CallOption) (*BMsg, error)
}

type bServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBServiceClient(cc grpc.ClientConnInterface) BServiceClient {
	return &bServiceClient{cc}
}

func (c *bServiceClient) EchoB(ctx context.Context, in *BMsg, opts ...grpc.CallOption) (*BMsg, error) {
	out := new(BMsg)
	err := c.cc.Invoke(ctx, "/pb.BService/EchoB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BServiceServer is the server API for BService service.
// All implementations must embed UnimplementedBServiceServer
// for forward compatibility
type BServiceServer interface {
	EchoB(context.Context, *BMsg) (*BMsg, error)
	mustEmbedUnimplementedBServiceServer()
}

// UnimplementedBServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBServiceServer struct {
}

func (UnimplementedBServiceServer) EchoB(context.Context, *BMsg) (*BMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoB not implemented")
}
func (UnimplementedBServiceServer) mustEmbedUnimplementedBServiceServer() {}

// UnsafeBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BServiceServer will
// result in compilation errors.
type UnsafeBServiceServer interface {
	mustEmbedUnimplementedBServiceServer()
}

func RegisterBServiceServer(s grpc.ServiceRegistrar, srv BServiceServer) {
	s.RegisterService(&_BService_serviceDesc, srv)
}

func _BService_EchoB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BServiceServer).EchoB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BService/EchoB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BServiceServer).EchoB(ctx, req.(*BMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _BService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BService",
	HandlerType: (*BServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoB",
			Handler:    _BService_EchoB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "b.proto",
}