// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: movies_people_service_v1.proto

package protos

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

// MoviesPeopleServiceV1Client is the client API for MoviesPeopleServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoviesPeopleServiceV1Client interface {
	GetPeople(ctx context.Context, in *GetMoviePeopleRequest, opts ...grpc.CallOption) (*Humans, error)
}

type moviesPeopleServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewMoviesPeopleServiceV1Client(cc grpc.ClientConnInterface) MoviesPeopleServiceV1Client {
	return &moviesPeopleServiceV1Client{cc}
}

func (c *moviesPeopleServiceV1Client) GetPeople(ctx context.Context, in *GetMoviePeopleRequest, opts ...grpc.CallOption) (*Humans, error) {
	out := new(Humans)
	err := c.cc.Invoke(ctx, "/movies_people_service.moviesPeopleServiceV1/GetPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesPeopleServiceV1Server is the server API for MoviesPeopleServiceV1 service.
// All implementations must embed UnimplementedMoviesPeopleServiceV1Server
// for forward compatibility
type MoviesPeopleServiceV1Server interface {
	GetPeople(context.Context, *GetMoviePeopleRequest) (*Humans, error)
	mustEmbedUnimplementedMoviesPeopleServiceV1Server()
}

// UnimplementedMoviesPeopleServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedMoviesPeopleServiceV1Server struct {
}

func (UnimplementedMoviesPeopleServiceV1Server) GetPeople(context.Context, *GetMoviePeopleRequest) (*Humans, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeople not implemented")
}
func (UnimplementedMoviesPeopleServiceV1Server) mustEmbedUnimplementedMoviesPeopleServiceV1Server() {}

// UnsafeMoviesPeopleServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoviesPeopleServiceV1Server will
// result in compilation errors.
type UnsafeMoviesPeopleServiceV1Server interface {
	mustEmbedUnimplementedMoviesPeopleServiceV1Server()
}

func RegisterMoviesPeopleServiceV1Server(s grpc.ServiceRegistrar, srv MoviesPeopleServiceV1Server) {
	s.RegisterService(&MoviesPeopleServiceV1_ServiceDesc, srv)
}

func _MoviesPeopleServiceV1_GetPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoviePeopleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesPeopleServiceV1Server).GetPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movies_people_service.moviesPeopleServiceV1/GetPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesPeopleServiceV1Server).GetPeople(ctx, req.(*GetMoviePeopleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MoviesPeopleServiceV1_ServiceDesc is the grpc.ServiceDesc for MoviesPeopleServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoviesPeopleServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "movies_people_service.moviesPeopleServiceV1",
	HandlerType: (*MoviesPeopleServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeople",
			Handler:    _MoviesPeopleServiceV1_GetPeople_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movies_people_service_v1.proto",
}
