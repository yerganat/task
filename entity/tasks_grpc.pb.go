// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package entity

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

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TasksClient interface {
	Save(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskId, error)
	Check(ctx context.Context, in *TaskId, opts ...grpc.CallOption) (*Status, error)
}

type tasksClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksClient(cc grpc.ClientConnInterface) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) Save(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskId, error) {
	out := new(TaskId)
	err := c.cc.Invoke(ctx, "/entity.Tasks/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) Check(ctx context.Context, in *TaskId, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/entity.Tasks/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
// All implementations must embed UnimplementedTasksServer
// for forward compatibility
type TasksServer interface {
	Save(context.Context, *Task) (*TaskId, error)
	Check(context.Context, *TaskId) (*Status, error)
	mustEmbedUnimplementedTasksServer()
}

// UnimplementedTasksServer must be embedded to have forward compatible implementations.
type UnimplementedTasksServer struct {
}

func (UnimplementedTasksServer) Save(context.Context, *Task) (*TaskId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedTasksServer) Check(context.Context, *TaskId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedTasksServer) mustEmbedUnimplementedTasksServer() {}

// UnsafeTasksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksServer will
// result in compilation errors.
type UnsafeTasksServer interface {
	mustEmbedUnimplementedTasksServer()
}

func RegisterTasksServer(s grpc.ServiceRegistrar, srv TasksServer) {
	s.RegisterService(&Tasks_ServiceDesc, srv)
}

func _Tasks_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.Tasks/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).Save(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.Tasks/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).Check(ctx, req.(*TaskId))
	}
	return interceptor(ctx, in, info, handler)
}

// Tasks_ServiceDesc is the grpc.ServiceDesc for Tasks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tasks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entity.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Tasks_Save_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Tasks_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entity/tasks.proto",
}
