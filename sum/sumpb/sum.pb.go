// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum/sumpb/sum.proto

package sumpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("sum/sumpb/sum.proto", fileDescriptor_1e82cd6e6acf87a3) }

var fileDescriptor_1e82cd6e6acf87a3 = []byte{
	// 73 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0xcd, 0xd5,
	0x2f, 0x2e, 0xcd, 0x2d, 0x48, 0x02, 0x91, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xcc, 0xc5,
	0xa5, 0xb9, 0x46, 0x3c, 0x5c, 0x5c, 0xc1, 0xa5, 0xb9, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9,
	0x4e, 0xec, 0x51, 0xac, 0x60, 0x55, 0x49, 0x6c, 0x60, 0x25, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe2, 0xcc, 0x9b, 0x20, 0x39, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
}

type sumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServiceClient(cc grpc.ClientConnInterface) SumServiceClient {
	return &sumServiceClient{cc}
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sum.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "sum/sumpb/sum.proto",
}