// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: api/board_service.proto

package protoc

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

// BoardServiceClient is the client API for BoardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardServiceClient interface {
	CreateBoard(ctx context.Context, in *CreateBoardRequest, opts ...grpc.CallOption) (*CreateBoardResponse, error)
}

type boardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardServiceClient(cc grpc.ClientConnInterface) BoardServiceClient {
	return &boardServiceClient{cc}
}

func (c *boardServiceClient) CreateBoard(ctx context.Context, in *CreateBoardRequest, opts ...grpc.CallOption) (*CreateBoardResponse, error) {
	out := new(CreateBoardResponse)
	err := c.cc.Invoke(ctx, "/BoardService/CreateBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServiceServer is the server API for BoardService service.
// All implementations must embed UnimplementedBoardServiceServer
// for forward compatibility
type BoardServiceServer interface {
	CreateBoard(context.Context, *CreateBoardRequest) (*CreateBoardResponse, error)
	mustEmbedUnimplementedBoardServiceServer()
}

// UnimplementedBoardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBoardServiceServer struct {
}

func (UnimplementedBoardServiceServer) CreateBoard(context.Context, *CreateBoardRequest) (*CreateBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBoard not implemented")
}
func (UnimplementedBoardServiceServer) mustEmbedUnimplementedBoardServiceServer() {}

// UnsafeBoardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServiceServer will
// result in compilation errors.
type UnsafeBoardServiceServer interface {
	mustEmbedUnimplementedBoardServiceServer()
}

func RegisterBoardServiceServer(s grpc.ServiceRegistrar, srv BoardServiceServer) {
	s.RegisterService(&BoardService_ServiceDesc, srv)
}

func _BoardService_CreateBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).CreateBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BoardService/CreateBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).CreateBoard(ctx, req.(*CreateBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoardService_ServiceDesc is the grpc.ServiceDesc for BoardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BoardService",
	HandlerType: (*BoardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBoard",
			Handler:    _BoardService_CreateBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/board_service.proto",
}