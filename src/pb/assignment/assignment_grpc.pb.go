// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: assignment.proto

package assignmentpb

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

// LogAssignmentServiceClient is the client API for LogAssignmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogAssignmentServiceClient interface {
	GetLogAssignmentByAssignment(ctx context.Context, in *GetLogAssignmentByAssignmentRequest, opts ...grpc.CallOption) (*GetLogAssignmentByAssignmentResponse, error)
	CreateLogAssignment(ctx context.Context, in *CreateLogAssignmentRequest, opts ...grpc.CallOption) (*CreateLogAssignmentResponse, error)
}

type logAssignmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogAssignmentServiceClient(cc grpc.ClientConnInterface) LogAssignmentServiceClient {
	return &logAssignmentServiceClient{cc}
}

func (c *logAssignmentServiceClient) GetLogAssignmentByAssignment(ctx context.Context, in *GetLogAssignmentByAssignmentRequest, opts ...grpc.CallOption) (*GetLogAssignmentByAssignmentResponse, error) {
	out := new(GetLogAssignmentByAssignmentResponse)
	err := c.cc.Invoke(ctx, "/assignment.LogAssignmentService/GetLogAssignmentByAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logAssignmentServiceClient) CreateLogAssignment(ctx context.Context, in *CreateLogAssignmentRequest, opts ...grpc.CallOption) (*CreateLogAssignmentResponse, error) {
	out := new(CreateLogAssignmentResponse)
	err := c.cc.Invoke(ctx, "/assignment.LogAssignmentService/CreateLogAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogAssignmentServiceServer is the server API for LogAssignmentService service.
// All implementations must embed UnimplementedLogAssignmentServiceServer
// for forward compatibility
type LogAssignmentServiceServer interface {
	GetLogAssignmentByAssignment(context.Context, *GetLogAssignmentByAssignmentRequest) (*GetLogAssignmentByAssignmentResponse, error)
	CreateLogAssignment(context.Context, *CreateLogAssignmentRequest) (*CreateLogAssignmentResponse, error)
	mustEmbedUnimplementedLogAssignmentServiceServer()
}

// UnimplementedLogAssignmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogAssignmentServiceServer struct {
}

func (UnimplementedLogAssignmentServiceServer) GetLogAssignmentByAssignment(context.Context, *GetLogAssignmentByAssignmentRequest) (*GetLogAssignmentByAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogAssignmentByAssignment not implemented")
}
func (UnimplementedLogAssignmentServiceServer) CreateLogAssignment(context.Context, *CreateLogAssignmentRequest) (*CreateLogAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLogAssignment not implemented")
}
func (UnimplementedLogAssignmentServiceServer) mustEmbedUnimplementedLogAssignmentServiceServer() {}

// UnsafeLogAssignmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogAssignmentServiceServer will
// result in compilation errors.
type UnsafeLogAssignmentServiceServer interface {
	mustEmbedUnimplementedLogAssignmentServiceServer()
}

func RegisterLogAssignmentServiceServer(s grpc.ServiceRegistrar, srv LogAssignmentServiceServer) {
	s.RegisterService(&LogAssignmentService_ServiceDesc, srv)
}

func _LogAssignmentService_GetLogAssignmentByAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogAssignmentByAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAssignmentServiceServer).GetLogAssignmentByAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assignment.LogAssignmentService/GetLogAssignmentByAssignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAssignmentServiceServer).GetLogAssignmentByAssignment(ctx, req.(*GetLogAssignmentByAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogAssignmentService_CreateLogAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLogAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAssignmentServiceServer).CreateLogAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assignment.LogAssignmentService/CreateLogAssignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAssignmentServiceServer).CreateLogAssignment(ctx, req.(*CreateLogAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogAssignmentService_ServiceDesc is the grpc.ServiceDesc for LogAssignmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogAssignmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assignment.LogAssignmentService",
	HandlerType: (*LogAssignmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLogAssignmentByAssignment",
			Handler:    _LogAssignmentService_GetLogAssignmentByAssignment_Handler,
		},
		{
			MethodName: "CreateLogAssignment",
			Handler:    _LogAssignmentService_CreateLogAssignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assignment.proto",
}
