// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common // import "github.com/my0sot1s/mpb/common"

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

type ReqQuery struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Anchor               string   `protobuf:"bytes,3,opt,name=anchor,proto3" json:"anchor,omitempty"`
	OrderBy              string   `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	ReqId                *ReqId   `protobuf:"bytes,5,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqQuery) Reset()         { *m = ReqQuery{} }
func (m *ReqQuery) String() string { return proto.CompactTextString(m) }
func (*ReqQuery) ProtoMessage()    {}
func (*ReqQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c807eb9e1e84ab3e, []int{0}
}
func (m *ReqQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqQuery.Unmarshal(m, b)
}
func (m *ReqQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqQuery.Marshal(b, m, deterministic)
}
func (dst *ReqQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQuery.Merge(dst, src)
}
func (m *ReqQuery) XXX_Size() int {
	return xxx_messageInfo_ReqQuery.Size(m)
}
func (m *ReqQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQuery proto.InternalMessageInfo

func (m *ReqQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReqQuery) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReqQuery) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

func (m *ReqQuery) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ReqQuery) GetReqId() *ReqId {
	if m != nil {
		return m.ReqId
	}
	return nil
}

type ReqId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqId) Reset()         { *m = ReqId{} }
func (m *ReqId) String() string { return proto.CompactTextString(m) }
func (*ReqId) ProtoMessage()    {}
func (*ReqId) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c807eb9e1e84ab3e, []int{1}
}
func (m *ReqId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqId.Unmarshal(m, b)
}
func (m *ReqId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqId.Marshal(b, m, deterministic)
}
func (dst *ReqId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqId.Merge(dst, src)
}
func (m *ReqId) XXX_Size() int {
	return xxx_messageInfo_ReqId.Size(m)
}
func (m *ReqId) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqId.DiscardUnknown(m)
}

var xxx_messageInfo_ReqId proto.InternalMessageInfo

func (m *ReqId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ServiceReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	By                   string   `protobuf:"bytes,2,opt,name=by,proto3" json:"by,omitempty"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	UserId               []string `protobuf:"bytes,4,rep,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId              []string `protobuf:"bytes,5,rep,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceReq) Reset()         { *m = ServiceReq{} }
func (m *ServiceReq) String() string { return proto.CompactTextString(m) }
func (*ServiceReq) ProtoMessage()    {}
func (*ServiceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c807eb9e1e84ab3e, []int{2}
}
func (m *ServiceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceReq.Unmarshal(m, b)
}
func (m *ServiceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceReq.Marshal(b, m, deterministic)
}
func (dst *ServiceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceReq.Merge(dst, src)
}
func (m *ServiceReq) XXX_Size() int {
	return xxx_messageInfo_ServiceReq.Size(m)
}
func (m *ServiceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceReq.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceReq proto.InternalMessageInfo

func (m *ServiceReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceReq) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

func (m *ServiceReq) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *ServiceReq) GetUserId() []string {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *ServiceReq) GetGroupId() []string {
	if m != nil {
		return m.GroupId
	}
	return nil
}

type Nil struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nil) Reset()         { *m = Nil{} }
func (m *Nil) String() string { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()    {}
func (*Nil) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c807eb9e1e84ab3e, []int{3}
}
func (m *Nil) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nil.Unmarshal(m, b)
}
func (m *Nil) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nil.Marshal(b, m, deterministic)
}
func (dst *Nil) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nil.Merge(dst, src)
}
func (m *Nil) XXX_Size() int {
	return xxx_messageInfo_Nil.Size(m)
}
func (m *Nil) XXX_DiscardUnknown() {
	xxx_messageInfo_Nil.DiscardUnknown(m)
}

var xxx_messageInfo_Nil proto.InternalMessageInfo

