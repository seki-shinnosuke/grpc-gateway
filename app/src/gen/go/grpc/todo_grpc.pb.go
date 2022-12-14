// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: todo.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TodoApiClient is the client API for TodoApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoApiClient interface {
	// Get Tasks
	GetTasks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TodoApi_GetTasksClient, error)
	// Get Task By Task ID
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*TodoDetailResponse, error)
	// Create Task
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TodoDetailResponse, error)
	// Update Task By Task ID
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TodoDetailResponse, error)
	// Delete Task By Task ID
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type todoApiClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoApiClient(cc grpc.ClientConnInterface) TodoApiClient {
	return &todoApiClient{cc}
}

func (c *todoApiClient) GetTasks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TodoApi_GetTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoApi_ServiceDesc.Streams[0], "/grpc.TodoApi/GetTasks", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoApiGetTasksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoApi_GetTasksClient interface {
	Recv() (*TodoResponse, error)
	grpc.ClientStream
}

type todoApiGetTasksClient struct {
	grpc.ClientStream
}

func (x *todoApiGetTasksClient) Recv() (*TodoResponse, error) {
	m := new(TodoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoApiClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*TodoDetailResponse, error) {
	out := new(TodoDetailResponse)
	err := c.cc.Invoke(ctx, "/grpc.TodoApi/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoApiClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TodoDetailResponse, error) {
	out := new(TodoDetailResponse)
	err := c.cc.Invoke(ctx, "/grpc.TodoApi/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoApiClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TodoDetailResponse, error) {
	out := new(TodoDetailResponse)
	err := c.cc.Invoke(ctx, "/grpc.TodoApi/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoApiClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/grpc.TodoApi/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoApiServer is the server API for TodoApi service.
// All implementations must embed UnimplementedTodoApiServer
// for forward compatibility
type TodoApiServer interface {
	// Get Tasks
	GetTasks(*emptypb.Empty, TodoApi_GetTasksServer) error
	// Get Task By Task ID
	GetTask(context.Context, *GetTaskRequest) (*TodoDetailResponse, error)
	// Create Task
	CreateTask(context.Context, *CreateTaskRequest) (*TodoDetailResponse, error)
	// Update Task By Task ID
	UpdateTask(context.Context, *UpdateTaskRequest) (*TodoDetailResponse, error)
	// Delete Task By Task ID
	DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTodoApiServer()
}

// UnimplementedTodoApiServer must be embedded to have forward compatible implementations.
type UnimplementedTodoApiServer struct {
}

func (UnimplementedTodoApiServer) GetTasks(*emptypb.Empty, TodoApi_GetTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedTodoApiServer) GetTask(context.Context, *GetTaskRequest) (*TodoDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTodoApiServer) CreateTask(context.Context, *CreateTaskRequest) (*TodoDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTodoApiServer) UpdateTask(context.Context, *UpdateTaskRequest) (*TodoDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTodoApiServer) DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedTodoApiServer) mustEmbedUnimplementedTodoApiServer() {}

// UnsafeTodoApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoApiServer will
// result in compilation errors.
type UnsafeTodoApiServer interface {
	mustEmbedUnimplementedTodoApiServer()
}

func RegisterTodoApiServer(s grpc.ServiceRegistrar, srv TodoApiServer) {
	s.RegisterService(&TodoApi_ServiceDesc, srv)
}

func _TodoApi_GetTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoApiServer).GetTasks(m, &todoApiGetTasksServer{stream})
}

type TodoApi_GetTasksServer interface {
	Send(*TodoResponse) error
	grpc.ServerStream
}

type todoApiGetTasksServer struct {
	grpc.ServerStream
}

func (x *todoApiGetTasksServer) Send(m *TodoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TodoApi_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoApiServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TodoApi/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoApiServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoApi_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoApiServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TodoApi/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoApiServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoApi_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoApiServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TodoApi/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoApiServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoApi_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoApiServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TodoApi/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoApiServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoApi_ServiceDesc is the grpc.ServiceDesc for TodoApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.TodoApi",
	HandlerType: (*TodoApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _TodoApi_GetTask_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _TodoApi_CreateTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TodoApi_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TodoApi_DeleteTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTasks",
			Handler:       _TodoApi_GetTasks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo.proto",
}
