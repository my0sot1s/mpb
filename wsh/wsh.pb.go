// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wsh/wsh.proto

/*
Package wsh is a generated protocol buffer package.

It is generated from these files:
	wsh/wsh.proto

It has these top-level messages:
	Resp
	Detail
	Message
*/
package wsh

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Resp struct {
	Created int64 `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Resp) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type Detail struct {
	Device  string `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	Author  string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Created int64  `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	Service string `protobuf:"bytes,4,opt,name=service" json:"service,omitempty"`
}

func (m *Detail) Reset()                    { *m = Detail{} }
func (m *Detail) String() string            { return proto.CompactTextString(m) }
func (*Detail) ProtoMessage()               {}
func (*Detail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Detail) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *Detail) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Detail) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Detail) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type Message struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Topic  string  `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Body   string  `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Detail *Detail `protobuf:"bytes,4,opt,name=detail" json:"detail,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Message) GetDetail() *Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func init() {
	proto.RegisterType((*Resp)(nil), "wsh.Resp")
	proto.RegisterType((*Detail)(nil), "wsh.Detail")
	proto.RegisterType((*Message)(nil), "wsh.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WsProvider service

type WsProviderClient interface {
	Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Resp, error)
}

type wsProviderClient struct {
	cc *grpc.ClientConn
}

func NewWsProviderClient(cc *grpc.ClientConn) WsProviderClient {
	return &wsProviderClient{cc}
}

func (c *wsProviderClient) Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/wsh.WsProvider/Broadcast", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WsProvider service

type WsProviderServer interface {
	Broadcast(context.Context, *Message) (*Resp, error)
}

func RegisterWsProviderServer(s *grpc.Server, srv WsProviderServer) {
	s.RegisterService(&_WsProvider_serviceDesc, srv)
}

func _WsProvider_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WsProviderServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wsh.WsProvider/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WsProviderServer).Broadcast(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _WsProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wsh.WsProvider",
	HandlerType: (*WsProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _WsProvider_Broadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wsh/wsh.proto",
}

func init() { proto.RegisterFile("wsh/wsh.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xd9, 0xb6, 0x76, 0xe9, 0xac, 0x7a, 0x18, 0x44, 0x8a, 0xa7, 0xa5, 0x7a, 0xf0, 0xb4,
	0x42, 0xfd, 0x07, 0xe2, 0x55, 0x90, 0x5c, 0x3c, 0x67, 0x9b, 0xc1, 0x04, 0x16, 0x52, 0x32, 0x71,
	0x8b, 0xff, 0x5e, 0x32, 0x49, 0xc1, 0xbd, 0xe5, 0x9b, 0x99, 0xbc, 0xf7, 0x66, 0xe0, 0x66, 0x61,
	0xfb, 0xb2, 0xb0, 0x3d, 0xcc, 0xc1, 0x47, 0x8f, 0xf5, 0xc2, 0x76, 0xd8, 0x43, 0xa3, 0x88, 0x67,
	0xec, 0x61, 0x3b, 0x05, 0xd2, 0x91, 0x4c, 0xbf, 0xd9, 0x6f, 0x9e, 0x6b, 0xb5, 0xe2, 0x70, 0x82,
	0xf6, 0x9d, 0xa2, 0x76, 0x27, 0xbc, 0x87, 0xd6, 0xd0, 0xd9, 0x4d, 0x24, 0x23, 0x9d, 0x2a, 0x94,
	0xea, 0xfa, 0x27, 0x5a, 0x1f, 0xfa, 0x2a, 0xd7, 0x33, 0xfd, 0xd7, 0xac, 0x2f, 0x34, 0x53, 0x87,
	0x29, 0x88, 0x54, 0x23, 0x5f, 0x56, 0x1c, 0x2c, 0x6c, 0x3f, 0x88, 0x59, 0x7f, 0x13, 0xde, 0x42,
	0xe5, 0x4c, 0xb1, 0xaa, 0x9c, 0xc1, 0x3b, 0xb8, 0x8a, 0x7e, 0x76, 0x53, 0x71, 0xc9, 0x80, 0x08,
	0xcd, 0xd1, 0x9b, 0x5f, 0x71, 0xe8, 0x94, 0xbc, 0xf1, 0x31, 0x05, 0x4d, 0x91, 0x45, 0x7d, 0x37,
	0xee, 0x0e, 0x69, 0xeb, 0xbc, 0x85, 0x2a, 0xad, 0x71, 0x04, 0xf8, 0xe2, 0xcf, 0xe0, 0xcf, 0xce,
	0x50, 0xc0, 0x27, 0xe8, 0xde, 0x82, 0xd7, 0x66, 0xd2, 0x1c, 0xf1, 0x5a, 0xe6, 0x4b, 0x8e, 0x87,
	0x4e, 0x28, 0x5d, 0xe9, 0xd8, 0xca, 0xe5, 0x5e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x17, 0x66,
	0x26, 0x2c, 0x4a, 0x01, 0x00, 0x00,
}
