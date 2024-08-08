// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/proto/serv.proto

package api

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

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	GenerateData(ctx context.Context, opts ...grpc.CallOption) (DataService_GenerateDataClient, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) GenerateData(ctx context.Context, opts ...grpc.CallOption) (DataService_GenerateDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[0], "/api.DataService/GenerateData", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceGenerateDataClient{stream}
	return x, nil
}

type DataService_GenerateDataClient interface {
	Send(*DataEntry) error
	Recv() (*DataEntry, error)
	grpc.ClientStream
}

type dataServiceGenerateDataClient struct {
	grpc.ClientStream
}

func (x *dataServiceGenerateDataClient) Send(m *DataEntry) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataServiceGenerateDataClient) Recv() (*DataEntry, error) {
	m := new(DataEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	GenerateData(DataService_GenerateDataServer) error
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) GenerateData(DataService_GenerateDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GenerateData not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_GenerateData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServiceServer).GenerateData(&dataServiceGenerateDataServer{stream})
}

type DataService_GenerateDataServer interface {
	Send(*DataEntry) error
	Recv() (*DataEntry, error)
	grpc.ServerStream
}

type dataServiceGenerateDataServer struct {
	grpc.ServerStream
}

func (x *dataServiceGenerateDataServer) Send(m *DataEntry) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataServiceGenerateDataServer) Recv() (*DataEntry, error) {
	m := new(DataEntry)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateData",
			Handler:       _DataService_GenerateData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/serv.proto",
}