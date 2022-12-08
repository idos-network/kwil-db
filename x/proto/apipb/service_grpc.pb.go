// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: kwil/apisvc/service.proto

package apipb

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

// KwilServiceClient is the client API for KwilService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KwilServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	DeploySchema(ctx context.Context, in *DeploySchemaRequest, opts ...grpc.CallOption) (*DeploySchemaResponse, error)
	GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error)
	Cud(ctx context.Context, in *CUDRequest, opts ...grpc.CallOption) (*CUDResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Wallets
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	ReturnFunds(ctx context.Context, in *ReturnFundsRequest, opts ...grpc.CallOption) (*ReturnFundsResponse, error)
	EstimateCost(ctx context.Context, in *EstimateCostRequest, opts ...grpc.CallOption) (*EstimateCostResponse, error)
}

type kwilServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKwilServiceClient(cc grpc.ClientConnInterface) KwilServiceClient {
	return &kwilServiceClient{cc}
}

func (c *kwilServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/apisvc.KwilService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kwilServiceClient) DeploySchema(ctx context.Context, in *DeploySchemaRequest, opts ...grpc.CallOption) (*DeploySchemaResponse, error) {
	out := new(DeploySchemaResponse)
	err := c.cc.Invoke(ctx, "/apisvc.KwilService/DeploySchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kwilServiceClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error) {
	out := new(GetMetadataResponse)
	err := c.cc.Invoke(ctx, "/apisvc.KwilService/GetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kwilServiceClient) Cud(ctx context.Context, in *CUDRequest, opts ...grpc.CallOption) (*CUDResponse, error) {
	out := new(CUDResponse)
	err := c.cc.Invoke(ctx, "/apisvc.KwilService/Cud", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kwilServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/apisvc.KwilService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kwilServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/apisvc.KwilService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kwilServiceClient) ReturnFunds(ctx context.Context, in *ReturnFundsRequest, opts ...grpc.CallOption) (*ReturnFundsResponse, error) {
	out := new(ReturnFundsResponse)
	err := c.cc.Invoke(ctx, "/apisvc.KwilService/ReturnFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kwilServiceClient) EstimateCost(ctx context.Context, in *EstimateCostRequest, opts ...grpc.CallOption) (*EstimateCostResponse, error) {
	out := new(EstimateCostResponse)
	err := c.cc.Invoke(ctx, "/apisvc.KwilService/EstimateCost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KwilServiceServer is the server API for KwilService service.
// All implementations must embed UnimplementedKwilServiceServer
// for forward compatibility
type KwilServiceServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	DeploySchema(context.Context, *DeploySchemaRequest) (*DeploySchemaResponse, error)
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)
	Cud(context.Context, *CUDRequest) (*CUDResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Wallets
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	ReturnFunds(context.Context, *ReturnFundsRequest) (*ReturnFundsResponse, error)
	EstimateCost(context.Context, *EstimateCostRequest) (*EstimateCostResponse, error)
	mustEmbedUnimplementedKwilServiceServer()
}

// UnimplementedKwilServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKwilServiceServer struct {
}

func (UnimplementedKwilServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedKwilServiceServer) DeploySchema(context.Context, *DeploySchemaRequest) (*DeploySchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploySchema not implemented")
}
func (UnimplementedKwilServiceServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedKwilServiceServer) Cud(context.Context, *CUDRequest) (*CUDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cud not implemented")
}
func (UnimplementedKwilServiceServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedKwilServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedKwilServiceServer) ReturnFunds(context.Context, *ReturnFundsRequest) (*ReturnFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnFunds not implemented")
}
func (UnimplementedKwilServiceServer) EstimateCost(context.Context, *EstimateCostRequest) (*EstimateCostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateCost not implemented")
}
func (UnimplementedKwilServiceServer) mustEmbedUnimplementedKwilServiceServer() {}

// UnsafeKwilServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KwilServiceServer will
// result in compilation errors.
type UnsafeKwilServiceServer interface {
	mustEmbedUnimplementedKwilServiceServer()
}

func RegisterKwilServiceServer(s grpc.ServiceRegistrar, srv KwilServiceServer) {
	s.RegisterService(&KwilService_ServiceDesc, srv)
}

func _KwilService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KwilServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apisvc.KwilService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KwilServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KwilService_DeploySchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploySchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KwilServiceServer).DeploySchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apisvc.KwilService/DeploySchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KwilServiceServer).DeploySchema(ctx, req.(*DeploySchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KwilService_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KwilServiceServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apisvc.KwilService/GetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KwilServiceServer).GetMetadata(ctx, req.(*GetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KwilService_Cud_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KwilServiceServer).Cud(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apisvc.KwilService/Cud",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KwilServiceServer).Cud(ctx, req.(*CUDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KwilService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KwilServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apisvc.KwilService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KwilServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KwilService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KwilServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apisvc.KwilService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KwilServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KwilService_ReturnFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnFundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KwilServiceServer).ReturnFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apisvc.KwilService/ReturnFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KwilServiceServer).ReturnFunds(ctx, req.(*ReturnFundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KwilService_EstimateCost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateCostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KwilServiceServer).EstimateCost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apisvc.KwilService/EstimateCost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KwilServiceServer).EstimateCost(ctx, req.(*EstimateCostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KwilService_ServiceDesc is the grpc.ServiceDesc for KwilService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KwilService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apisvc.KwilService",
	HandlerType: (*KwilServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _KwilService_Connect_Handler,
		},
		{
			MethodName: "DeploySchema",
			Handler:    _KwilService_DeploySchema_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _KwilService_GetMetadata_Handler,
		},
		{
			MethodName: "Cud",
			Handler:    _KwilService_Cud_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _KwilService_Read_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _KwilService_GetBalance_Handler,
		},
		{
			MethodName: "ReturnFunds",
			Handler:    _KwilService_ReturnFunds_Handler,
		},
		{
			MethodName: "EstimateCost",
			Handler:    _KwilService_EstimateCost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kwil/apisvc/service.proto",
}
