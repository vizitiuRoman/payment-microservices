// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/invoice.proto

package invoice

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

// InvoiceServiceClient is the client API for InvoiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoiceServiceClient interface {
	GetInvoice(ctx context.Context, in *GetInvoiceInput, opts ...grpc.CallOption) (*InvoiceOutput, error)
	CreateInvoice(ctx context.Context, in *CreateInvoiceInput, opts ...grpc.CallOption) (*InvoiceOutput, error)
}

type invoiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceServiceClient(cc grpc.ClientConnInterface) InvoiceServiceClient {
	return &invoiceServiceClient{cc}
}

func (c *invoiceServiceClient) GetInvoice(ctx context.Context, in *GetInvoiceInput, opts ...grpc.CallOption) (*InvoiceOutput, error) {
	out := new(InvoiceOutput)
	err := c.cc.Invoke(ctx, "/proto.InvoiceService/GetInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) CreateInvoice(ctx context.Context, in *CreateInvoiceInput, opts ...grpc.CallOption) (*InvoiceOutput, error) {
	out := new(InvoiceOutput)
	err := c.cc.Invoke(ctx, "/proto.InvoiceService/CreateInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServiceServer is the server API for InvoiceService service.
// All implementations must embed UnimplementedInvoiceServiceServer
// for forward compatibility
type InvoiceServiceServer interface {
	GetInvoice(context.Context, *GetInvoiceInput) (*InvoiceOutput, error)
	CreateInvoice(context.Context, *CreateInvoiceInput) (*InvoiceOutput, error)
	mustEmbedUnimplementedInvoiceServiceServer()
}

// UnimplementedInvoiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInvoiceServiceServer struct {
}

func (UnimplementedInvoiceServiceServer) GetInvoice(context.Context, *GetInvoiceInput) (*InvoiceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) CreateInvoice(context.Context, *CreateInvoiceInput) (*InvoiceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) mustEmbedUnimplementedInvoiceServiceServer() {}

// UnsafeInvoiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceServiceServer will
// result in compilation errors.
type UnsafeInvoiceServiceServer interface {
	mustEmbedUnimplementedInvoiceServiceServer()
}

func RegisterInvoiceServiceServer(s grpc.ServiceRegistrar, srv InvoiceServiceServer) {
	s.RegisterService(&InvoiceService_ServiceDesc, srv)
}

func _InvoiceService_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.InvoiceService/GetInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoice(ctx, req.(*GetInvoiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_CreateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.InvoiceService/CreateInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateInvoice(ctx, req.(*CreateInvoiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

// InvoiceService_ServiceDesc is the grpc.ServiceDesc for InvoiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InvoiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.InvoiceService",
	HandlerType: (*InvoiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInvoice",
			Handler:    _InvoiceService_GetInvoice_Handler,
		},
		{
			MethodName: "CreateInvoice",
			Handler:    _InvoiceService_CreateInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/invoice.proto",
}
