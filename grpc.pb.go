// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package main

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

type TransactionSignature struct {
	TrxPrefix            string   `protobuf:"bytes,1,opt,name=trxPrefix,proto3" json:"trxPrefix,omitempty"`
	Signature            string   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Account              string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionSignature) Reset()         { *m = TransactionSignature{} }
func (m *TransactionSignature) String() string { return proto.CompactTextString(m) }
func (*TransactionSignature) ProtoMessage()    {}
func (*TransactionSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *TransactionSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionSignature.Unmarshal(m, b)
}
func (m *TransactionSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionSignature.Marshal(b, m, deterministic)
}
func (m *TransactionSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionSignature.Merge(m, src)
}
func (m *TransactionSignature) XXX_Size() int {
	return xxx_messageInfo_TransactionSignature.Size(m)
}
func (m *TransactionSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionSignature.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionSignature proto.InternalMessageInfo

func (m *TransactionSignature) GetTrxPrefix() string {
	if m != nil {
		return m.TrxPrefix
	}
	return ""
}

func (m *TransactionSignature) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *TransactionSignature) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type Receipt struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receipt.Unmarshal(m, b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return xxx_messageInfo_Receipt.Size(m)
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*TransactionSignature)(nil), "main.TransactionSignature")
	proto.RegisterType((*Receipt)(nil), "main.Receipt")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x31, 0xcf, 0x82, 0x30,
	0x10, 0x86, 0x3f, 0xf8, 0x8c, 0xc0, 0x19, 0x1d, 0x1a, 0x87, 0x4a, 0x1c, 0x0c, 0x93, 0x13, 0x83,
	0x0e, 0xce, 0x26, 0xc6, 0xd9, 0xa0, 0x7f, 0xa0, 0x36, 0x27, 0x69, 0xc0, 0xb6, 0xb9, 0x96, 0x84,
	0x9f, 0x6f, 0x00, 0xd1, 0xc5, 0xf1, 0xde, 0xe7, 0xf2, 0xe6, 0x79, 0x01, 0x4a, 0xb2, 0x32, 0xb7,
	0x64, 0xbc, 0x61, 0x93, 0xa7, 0x50, 0x3a, 0xab, 0x61, 0x79, 0x23, 0xa1, 0x9d, 0x90, 0x5e, 0x19,
	0x7d, 0x55, 0xa5, 0x16, 0xbe, 0x21, 0x64, 0x6b, 0x48, 0x3c, 0xb5, 0x17, 0xc2, 0x87, 0x6a, 0x79,
	0xb0, 0x09, 0xb6, 0x49, 0xf1, 0x0d, 0x3a, 0xea, 0xc6, 0x57, 0x1e, 0x0e, 0xf4, 0x13, 0x30, 0x0e,
	0x91, 0x90, 0xd2, 0x34, 0xda, 0xf3, 0xff, 0x9e, 0x8d, 0x67, 0xb6, 0x82, 0xa8, 0x40, 0x89, 0xca,
	0x7a, 0xb6, 0x80, 0xd0, 0x54, 0x7d, 0x73, 0x5c, 0x84, 0xa6, 0xda, 0x9d, 0x61, 0x76, 0x94, 0x12,
	0x6b, 0x24, 0xe1, 0x0d, 0xb1, 0x03, 0xc4, 0x27, 0xe5, 0x2c, 0x92, 0x43, 0x96, 0xe6, 0x9d, 0x6a,
	0xfe, 0xcb, 0x33, 0x9d, 0x0f, 0xec, 0xdd, 0x9a, 0xfd, 0xdd, 0xa7, 0xfd, 0xba, 0xfd, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x23, 0x67, 0x7d, 0x55, 0xeb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AcceleratorClient is the client API for Accelerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AcceleratorClient interface {
	Disperse(ctx context.Context, in *TransactionSignature, opts ...grpc.CallOption) (*Receipt, error)
}

type acceleratorClient struct {
	cc *grpc.ClientConn
}

func NewAcceleratorClient(cc *grpc.ClientConn) AcceleratorClient {
	return &acceleratorClient{cc}
}

func (c *acceleratorClient) Disperse(ctx context.Context, in *TransactionSignature, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/main.Accelerator/Disperse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AcceleratorServer is the server API for Accelerator service.
type AcceleratorServer interface {
	Disperse(context.Context, *TransactionSignature) (*Receipt, error)
}

func RegisterAcceleratorServer(s *grpc.Server, srv AcceleratorServer) {
	s.RegisterService(&_Accelerator_serviceDesc, srv)
}

func _Accelerator_Disperse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionSignature)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcceleratorServer).Disperse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Accelerator/Disperse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcceleratorServer).Disperse(ctx, req.(*TransactionSignature))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accelerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Accelerator",
	HandlerType: (*AcceleratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Disperse",
			Handler:    _Accelerator_Disperse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