type ErrorMsg struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorMsg) Reset()         { *m = ErrorMsg{} }
func (m *ErrorMsg) String() string { return proto.CompactTextString(m) }
func (*ErrorMsg) ProtoMessage()    {}
func (*ErrorMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c807eb9e1e84ab3e, []int{4}
}
func (m *ErrorMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorMsg.Unmarshal(m, b)
}
func (m *ErrorMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorMsg.Marshal(b, m, deterministic)
}
func (dst *ErrorMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMsg.Merge(dst, src)
}
func (m *ErrorMsg) XXX_Size() int {
	return xxx_messageInfo_ErrorMsg.Size(m)
}
func (m *ErrorMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMsg proto.InternalMessageInfo

func (m *ErrorMsg) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ErrorMsg) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type MediaInfo struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Width                int32    `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	SignatureKey         string   `protobuf:"bytes,5,opt,name=signature_key,json=signatureKey,proto3" json:"signature_key,omitempty"`
	PublicId             string   `protobuf:"bytes,6,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	Format               string   `protobuf:"bytes,7,opt,name=format,proto3" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MediaInfo) Reset()         { *m = MediaInfo{} }
func (m *MediaInfo) String() string { return proto.CompactTextString(m) }
func (*MediaInfo) ProtoMessage()    {}
func (*MediaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c807eb9e1e84ab3e, []int{5}
}
func (m *MediaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaInfo.Unmarshal(m, b)
}
func (m *MediaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaInfo.Marshal(b, m, deterministic)
}
func (dst *MediaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaInfo.Merge(dst, src)
}
func (m *MediaInfo) XXX_Size() int {
	return xxx_messageInfo_MediaInfo.Size(m)
}
func (m *MediaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MediaInfo proto.InternalMessageInfo

func (m *MediaInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *MediaInfo) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *MediaInfo) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *MediaInfo) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MediaInfo) GetSignatureKey() string {
	if m != nil {
		return m.SignatureKey
	}
	return ""
}

func (m *MediaInfo) GetPublicId() string {
	if m != nil {
		return m.PublicId
	}
	return ""
}

func (m *MediaInfo) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqQuery)(nil), "common.ReqQuery")
	proto.RegisterType((*ReqId)(nil), "common.ReqId")
	proto.RegisterType((*ServiceReq)(nil), "common.ServiceReq")
	proto.RegisterType((*Nil)(nil), "common.Nil")
	proto.RegisterType((*ErrorMsg)(nil), "common.ErrorMsg")
	proto.RegisterType((*MediaInfo)(nil), "common.MediaInfo")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_common_c807eb9e1e84ab3e) }

