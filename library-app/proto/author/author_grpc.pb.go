// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package libraryService

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

// WriterServiceClient is the client API for WriterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriterServiceClient interface {
	CreateAuthor(ctx context.Context, in *CreateAuthorReq, opts ...grpc.CallOption) (*CreateAuthorRes, error)
}

type writerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWriterServiceClient(cc grpc.ClientConnInterface) WriterServiceClient {
	return &writerServiceClient{cc}
}

func (c *writerServiceClient) CreateAuthor(ctx context.Context, in *CreateAuthorReq, opts ...grpc.CallOption) (*CreateAuthorRes, error) {
	out := new(CreateAuthorRes)
	err := c.cc.Invoke(ctx, "/libraryService.writerService/CreateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriterServiceServer is the server API for WriterService service.
// All implementations should embed UnimplementedWriterServiceServer
// for forward compatibility
type WriterServiceServer interface {
	CreateAuthor(context.Context, *CreateAuthorReq) (*CreateAuthorRes, error)
}

// UnimplementedWriterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWriterServiceServer struct {
}

func (UnimplementedWriterServiceServer) CreateAuthor(context.Context, *CreateAuthorReq) (*CreateAuthorRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthor not implemented")
}

// UnsafeWriterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriterServiceServer will
// result in compilation errors.
type UnsafeWriterServiceServer interface {
	mustEmbedUnimplementedWriterServiceServer()
}

func RegisterWriterServiceServer(s grpc.ServiceRegistrar, srv WriterServiceServer) {
	s.RegisterService(&WriterService_ServiceDesc, srv)
}

func _WriterService_CreateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServiceServer).CreateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libraryService.writerService/CreateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServiceServer).CreateAuthor(ctx, req.(*CreateAuthorReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WriterService_ServiceDesc is the grpc.ServiceDesc for WriterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libraryService.writerService",
	HandlerType: (*WriterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuthor",
			Handler:    _WriterService_CreateAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "author.proto",
}