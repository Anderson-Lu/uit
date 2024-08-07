// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: todo/toto.proto

package todo

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

// UitTodoClient is the client API for UitTodo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UitTodoClient interface {
	Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRsp, error)
	Remove(ctx context.Context, in *RemoveReq, opts ...grpc.CallOption) (*RemoveRsp, error)
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListRsp, error)
	Modify(ctx context.Context, in *ModifyReq, opts ...grpc.CallOption) (*ModifyRsp, error)
}

type uitTodoClient struct {
	cc grpc.ClientConnInterface
}

func NewUitTodoClient(cc grpc.ClientConnInterface) UitTodoClient {
	return &uitTodoClient{cc}
}

func (c *uitTodoClient) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRsp, error) {
	out := new(AddRsp)
	err := c.cc.Invoke(ctx, "/todo.UitTodo/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uitTodoClient) Remove(ctx context.Context, in *RemoveReq, opts ...grpc.CallOption) (*RemoveRsp, error) {
	out := new(RemoveRsp)
	err := c.cc.Invoke(ctx, "/todo.UitTodo/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uitTodoClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListRsp, error) {
	out := new(ListRsp)
	err := c.cc.Invoke(ctx, "/todo.UitTodo/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uitTodoClient) Modify(ctx context.Context, in *ModifyReq, opts ...grpc.CallOption) (*ModifyRsp, error) {
	out := new(ModifyRsp)
	err := c.cc.Invoke(ctx, "/todo.UitTodo/Modify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UitTodoServer is the server API for UitTodo service.
// All implementations must embed UnimplementedUitTodoServer
// for forward compatibility
type UitTodoServer interface {
	Add(context.Context, *AddReq) (*AddRsp, error)
	Remove(context.Context, *RemoveReq) (*RemoveRsp, error)
	List(context.Context, *ListReq) (*ListRsp, error)
	Modify(context.Context, *ModifyReq) (*ModifyRsp, error)
	mustEmbedUnimplementedUitTodoServer()
}

// UnimplementedUitTodoServer must be embedded to have forward compatible implementations.
type UnimplementedUitTodoServer struct {
}

func (UnimplementedUitTodoServer) Add(context.Context, *AddReq) (*AddRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedUitTodoServer) Remove(context.Context, *RemoveReq) (*RemoveRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedUitTodoServer) List(context.Context, *ListReq) (*ListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUitTodoServer) Modify(context.Context, *ModifyReq) (*ModifyRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modify not implemented")
}
func (UnimplementedUitTodoServer) mustEmbedUnimplementedUitTodoServer() {}

// UnsafeUitTodoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UitTodoServer will
// result in compilation errors.
type UnsafeUitTodoServer interface {
	mustEmbedUnimplementedUitTodoServer()
}

func RegisterUitTodoServer(s grpc.ServiceRegistrar, srv UitTodoServer) {
	s.RegisterService(&UitTodo_ServiceDesc, srv)
}

func _UitTodo_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UitTodoServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.UitTodo/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UitTodoServer).Add(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UitTodo_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UitTodoServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.UitTodo/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UitTodoServer).Remove(ctx, req.(*RemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UitTodo_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UitTodoServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.UitTodo/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UitTodoServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UitTodo_Modify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UitTodoServer).Modify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.UitTodo/Modify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UitTodoServer).Modify(ctx, req.(*ModifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UitTodo_ServiceDesc is the grpc.ServiceDesc for UitTodo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UitTodo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.UitTodo",
	HandlerType: (*UitTodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _UitTodo_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _UitTodo_Remove_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UitTodo_List_Handler,
		},
		{
			MethodName: "Modify",
			Handler:    _UitTodo_Modify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo/toto.proto",
}
