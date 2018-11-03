// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cutter.proto

package cutter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CutRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CutRequest) Reset()         { *m = CutRequest{} }
func (m *CutRequest) String() string { return proto.CompactTextString(m) }
func (*CutRequest) ProtoMessage()    {}
func (*CutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2c736eb4a04fdb, []int{0}
}

func (m *CutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CutRequest.Unmarshal(m, b)
}
func (m *CutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CutRequest.Marshal(b, m, deterministic)
}
func (m *CutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CutRequest.Merge(m, src)
}
func (m *CutRequest) XXX_Size() int {
	return xxx_messageInfo_CutRequest.Size(m)
}
func (m *CutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CutRequest proto.InternalMessageInfo

func (m *CutRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type CutReply struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CutReply) Reset()         { *m = CutReply{} }
func (m *CutReply) String() string { return proto.CompactTextString(m) }
func (*CutReply) ProtoMessage()    {}
func (*CutReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef2c736eb4a04fdb, []int{1}
}

func (m *CutReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CutReply.Unmarshal(m, b)
}
func (m *CutReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CutReply.Marshal(b, m, deterministic)
}
func (m *CutReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CutReply.Merge(m, src)
}
func (m *CutReply) XXX_Size() int {
	return xxx_messageInfo_CutReply.Size(m)
}
func (m *CutReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CutReply.DiscardUnknown(m)
}

var xxx_messageInfo_CutReply proto.InternalMessageInfo

func (m *CutReply) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*CutRequest)(nil), "cutter.CutRequest")
	proto.RegisterType((*CutReply)(nil), "cutter.CutReply")
}

func init() { proto.RegisterFile("cutter.proto", fileDescriptor_ef2c736eb4a04fdb) }

var fileDescriptor_ef2c736eb4a04fdb = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2e, 0x2d, 0x29,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x14, 0xb8, 0xb8,
	0x9c, 0x4b, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x52, 0x12,
	0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x25, 0x2e, 0x0e, 0xb0,
	0x8a, 0x82, 0x9c, 0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0xb0, 0x0a,
	0xd6, 0x20, 0x28, 0xcf, 0xc8, 0x86, 0x8b, 0xcd, 0x19, 0x6c, 0x9e, 0x90, 0x11, 0x58, 0xb5, 0x67,
	0x6e, 0x62, 0x7a, 0xaa, 0x90, 0x90, 0x1e, 0xd4, 0x4a, 0x84, 0x0d, 0x52, 0x02, 0x28, 0x62, 0x05,
	0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x27, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc7,
	0xe9, 0x47, 0x9f, 0xa2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CutterClient is the client API for Cutter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CutterClient interface {
	// Sends a greeting
	CutImage(ctx context.Context, in *CutRequest, opts ...grpc.CallOption) (*CutReply, error)
}

type cutterClient struct {
	cc *grpc.ClientConn
}

func NewCutterClient(cc *grpc.ClientConn) CutterClient {
	return &cutterClient{cc}
}

func (c *cutterClient) CutImage(ctx context.Context, in *CutRequest, opts ...grpc.CallOption) (*CutReply, error) {
	out := new(CutReply)
	err := c.cc.Invoke(ctx, "/cutter.Cutter/CutImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CutterServer is the server API for Cutter service.
type CutterServer interface {
	// Sends a greeting
	CutImage(context.Context, *CutRequest) (*CutReply, error)
}

func RegisterCutterServer(s *grpc.Server, srv CutterServer) {
	s.RegisterService(&_Cutter_serviceDesc, srv)
}

func _Cutter_CutImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CutterServer).CutImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cutter.Cutter/CutImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CutterServer).CutImage(ctx, req.(*CutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cutter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cutter.Cutter",
	HandlerType: (*CutterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CutImage",
			Handler:    _Cutter_CutImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cutter.proto",
}