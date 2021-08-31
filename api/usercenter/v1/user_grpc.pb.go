// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	SendValidateCode(ctx context.Context, in *ValidateCodeReq, opts ...grpc.CallOption) (*ValidateCodeResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ResetPassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ForgetPassword(ctx context.Context, in *ForgetPasswordReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*LoginResp, error)
	GetMembers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMembersResp, error)
	GetPaymentMethod(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPaymentMethodResp, error)
	GetAllUserInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllUserInfoResp, error)
	GetUserMasterByInViteUser(ctx context.Context, in *GetUserMasterReq, opts ...grpc.CallOption) (*UserMasterResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendValidateCode(ctx context.Context, in *ValidateCodeReq, opts ...grpc.CallOption) (*ValidateCodeResp, error) {
	out := new(ValidateCodeResp)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/SendValidateCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ResetPassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ForgetPassword(ctx context.Context, in *ForgetPasswordReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/ForgetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetMembers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMembersResp, error) {
	out := new(GetMembersResp)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/GetMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetPaymentMethod(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPaymentMethodResp, error) {
	out := new(GetPaymentMethodResp)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/GetPaymentMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAllUserInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllUserInfoResp, error) {
	out := new(AllUserInfoResp)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/GetAllUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserMasterByInViteUser(ctx context.Context, in *GetUserMasterReq, opts ...grpc.CallOption) (*UserMasterResp, error) {
	out := new(UserMasterResp)
	err := c.cc.Invoke(ctx, "/api.usercenter.v1.User/GetUserMasterByInViteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	SendValidateCode(context.Context, *ValidateCodeReq) (*ValidateCodeResp, error)
	Register(context.Context, *RegisterReq) (*emptypb.Empty, error)
	ResetPassword(context.Context, *ChangePasswordReq) (*emptypb.Empty, error)
	ForgetPassword(context.Context, *ForgetPasswordReq) (*emptypb.Empty, error)
	UpdateUser(context.Context, *UpdateUserReq) (*emptypb.Empty, error)
	GetUserInfo(context.Context, *UserInfoReq) (*LoginResp, error)
	GetMembers(context.Context, *emptypb.Empty) (*GetMembersResp, error)
	GetPaymentMethod(context.Context, *emptypb.Empty) (*GetPaymentMethodResp, error)
	GetAllUserInfo(context.Context, *emptypb.Empty) (*AllUserInfoResp, error)
	GetUserMasterByInViteUser(context.Context, *GetUserMasterReq) (*UserMasterResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) SendValidateCode(context.Context, *ValidateCodeReq) (*ValidateCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendValidateCode not implemented")
}
func (UnimplementedUserServer) Register(context.Context, *RegisterReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) ResetPassword(context.Context, *ChangePasswordReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServer) ForgetPassword(context.Context, *ForgetPasswordReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgetPassword not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) GetUserInfo(context.Context, *UserInfoReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServer) GetMembers(context.Context, *emptypb.Empty) (*GetMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}
func (UnimplementedUserServer) GetPaymentMethod(context.Context, *emptypb.Empty) (*GetPaymentMethodResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentMethod not implemented")
}
func (UnimplementedUserServer) GetAllUserInfo(context.Context, *emptypb.Empty) (*AllUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserInfo not implemented")
}
func (UnimplementedUserServer) GetUserMasterByInViteUser(context.Context, *GetUserMasterReq) (*UserMasterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMasterByInViteUser not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendValidateCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendValidateCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/SendValidateCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendValidateCode(ctx, req.(*ValidateCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ResetPassword(ctx, req.(*ChangePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ForgetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ForgetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/ForgetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ForgetPassword(ctx, req.(*ForgetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/GetMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetMembers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/GetPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetPaymentMethod(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetAllUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAllUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/GetAllUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAllUserInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserMasterByInViteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserMasterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserMasterByInViteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.usercenter.v1.User/GetUserMasterByInViteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserMasterByInViteUser(ctx, req.(*GetUserMasterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.usercenter.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "SendValidateCode",
			Handler:    _User_SendValidateCode_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _User_ResetPassword_Handler,
		},
		{
			MethodName: "ForgetPassword",
			Handler:    _User_ForgetPassword_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _User_GetUserInfo_Handler,
		},
		{
			MethodName: "GetMembers",
			Handler:    _User_GetMembers_Handler,
		},
		{
			MethodName: "GetPaymentMethod",
			Handler:    _User_GetPaymentMethod_Handler,
		},
		{
			MethodName: "GetAllUserInfo",
			Handler:    _User_GetAllUserInfo_Handler,
		},
		{
			MethodName: "GetUserMasterByInViteUser",
			Handler:    _User_GetUserMasterByInViteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/usercenter/v1/user.proto",
}