var fileDescriptor_common_c807eb9e1e84ab3e = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x8e, 0x94, 0x40,
	0x10, 0x80, 0x03, 0x33, 0xc0, 0x50, 0xeb, 0x6e, 0x4c, 0x6b, 0x5c, 0x8c, 0x89, 0x21, 0xe8, 0x61,
	0x4e, 0x3b, 0xfe, 0xbc, 0xc1, 0x24, 0x1e, 0x88, 0x59, 0x13, 0xdb, 0x9b, 0x97, 0x09, 0xd0, 0xb5,
	0xd0, 0x11, 0x68, 0xa6, 0x00, 0x4d, 0xef, 0x1b, 0xf8, 0x48, 0xbe, 0x9d, 0xe9, 0x1f, 0x8d, 0x71,
	0x4f, 0x5d, 0x5f, 0x75, 0xa7, 0xfb, 0xeb, 0xaa, 0x82, 0x27, 0x8d, 0x1a, 0x06, 0x35, 0x1e, 0xdc,
	0x72, 0x33, 0x91, 0x5a, 0x14, 0x8b, 0x1d, 0x15, 0x3f, 0x03, 0xd8, 0x71, 0x3c, 0x7f, 0x5e, 0x91,
	0x34, 0xbb, 0x82, 0x50, 0x8a, 0x2c, 0xc8, 0x83, 0x7d, 0xca, 0x43, 0x29, 0xd8, 0x53, 0x88, 0x7a,
	0x39, 0xc8, 0x25, 0x0b, 0xf3, 0x60, 0x1f, 0x71, 0x07, 0xec, 0x19, 0xc4, 0xd5, 0xd8, 0x74, 0x8a,
	0xb2, 0x8d, 0x3d, 0xe9, 0x89, 0x3d, 0x87, 0x9d, 0x22, 0x81, 0x74, 0xaa, 0x75, 0xb6, 0xb5, 0x3b,
	0x89, 0xe5, 0xa3, 0x66, 0xaf, 0x21, 0x26, 0x3c, 0x9f, 0xa4, 0xc8, 0xa2, 0x3c, 0xd8, 0x5f, 0xbc,
	0xbb, 0xbc, 0xf1, 0x32, 0x1c, 0xcf, 0xa5, 0xe0, 0x11, 0x99, 0xa5, 0xb8, 0x86, 0xc8, 0xf2, 0xff,
	0x1e, 0xc5, 0x3d, 0xc0, 0x17, 0xa4, 0xef, 0xb2, 0x41, 0x8e, 0xe7, 0x07, 0x96, 0x57, 0x10, 0xd6,
	0xda, 0x2a, 0xa6, 0x3c, 0xac, 0x35, 0xcb, 0x20, 0x99, 0x2a, 0xdd, 0xab, 0x4a, 0x78, 0xc1, 0x3f,
	0xc8, 0xae, 0x21, 0x59, 0x67, 0x24, 0xe3, 0xb1, 0xcd, 0x37, 0x46, 0xdd, 0x60, 0x29, 0x8c, 0x7a,
	0x4b, 0x6a, 0x9d, 0x9c, 0xa1, 0xd9, 0x49, 0x2c, 0x97, 0xa2, 0x88, 0x60, 0xf3, 0x49, 0xf6, 0xc5,
	0x11, 0x76, 0x1f, 0x88, 0x14, 0xdd, 0xce, 0xad, 0x29, 0x0b, 0x9a, 0xd8, 0x3b, 0x38, 0x60, 0x39,
	0x5c, 0x08, 0x9c, 0x1b, 0x92, 0xd3, 0x22, 0xd5, 0xe8, 0x7d, 0xfe, 0x4d, 0x15, 0xbf, 0x02, 0x48,
	0x6f, 0x51, 0xc8, 0xaa, 0x1c, 0xef, 0x14, 0x7b, 0x0c, 0x9b, 0x95, 0x7a, 0x7f, 0x87, 0x09, 0x19,
	0x83, 0xed, 0x2c, 0xef, 0xd1, 0x57, 0xdb, 0xc6, 0xe6, 0xad, 0x1f, 0x52, 0x2c, 0x9d, 0xfd, 0x4a,
	0xc4, 0x1d, 0x98, 0x16, 0x74, 0x28, 0xdb, 0x6e, 0xb1, 0x85, 0x8e, 0xb8, 0x27, 0xf6, 0x0a, 0x2e,
	0x67, 0xd9, 0x8e, 0xd5, 0xb2, 0x12, 0x9e, 0xbe, 0xa1, 0xb6, 0xe5, 0x4e, 0xf9, 0xa3, 0xbf, 0xc9,
	0x8f, 0xa8, 0xd9, 0x0b, 0x48, 0xa7, 0xb5, 0xee, 0x65, 0x63, 0x7e, 0x1b, 0xdb, 0x03, 0x3b, 0x97,
	0x28, 0x85, 0xb9, 0xf9, 0x4e, 0xd1, 0x50, 0x2d, 0x59, 0xe2, 0x9a, 0xeb, 0xe8, 0x98, 0x7f, 0x7d,
	0xd9, 0xca, 0xa5, 0x5b, 0x6b, 0xd3, 0xb9, 0xc3, 0xa0, 0xdf, 0xcc, 0x6a, 0x79, 0x3b, 0x1f, 0x86,
	0xa9, 0xf6, 0x73, 0x55, 0xc7, 0x76, 0xb0, 0xde, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x40,
	0x8b, 0xce, 0x6f, 0x02, 0x00, 0x00,
}
