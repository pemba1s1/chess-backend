// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: chess/chess.proto

package chess

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

// ChessClient is the client API for Chess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChessClient interface {
	NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error)
}

type chessClient struct {
	cc grpc.ClientConnInterface
}

func NewChessClient(cc grpc.ClientConnInterface) ChessClient {
	return &chessClient{cc}
}

func (c *chessClient) NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error) {
	out := new(NewGameResponse)
	err := c.cc.Invoke(ctx, "/Chess/NewGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChessServer is the server API for Chess service.
// All implementations must embed UnimplementedChessServer
// for forward compatibility
type ChessServer interface {
	NewGame(context.Context, *NewGameRequest) (*NewGameResponse, error)
	mustEmbedUnimplementedChessServer()
}

// UnimplementedChessServer must be embedded to have forward compatible implementations.
type UnimplementedChessServer struct {
}

func (UnimplementedChessServer) NewGame(context.Context, *NewGameRequest) (*NewGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewGame not implemented")
}
func (UnimplementedChessServer) mustEmbedUnimplementedChessServer() {}

// UnsafeChessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChessServer will
// result in compilation errors.
type UnsafeChessServer interface {
	mustEmbedUnimplementedChessServer()
}

func RegisterChessServer(s grpc.ServiceRegistrar, srv ChessServer) {
	s.RegisterService(&Chess_ServiceDesc, srv)
}

func _Chess_NewGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChessServer).NewGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chess/NewGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChessServer).NewGame(ctx, req.(*NewGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chess_ServiceDesc is the grpc.ServiceDesc for Chess service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chess_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Chess",
	HandlerType: (*ChessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGame",
			Handler:    _Chess_NewGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chess/chess.proto",
}
