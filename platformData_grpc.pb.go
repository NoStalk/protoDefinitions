// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/platformData.proto

package platformDatapb

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

// FetchPlatformDataClient is the client API for FetchPlatformData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchPlatformDataClient interface {
	GetUserSubmissions(ctx context.Context, in *Request, opts ...grpc.CallOption) (*SubmissionResponse, error)
	GetUserContests(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ContestResponse, error)
	GetAllUserData(ctx context.Context, opts ...grpc.CallOption) (FetchPlatformData_GetAllUserDataClient, error)
}

type fetchPlatformDataClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchPlatformDataClient(cc grpc.ClientConnInterface) FetchPlatformDataClient {
	return &fetchPlatformDataClient{cc}
}

func (c *fetchPlatformDataClient) GetUserSubmissions(ctx context.Context, in *Request, opts ...grpc.CallOption) (*SubmissionResponse, error) {
	out := new(SubmissionResponse)
	err := c.cc.Invoke(ctx, "/proto.FetchPlatformData/getUserSubmissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchPlatformDataClient) GetUserContests(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ContestResponse, error) {
	out := new(ContestResponse)
	err := c.cc.Invoke(ctx, "/proto.FetchPlatformData/getUserContests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchPlatformDataClient) GetAllUserData(ctx context.Context, opts ...grpc.CallOption) (FetchPlatformData_GetAllUserDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &FetchPlatformData_ServiceDesc.Streams[0], "/proto.FetchPlatformData/getAllUserData", opts...)
	if err != nil {
		return nil, err
	}
	x := &fetchPlatformDataGetAllUserDataClient{stream}
	return x, nil
}

type FetchPlatformData_GetAllUserDataClient interface {
	Send(*Request) error
	Recv() (*OperationStatus, error)
	grpc.ClientStream
}

type fetchPlatformDataGetAllUserDataClient struct {
	grpc.ClientStream
}

func (x *fetchPlatformDataGetAllUserDataClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fetchPlatformDataGetAllUserDataClient) Recv() (*OperationStatus, error) {
	m := new(OperationStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FetchPlatformDataServer is the server API for FetchPlatformData service.
// All implementations must embed UnimplementedFetchPlatformDataServer
// for forward compatibility
type FetchPlatformDataServer interface {
	GetUserSubmissions(context.Context, *Request) (*SubmissionResponse, error)
	GetUserContests(context.Context, *Request) (*ContestResponse, error)
	GetAllUserData(FetchPlatformData_GetAllUserDataServer) error
	mustEmbedUnimplementedFetchPlatformDataServer()
}

// UnimplementedFetchPlatformDataServer must be embedded to have forward compatible implementations.
type UnimplementedFetchPlatformDataServer struct {
}

func (UnimplementedFetchPlatformDataServer) GetUserSubmissions(context.Context, *Request) (*SubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSubmissions not implemented")
}
func (UnimplementedFetchPlatformDataServer) GetUserContests(context.Context, *Request) (*ContestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserContests not implemented")
}
func (UnimplementedFetchPlatformDataServer) GetAllUserData(FetchPlatformData_GetAllUserDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllUserData not implemented")
}
func (UnimplementedFetchPlatformDataServer) mustEmbedUnimplementedFetchPlatformDataServer() {}

// UnsafeFetchPlatformDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchPlatformDataServer will
// result in compilation errors.
type UnsafeFetchPlatformDataServer interface {
	mustEmbedUnimplementedFetchPlatformDataServer()
}

func RegisterFetchPlatformDataServer(s grpc.ServiceRegistrar, srv FetchPlatformDataServer) {
	s.RegisterService(&FetchPlatformData_ServiceDesc, srv)
}

func _FetchPlatformData_GetUserSubmissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchPlatformDataServer).GetUserSubmissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FetchPlatformData/getUserSubmissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchPlatformDataServer).GetUserSubmissions(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FetchPlatformData_GetUserContests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchPlatformDataServer).GetUserContests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FetchPlatformData/getUserContests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchPlatformDataServer).GetUserContests(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FetchPlatformData_GetAllUserData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FetchPlatformDataServer).GetAllUserData(&fetchPlatformDataGetAllUserDataServer{stream})
}

type FetchPlatformData_GetAllUserDataServer interface {
	Send(*OperationStatus) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type fetchPlatformDataGetAllUserDataServer struct {
	grpc.ServerStream
}

func (x *fetchPlatformDataGetAllUserDataServer) Send(m *OperationStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fetchPlatformDataGetAllUserDataServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FetchPlatformData_ServiceDesc is the grpc.ServiceDesc for FetchPlatformData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FetchPlatformData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FetchPlatformData",
	HandlerType: (*FetchPlatformDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUserSubmissions",
			Handler:    _FetchPlatformData_GetUserSubmissions_Handler,
		},
		{
			MethodName: "getUserContests",
			Handler:    _FetchPlatformData_GetUserContests_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getAllUserData",
			Handler:       _FetchPlatformData_GetAllUserData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/platformData.proto",
}
