// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package mpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User_State int32

const (
	User_ACTIVE   User_State = 0
	User_INACTIVE User_State = 1
	User_BAN      User_State = 2
)

var User_State_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
	2: "BAN",
}
var User_State_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
	"BAN":      2,
}

func (x User_State) String() string {
	return proto.EnumName(User_State_name, int32(x))
}
func (User_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_22c0f9194e11d042, []int{1, 0}
}

type SocialDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	IdInfo               string   `protobuf:"bytes,3,opt,name=id_info,json=idInfo,proto3" json:"id_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocialDetail) Reset()         { *m = SocialDetail{} }
func (m *SocialDetail) String() string { return proto.CompactTextString(m) }
func (*SocialDetail) ProtoMessage()    {}
func (*SocialDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_22c0f9194e11d042, []int{0}
}
func (m *SocialDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocialDetail.Unmarshal(m, b)
}
func (m *SocialDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocialDetail.Marshal(b, m, deterministic)
}
func (dst *SocialDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocialDetail.Merge(dst, src)
}
func (m *SocialDetail) XXX_Size() int {
	return xxx_messageInfo_SocialDetail.Size(m)
}
func (m *SocialDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_SocialDetail.DiscardUnknown(m)
}

var xxx_messageInfo_SocialDetail proto.InternalMessageInfo

func (m *SocialDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SocialDetail) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SocialDetail) GetIdInfo() string {
	if m != nil {
		return m.IdInfo
	}
	return ""
}

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string               `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Phones               []string             `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	Emails               []string             `protobuf:"bytes,5,rep,name=emails,proto3" json:"emails,omitempty"`
	Websites             []string             `protobuf:"bytes,6,rep,name=websites,proto3" json:"websites,omitempty"`
	Socials              []*SocialDetail      `protobuf:"bytes,7,rep,name=socials,proto3" json:"socials,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
	State                User_State           `protobuf:"varint,9,opt,name=state,proto3,enum=mpb.User_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_22c0f9194e11d042, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetPhones() []string {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *User) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *User) GetWebsites() []string {
	if m != nil {
		return m.Websites
	}
	return nil
}

func (m *User) GetSocials() []*SocialDetail {
	if m != nil {
		return m.Socials
	}
	return nil
}

func (m *User) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *User) GetState() User_State {
	if m != nil {
		return m.State
	}
	return User_ACTIVE
}

func init() {
	proto.RegisterType((*SocialDetail)(nil), "mpb.SocialDetail")
	proto.RegisterType((*User)(nil), "mpb.User")
	proto.RegisterEnum("mpb.User_State", User_State_name, User_State_value)
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_22c0f9194e11d042) }

var fileDescriptor_user_22c0f9194e11d042 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x6d, 0x0b, 0x2d, 0x0c, 0x04, 0x71, 0x0f, 0xba, 0xe1, 0x62, 0x25, 0x31, 0x69, 0x34,
	0x59, 0x12, 0xf4, 0x0b, 0xe0, 0x9f, 0x03, 0x07, 0x39, 0x14, 0xf4, 0x4a, 0xb6, 0x74, 0xc0, 0x4d,
	0xda, 0x6e, 0xd3, 0x5d, 0x62, 0xfc, 0x94, 0x7e, 0x25, 0xd3, 0xdd, 0x96, 0x70, 0x9b, 0xf7, 0x9b,
	0x97, 0x37, 0x99, 0x07, 0x70, 0x54, 0x58, 0xb1, 0xb2, 0x92, 0x5a, 0x12, 0x2f, 0x2f, 0x93, 0xc9,
	0xed, 0x41, 0xca, 0x43, 0x86, 0x33, 0x83, 0x92, 0xe3, 0x7e, 0xa6, 0x45, 0x8e, 0x4a, 0xf3, 0xbc,
	0xb4, 0xae, 0xe9, 0x07, 0x0c, 0xd7, 0x72, 0x27, 0x78, 0xf6, 0x86, 0x9a, 0x8b, 0x8c, 0x10, 0xe8,
	0x14, 0x3c, 0x47, 0xea, 0x84, 0x4e, 0xd4, 0x8f, 0xcd, 0x4c, 0xc6, 0xe0, 0x1d, 0xab, 0x8c, 0xba,
	0x06, 0xd5, 0x23, 0xb9, 0x81, 0x40, 0xa4, 0x5b, 0x51, 0xec, 0x25, 0xf5, 0x0c, 0xf5, 0x45, 0xba,
	0x2c, 0xf6, 0x72, 0xfa, 0xe7, 0x42, 0xe7, 0x53, 0x61, 0x45, 0x46, 0xe0, 0x8a, 0xb4, 0x49, 0x71,
	0x45, 0x7a, 0xca, 0x75, 0xcf, 0x72, 0xef, 0x60, 0x98, 0x0a, 0x55, 0x66, 0xfc, 0x77, 0x6b, 0x76,
	0x36, 0x6a, 0xd0, 0xb0, 0x55, 0x6d, 0xb9, 0x06, 0xbf, 0xfc, 0x96, 0x05, 0x2a, 0xda, 0x09, 0xbd,
	0xfa, 0x8e, 0x55, 0x35, 0xc7, 0x9c, 0x8b, 0x4c, 0xd1, 0xae, 0xe5, 0x56, 0x91, 0x09, 0xf4, 0x7e,
	0x30, 0x51, 0x42, 0xa3, 0xa2, 0xbe, 0xd9, 0x9c, 0x34, 0x79, 0x84, 0x40, 0x99, 0x57, 0x15, 0x0d,
	0x42, 0x2f, 0x1a, 0xcc, 0xaf, 0x58, 0x5e, 0x26, 0xec, 0xfc, 0xfd, 0xb8, 0x75, 0x90, 0x67, 0x08,
	0x76, 0x15, 0x72, 0x8d, 0x29, 0xed, 0x85, 0x4e, 0x34, 0x98, 0x4f, 0x98, 0xad, 0x92, 0xb5, 0x55,
	0xb2, 0x4d, 0x5b, 0x65, 0xdc, 0x5a, 0xc9, 0x3d, 0x74, 0x95, 0xe6, 0x1a, 0x69, 0x3f, 0x74, 0xa2,
	0xd1, 0xfc, 0xd2, 0x1c, 0xa8, 0xfb, 0x60, 0xeb, 0x1a, 0xc7, 0x76, 0x3b, 0x7d, 0x80, 0xae, 0xd1,
	0x04, 0xc0, 0x5f, 0xbc, 0x6e, 0x96, 0x5f, 0xef, 0xe3, 0x0b, 0x32, 0x84, 0xde, 0x72, 0xd5, 0x28,
	0x87, 0x04, 0xe0, 0xbd, 0x2c, 0x56, 0x63, 0x37, 0xf1, 0xcd, 0xbd, 0xa7, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0x3c, 0xc2, 0x5f, 0xdb, 0x01, 0x00, 0x00,
}
