// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: oyster/rpc/payment/v1/payment_info_rpc.proto

package paymentv1

import (
	context "context"
	v1 "github.com/yaydoo/proto.boilerplate/go/oyster/entities/pb/payment/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PaymentInfoServiceClient is the client API for PaymentInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentInfoServiceClient interface {
	GetPaymentTransaction(ctx context.Context, in *v1.GetPaymentTransactionRequest, opts ...grpc.CallOption) (*v1.GetPaymentTransactionResponse, error)
}

type paymentInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentInfoServiceClient(cc grpc.ClientConnInterface) PaymentInfoServiceClient {
	return &paymentInfoServiceClient{cc}
}

func (c *paymentInfoServiceClient) GetPaymentTransaction(ctx context.Context, in *v1.GetPaymentTransactionRequest, opts ...grpc.CallOption) (*v1.GetPaymentTransactionResponse, error) {
	out := new(v1.GetPaymentTransactionResponse)
	err := c.cc.Invoke(ctx, "/oyster.rpc.payment.v1.PaymentInfoService/GetPaymentTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentInfoServiceServer is the server API for PaymentInfoService service.
// All implementations should embed UnimplementedPaymentInfoServiceServer
// for forward compatibility
type PaymentInfoServiceServer interface {
	GetPaymentTransaction(context.Context, *v1.GetPaymentTransactionRequest) (*v1.GetPaymentTransactionResponse, error)
}

// UnimplementedPaymentInfoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPaymentInfoServiceServer struct {
}

func (UnimplementedPaymentInfoServiceServer) GetPaymentTransaction(context.Context, *v1.GetPaymentTransactionRequest) (*v1.GetPaymentTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentTransaction not implemented")
}

// UnsafePaymentInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentInfoServiceServer will
// result in compilation errors.
type UnsafePaymentInfoServiceServer interface {
	mustEmbedUnimplementedPaymentInfoServiceServer()
}

func RegisterPaymentInfoServiceServer(s grpc.ServiceRegistrar, srv PaymentInfoServiceServer) {
	s.RegisterService(&PaymentInfoService_ServiceDesc, srv)
}

func _PaymentInfoService_GetPaymentTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetPaymentTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentInfoServiceServer).GetPaymentTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oyster.rpc.payment.v1.PaymentInfoService/GetPaymentTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentInfoServiceServer).GetPaymentTransaction(ctx, req.(*v1.GetPaymentTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentInfoService_ServiceDesc is the grpc.ServiceDesc for PaymentInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oyster.rpc.payment.v1.PaymentInfoService",
	HandlerType: (*PaymentInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPaymentTransaction",
			Handler:    _PaymentInfoService_GetPaymentTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oyster/rpc/payment/v1/payment_info_rpc.proto",
}