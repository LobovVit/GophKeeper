// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/gophkeeper.proto

// Code generated by protoc-gen-go-grpchandler. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpchandler v1.4.0
// - protoc             v5.26.1
// source: proto/gophkeeper.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpchandler package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	GophKeeper_Ping_FullMethodName           = "/keeper.GophKeeper/Ping"
	GophKeeper_UserExist_FullMethodName      = "/keeper.GophKeeper/UserExist"
	GophKeeper_Authentication_FullMethodName = "/keeper.GophKeeper/Authentication"
	GophKeeper_Registration_FullMethodName   = "/keeper.GophKeeper/Registration"
	GophKeeper_FileUpload_FullMethodName     = "/keeper.GophKeeper/FileUpload"
	GophKeeper_FileGetList_FullMethodName    = "/keeper.GophKeeper/FileGetList"
	GophKeeper_FileRemove_FullMethodName     = "/keeper.GophKeeper/FileRemove"
	GophKeeper_FileDownload_FullMethodName   = "/keeper.GophKeeper/FileDownload"
	GophKeeper_EntityCreate_FullMethodName   = "/keeper.GophKeeper/EntityCreate"
	GophKeeper_EntityGetList_FullMethodName  = "/keeper.GophKeeper/EntityGetList"
	GophKeeper_EntityDelete_FullMethodName   = "/keeper.GophKeeper/EntityDelete"
	GophKeeper_EntityUpdate_FullMethodName   = "/keeper.GophKeeper/EntityUpdate"
)

// GophKeeperClient is the client API for GophKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophKeeperClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	UserExist(ctx context.Context, in *UserExistRequest, opts ...grpc.CallOption) (*UserExistResponse, error)
	Authentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
	Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	FileUpload(ctx context.Context, in *UploadBinaryRequest, opts ...grpc.CallOption) (*UploadBinaryResponse, error)
	FileGetList(ctx context.Context, in *GetListBinaryRequest, opts ...grpc.CallOption) (*GetListBinaryResponse, error)
	FileRemove(ctx context.Context, in *DeleteBinaryRequest, opts ...grpc.CallOption) (*DeleteBinaryResponse, error)
	FileDownload(ctx context.Context, in *DownloadBinaryRequest, opts ...grpc.CallOption) (*DownloadBinaryResponse, error)
	EntityCreate(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error)
	EntityGetList(ctx context.Context, in *GetListEntityRequest, opts ...grpc.CallOption) (*GetListEntityResponse, error)
	EntityDelete(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteEntityResponse, error)
	EntityUpdate(ctx context.Context, in *UpdateEntityRequest, opts ...grpc.CallOption) (*UpdateEntityResponse, error)
}

type gophKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophKeeperClient(cc grpc.ClientConnInterface) GophKeeperClient {
	return &gophKeeperClient{cc}
}

func (c *gophKeeperClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, GophKeeper_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) UserExist(ctx context.Context, in *UserExistRequest, opts ...grpc.CallOption) (*UserExistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserExistResponse)
	err := c.cc.Invoke(ctx, GophKeeper_UserExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) Authentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, GophKeeper_Authentication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, GophKeeper_Registration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) FileUpload(ctx context.Context, in *UploadBinaryRequest, opts ...grpc.CallOption) (*UploadBinaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadBinaryResponse)
	err := c.cc.Invoke(ctx, GophKeeper_FileUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) FileGetList(ctx context.Context, in *GetListBinaryRequest, opts ...grpc.CallOption) (*GetListBinaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListBinaryResponse)
	err := c.cc.Invoke(ctx, GophKeeper_FileGetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) FileRemove(ctx context.Context, in *DeleteBinaryRequest, opts ...grpc.CallOption) (*DeleteBinaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBinaryResponse)
	err := c.cc.Invoke(ctx, GophKeeper_FileRemove_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) FileDownload(ctx context.Context, in *DownloadBinaryRequest, opts ...grpc.CallOption) (*DownloadBinaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadBinaryResponse)
	err := c.cc.Invoke(ctx, GophKeeper_FileDownload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) EntityCreate(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEntityResponse)
	err := c.cc.Invoke(ctx, GophKeeper_EntityCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) EntityGetList(ctx context.Context, in *GetListEntityRequest, opts ...grpc.CallOption) (*GetListEntityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListEntityResponse)
	err := c.cc.Invoke(ctx, GophKeeper_EntityGetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) EntityDelete(ctx context.Context, in *DeleteEntityRequest, opts ...grpc.CallOption) (*DeleteEntityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteEntityResponse)
	err := c.cc.Invoke(ctx, GophKeeper_EntityDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) EntityUpdate(ctx context.Context, in *UpdateEntityRequest, opts ...grpc.CallOption) (*UpdateEntityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEntityResponse)
	err := c.cc.Invoke(ctx, GophKeeper_EntityUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophKeeperServer is the server API for GophKeeper service.
// All implementations must embed UnimplementedGophKeeperServer
// for forward compatibility
type GophKeeperServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	UserExist(context.Context, *UserExistRequest) (*UserExistResponse, error)
	Authentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
	Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	FileUpload(context.Context, *UploadBinaryRequest) (*UploadBinaryResponse, error)
	FileGetList(context.Context, *GetListBinaryRequest) (*GetListBinaryResponse, error)
	FileRemove(context.Context, *DeleteBinaryRequest) (*DeleteBinaryResponse, error)
	FileDownload(context.Context, *DownloadBinaryRequest) (*DownloadBinaryResponse, error)
	EntityCreate(context.Context, *CreateEntityRequest) (*CreateEntityResponse, error)
	EntityGetList(context.Context, *GetListEntityRequest) (*GetListEntityResponse, error)
	EntityDelete(context.Context, *DeleteEntityRequest) (*DeleteEntityResponse, error)
	EntityUpdate(context.Context, *UpdateEntityRequest) (*UpdateEntityResponse, error)
	mustEmbedUnimplementedGophKeeperServer()
}

// UnimplementedGophKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophKeeperServer struct {
}

func (UnimplementedGophKeeperServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGophKeeperServer) UserExist(context.Context, *UserExistRequest) (*UserExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserExist not implemented")
}
func (UnimplementedGophKeeperServer) Authentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authentication not implemented")
}
func (UnimplementedGophKeeperServer) Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedGophKeeperServer) FileUpload(context.Context, *UploadBinaryRequest) (*UploadBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileUpload not implemented")
}
func (UnimplementedGophKeeperServer) FileGetList(context.Context, *GetListBinaryRequest) (*GetListBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileGetList not implemented")
}
func (UnimplementedGophKeeperServer) FileRemove(context.Context, *DeleteBinaryRequest) (*DeleteBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileRemove not implemented")
}
func (UnimplementedGophKeeperServer) FileDownload(context.Context, *DownloadBinaryRequest) (*DownloadBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileDownload not implemented")
}
func (UnimplementedGophKeeperServer) EntityCreate(context.Context, *CreateEntityRequest) (*CreateEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntityCreate not implemented")
}
func (UnimplementedGophKeeperServer) EntityGetList(context.Context, *GetListEntityRequest) (*GetListEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntityGetList not implemented")
}
func (UnimplementedGophKeeperServer) EntityDelete(context.Context, *DeleteEntityRequest) (*DeleteEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntityDelete not implemented")
}
func (UnimplementedGophKeeperServer) EntityUpdate(context.Context, *UpdateEntityRequest) (*UpdateEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntityUpdate not implemented")
}
func (UnimplementedGophKeeperServer) mustEmbedUnimplementedGophKeeperServer() {}

// UnsafeGophKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophKeeperServer will
// result in compilation errors.
type UnsafeGophKeeperServer interface {
	mustEmbedUnimplementedGophKeeperServer()
}

func RegisterGophKeeperServer(s grpc.ServiceRegistrar, srv GophKeeperServer) {
	s.RegisterService(&GophKeeper_ServiceDesc, srv)
}

func _GophKeeper_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_UserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).UserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_UserExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).UserExist(ctx, req.(*UserExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_Authentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Authentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Authentication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Authentication(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Registration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Registration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_FileUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).FileUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_FileUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).FileUpload(ctx, req.(*UploadBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_FileGetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).FileGetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_FileGetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).FileGetList(ctx, req.(*GetListBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_FileRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).FileRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_FileRemove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).FileRemove(ctx, req.(*DeleteBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_FileDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).FileDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_FileDownload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).FileDownload(ctx, req.(*DownloadBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_EntityCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).EntityCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_EntityCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).EntityCreate(ctx, req.(*CreateEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_EntityGetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).EntityGetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_EntityGetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).EntityGetList(ctx, req.(*GetListEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_EntityDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).EntityDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_EntityDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).EntityDelete(ctx, req.(*DeleteEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_EntityUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).EntityUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_EntityUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).EntityUpdate(ctx, req.(*UpdateEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GophKeeper_ServiceDesc is the grpc.ServiceDesc for GophKeeper service.
// It's only intended for direct use with grpchandler.RegisterService,
// and not to be introspected or modified (even as a copy)
var GophKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.GophKeeper",
	HandlerType: (*GophKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GophKeeper_Ping_Handler,
		},
		{
			MethodName: "UserExist",
			Handler:    _GophKeeper_UserExist_Handler,
		},
		{
			MethodName: "Authentication",
			Handler:    _GophKeeper_Authentication_Handler,
		},
		{
			MethodName: "Registration",
			Handler:    _GophKeeper_Registration_Handler,
		},
		{
			MethodName: "FileUpload",
			Handler:    _GophKeeper_FileUpload_Handler,
		},
		{
			MethodName: "FileGetList",
			Handler:    _GophKeeper_FileGetList_Handler,
		},
		{
			MethodName: "FileRemove",
			Handler:    _GophKeeper_FileRemove_Handler,
		},
		{
			MethodName: "FileDownload",
			Handler:    _GophKeeper_FileDownload_Handler,
		},
		{
			MethodName: "EntityCreate",
			Handler:    _GophKeeper_EntityCreate_Handler,
		},
		{
			MethodName: "EntityGetList",
			Handler:    _GophKeeper_EntityGetList_Handler,
		},
		{
			MethodName: "EntityDelete",
			Handler:    _GophKeeper_EntityDelete_Handler,
		},
		{
			MethodName: "EntityUpdate",
			Handler:    _GophKeeper_EntityUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gophkeeper.proto",
}