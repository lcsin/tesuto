// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: projectpb/module.proto

package projectpb

import (
	context "context"
	commonpb "github.com/lcsin/tesuto/proto/v1/rpc/commonpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ModuleService_GetModuleByID_FullMethodName         = "/project.ModuleService/GetModuleByID"
	ModuleService_GetModulesByProjectID_FullMethodName = "/project.ModuleService/GetModulesByProjectID"
	ModuleService_CreateModule_FullMethodName          = "/project.ModuleService/CreateModule"
	ModuleService_UpdateModuleByID_FullMethodName      = "/project.ModuleService/UpdateModuleByID"
	ModuleService_DeleteModuleByID_FullMethodName      = "/project.ModuleService/DeleteModuleByID"
)

// ModuleServiceClient is the client API for ModuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModuleServiceClient interface {
	// 根据模块ID获取项目信息
	GetModuleByID(ctx context.Context, in *GetModuleByIDReq, opts ...grpc.CallOption) (*GetModuleByIDRep, error)
	// 根据项目ID获取模块列表
	GetModulesByProjectID(ctx context.Context, in *GetModulesByProjectIDReq, opts ...grpc.CallOption) (*GetModulesByProjectIDRep, error)
	// 创建模块
	CreateModule(ctx context.Context, in *CreateModuleReq, opts ...grpc.CallOption) (*commonpb.Empty, error)
	// 更新模块信息
	UpdateModuleByID(ctx context.Context, in *UpdateModuleByIDReq, opts ...grpc.CallOption) (*commonpb.Empty, error)
	// 删除模块信息
	DeleteModuleByID(ctx context.Context, in *DeleteModuleByIDReq, opts ...grpc.CallOption) (*commonpb.Empty, error)
}

type moduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModuleServiceClient(cc grpc.ClientConnInterface) ModuleServiceClient {
	return &moduleServiceClient{cc}
}

func (c *moduleServiceClient) GetModuleByID(ctx context.Context, in *GetModuleByIDReq, opts ...grpc.CallOption) (*GetModuleByIDRep, error) {
	out := new(GetModuleByIDRep)
	err := c.cc.Invoke(ctx, ModuleService_GetModuleByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) GetModulesByProjectID(ctx context.Context, in *GetModulesByProjectIDReq, opts ...grpc.CallOption) (*GetModulesByProjectIDRep, error) {
	out := new(GetModulesByProjectIDRep)
	err := c.cc.Invoke(ctx, ModuleService_GetModulesByProjectID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) CreateModule(ctx context.Context, in *CreateModuleReq, opts ...grpc.CallOption) (*commonpb.Empty, error) {
	out := new(commonpb.Empty)
	err := c.cc.Invoke(ctx, ModuleService_CreateModule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) UpdateModuleByID(ctx context.Context, in *UpdateModuleByIDReq, opts ...grpc.CallOption) (*commonpb.Empty, error) {
	out := new(commonpb.Empty)
	err := c.cc.Invoke(ctx, ModuleService_UpdateModuleByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleServiceClient) DeleteModuleByID(ctx context.Context, in *DeleteModuleByIDReq, opts ...grpc.CallOption) (*commonpb.Empty, error) {
	out := new(commonpb.Empty)
	err := c.cc.Invoke(ctx, ModuleService_DeleteModuleByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModuleServiceServer is the server API for ModuleService service.
// All implementations must embed UnimplementedModuleServiceServer
// for forward compatibility
type ModuleServiceServer interface {
	// 根据模块ID获取项目信息
	GetModuleByID(context.Context, *GetModuleByIDReq) (*GetModuleByIDRep, error)
	// 根据项目ID获取模块列表
	GetModulesByProjectID(context.Context, *GetModulesByProjectIDReq) (*GetModulesByProjectIDRep, error)
	// 创建模块
	CreateModule(context.Context, *CreateModuleReq) (*commonpb.Empty, error)
	// 更新模块信息
	UpdateModuleByID(context.Context, *UpdateModuleByIDReq) (*commonpb.Empty, error)
	// 删除模块信息
	DeleteModuleByID(context.Context, *DeleteModuleByIDReq) (*commonpb.Empty, error)
	mustEmbedUnimplementedModuleServiceServer()
}

// UnimplementedModuleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModuleServiceServer struct {
}

func (UnimplementedModuleServiceServer) GetModuleByID(context.Context, *GetModuleByIDReq) (*GetModuleByIDRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModuleByID not implemented")
}
func (UnimplementedModuleServiceServer) GetModulesByProjectID(context.Context, *GetModulesByProjectIDReq) (*GetModulesByProjectIDRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModulesByProjectID not implemented")
}
func (UnimplementedModuleServiceServer) CreateModule(context.Context, *CreateModuleReq) (*commonpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModule not implemented")
}
func (UnimplementedModuleServiceServer) UpdateModuleByID(context.Context, *UpdateModuleByIDReq) (*commonpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModuleByID not implemented")
}
func (UnimplementedModuleServiceServer) DeleteModuleByID(context.Context, *DeleteModuleByIDReq) (*commonpb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModuleByID not implemented")
}
func (UnimplementedModuleServiceServer) mustEmbedUnimplementedModuleServiceServer() {}

// UnsafeModuleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModuleServiceServer will
// result in compilation errors.
type UnsafeModuleServiceServer interface {
	mustEmbedUnimplementedModuleServiceServer()
}

func RegisterModuleServiceServer(s grpc.ServiceRegistrar, srv ModuleServiceServer) {
	s.RegisterService(&ModuleService_ServiceDesc, srv)
}

func _ModuleService_GetModuleByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModuleByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetModuleByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleService_GetModuleByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetModuleByID(ctx, req.(*GetModuleByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_GetModulesByProjectID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModulesByProjectIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).GetModulesByProjectID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleService_GetModulesByProjectID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).GetModulesByProjectID(ctx, req.(*GetModulesByProjectIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_CreateModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).CreateModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleService_CreateModule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).CreateModule(ctx, req.(*CreateModuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_UpdateModuleByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModuleByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).UpdateModuleByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleService_UpdateModuleByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).UpdateModuleByID(ctx, req.(*UpdateModuleByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModuleService_DeleteModuleByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModuleByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServiceServer).DeleteModuleByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModuleService_DeleteModuleByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServiceServer).DeleteModuleByID(ctx, req.(*DeleteModuleByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ModuleService_ServiceDesc is the grpc.ServiceDesc for ModuleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModuleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.ModuleService",
	HandlerType: (*ModuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModuleByID",
			Handler:    _ModuleService_GetModuleByID_Handler,
		},
		{
			MethodName: "GetModulesByProjectID",
			Handler:    _ModuleService_GetModulesByProjectID_Handler,
		},
		{
			MethodName: "CreateModule",
			Handler:    _ModuleService_CreateModule_Handler,
		},
		{
			MethodName: "UpdateModuleByID",
			Handler:    _ModuleService_UpdateModuleByID_Handler,
		},
		{
			MethodName: "DeleteModuleByID",
			Handler:    _ModuleService_DeleteModuleByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "projectpb/module.proto",
}
