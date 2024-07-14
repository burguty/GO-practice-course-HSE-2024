// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: proto/echo.proto

package proto

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

// BankClient is the client API for Bank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountReply, error)
	UpdateAmount(ctx context.Context, in *UpdateAmountRequest, opts ...grpc.CallOption) (*UpdateAmountReply, error)
	UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...grpc.CallOption) (*UpdateNameReply, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error)
}

type bankClient struct {
	cc grpc.ClientConnInterface
}

func NewBankClient(cc grpc.ClientConnInterface) BankClient {
	return &bankClient{cc}
}

func (c *bankClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error) {
	out := new(CreateAccountReply)
	err := c.cc.Invoke(ctx, "/proto.Bank/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountReply, error) {
	out := new(DeleteAccountReply)
	err := c.cc.Invoke(ctx, "/proto.Bank/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) UpdateAmount(ctx context.Context, in *UpdateAmountRequest, opts ...grpc.CallOption) (*UpdateAmountReply, error) {
	out := new(UpdateAmountReply)
	err := c.cc.Invoke(ctx, "/proto.Bank/UpdateAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...grpc.CallOption) (*UpdateNameReply, error) {
	out := new(UpdateNameReply)
	err := c.cc.Invoke(ctx, "/proto.Bank/UpdateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error) {
	out := new(GetAccountReply)
	err := c.cc.Invoke(ctx, "/proto.Bank/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServer is the server API for Bank service.
// All implementations must embed UnimplementedBankServer
// for forward compatibility
type BankServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountReply, error)
	UpdateAmount(context.Context, *UpdateAmountRequest) (*UpdateAmountReply, error)
	UpdateName(context.Context, *UpdateNameRequest) (*UpdateNameReply, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountReply, error)
	mustEmbedUnimplementedBankServer()
}

// UnimplementedBankServer must be embedded to have forward compatible implementations.
type UnimplementedBankServer struct {
}

func (UnimplementedBankServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedBankServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedBankServer) UpdateAmount(context.Context, *UpdateAmountRequest) (*UpdateAmountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAmount not implemented")
}
func (UnimplementedBankServer) UpdateName(context.Context, *UpdateNameRequest) (*UpdateNameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedBankServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedBankServer) mustEmbedUnimplementedBankServer() {}

// UnsafeBankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServer will
// result in compilation errors.
type UnsafeBankServer interface {
	mustEmbedUnimplementedBankServer()
}

func RegisterBankServer(s grpc.ServiceRegistrar, srv BankServer) {
	s.RegisterService(&Bank_ServiceDesc, srv)
}

func _Bank_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bank/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bank/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_UpdateAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).UpdateAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bank/UpdateAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).UpdateAmount(ctx, req.(*UpdateAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bank/UpdateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).UpdateName(ctx, req.(*UpdateNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bank/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bank_ServiceDesc is the grpc.ServiceDesc for Bank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Bank",
	HandlerType: (*BankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Bank_CreateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Bank_DeleteAccount_Handler,
		},
		{
			MethodName: "UpdateAmount",
			Handler:    _Bank_UpdateAmount_Handler,
		},
		{
			MethodName: "UpdateName",
			Handler:    _Bank_UpdateName_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Bank_GetAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/echo.proto",
}
