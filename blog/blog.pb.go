// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog.proto

package blog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Products             string   `protobuf:"bytes,2,opt,name=products,proto3" json:"products,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price                int32    `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Author               string   `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	Created              string   `protobuf:"bytes,7,opt,name=created,proto3" json:"created,omitempty"`
	State                int32    `protobuf:"varint,8,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f8408f4f97d27b79, []int{0}
}
func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (dst *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(dst, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetProducts() string {
	if m != nil {
		return m.Products
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Blog) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Blog) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Blog) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Blog) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
}

func init() { proto.RegisterFile("blog.proto", fileDescriptor_blog_f8408f4f97d27b79) }

var fileDescriptor_blog_f8408f4f97d27b79 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x4f, 0xca, 0xc2, 0x40,
	0x0c, 0xc5, 0x99, 0x7e, 0xfd, 0xf7, 0x45, 0x70, 0x11, 0x44, 0x82, 0xab, 0xe2, 0xaa, 0x2b, 0x37,
	0xde, 0xc0, 0x23, 0xf4, 0x06, 0xed, 0xcc, 0x50, 0x07, 0x8a, 0x19, 0xa6, 0xe9, 0x1d, 0x3d, 0x96,
	0x34, 0x55, 0x71, 0x97, 0xdf, 0xfb, 0xf1, 0x1e, 0x04, 0x60, 0x98, 0x78, 0xbc, 0xc4, 0xc4, 0xc2,
	0x98, 0xaf, 0xf7, 0xf9, 0x69, 0x20, 0xbf, 0x4d, 0x3c, 0xe2, 0x1e, 0xb2, 0xe0, 0xc8, 0x34, 0xa6,
	0xfd, 0xef, 0xb2, 0xe0, 0xf0, 0x04, 0x75, 0x4c, 0xec, 0x16, 0x2b, 0x33, 0x65, 0x9a, 0x7e, 0x19,
	0x0f, 0x50, 0x48, 0x90, 0xc9, 0xd3, 0x9f, 0x8a, 0x0d, 0xb0, 0x81, 0x9d, 0xf3, 0xb3, 0x4d, 0x21,
	0x4a, 0xe0, 0x07, 0xe5, 0xea, 0x7e, 0xa3, 0xb5, 0x17, 0x53, 0xb0, 0x9e, 0x8a, 0xc6, 0xb4, 0x45,
	0xb7, 0x01, 0x1e, 0xa1, 0xec, 0x17, 0xb9, 0x73, 0xa2, 0x52, 0x2b, 0x6f, 0x42, 0x82, 0xca, 0x26,
	0xdf, 0x8b, 0x77, 0x54, 0xa9, 0xf8, 0xe0, 0xba, 0x33, 0x4b, 0x2f, 0x9e, 0xea, 0x6d, 0x47, 0x61,
	0x28, 0xf5, 0xaf, 0xeb, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x10, 0x21, 0xf8, 0xe5, 0x00, 0x00,
	0x00,
}
