// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/my0sot1s/mpb/common"

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

type UserType int32

const (
	UserType_NormalUser UserType = 0
	UserType_SuperUser  UserType = 1
)

var UserType_name = map[int32]string{
	0: "NormalUser",
	1: "SuperUser",
}
var UserType_value = map[string]int32{
	"NormalUser": 0,
	"SuperUser":  1,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}
func (UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_3b9320c540399847, []int{0}
}

type EmailSubscribers struct {
	SubList              []*EmailSubscriber `protobuf:"bytes,1,rep,name=sub_list,json=subList,proto3" json:"sub_list,omitempty"`
	SendDone             []bool             `protobuf:"varint,2,rep,packed,name=send_done,json=sendDone,proto3" json:"send_done,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EmailSubscribers) Reset()         { *m = EmailSubscribers{} }
func (m *EmailSubscribers) String() string { return proto.CompactTextString(m) }
func (*EmailSubscribers) ProtoMessage()    {}
func (*EmailSubscribers) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3b9320c540399847, []int{0}
}
func (m *EmailSubscribers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailSubscribers.Unmarshal(m, b)
}
func (m *EmailSubscribers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailSubscribers.Marshal(b, m, deterministic)
}
func (dst *EmailSubscribers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailSubscribers.Merge(dst, src)
}
func (m *EmailSubscribers) XXX_Size() int {
	return xxx_messageInfo_EmailSubscribers.Size(m)
}
func (m *EmailSubscribers) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailSubscribers.DiscardUnknown(m)
}

var xxx_messageInfo_EmailSubscribers proto.InternalMessageInfo

func (m *EmailSubscribers) GetSubList() []*EmailSubscriber {
	if m != nil {
		return m.SubList
	}
	return nil
}

func (m *EmailSubscribers) GetSendDone() []bool {
	if m != nil {
		return m.SendDone
	}
	return nil
}

type EmailSubscriber struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailSubscriber) Reset()         { *m = EmailSubscriber{} }
func (m *EmailSubscriber) String() string { return proto.CompactTextString(m) }
func (*EmailSubscriber) ProtoMessage()    {}
func (*EmailSubscriber) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3b9320c540399847, []int{1}
}
func (m *EmailSubscriber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailSubscriber.Unmarshal(m, b)
}
func (m *EmailSubscriber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailSubscriber.Marshal(b, m, deterministic)
}
func (dst *EmailSubscriber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailSubscriber.Merge(dst, src)
}
func (m *EmailSubscriber) XXX_Size() int {
	return xxx_messageInfo_EmailSubscriber.Size(m)
}
func (m *EmailSubscriber) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailSubscriber.DiscardUnknown(m)
}

var xxx_messageInfo_EmailSubscriber proto.InternalMessageInfo

func (m *EmailSubscriber) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EmailSubscriber) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *EmailSubscriber) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailSubscriber) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type SocialDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IdInfo               string   `protobuf:"bytes,3,opt,name=id_info,json=idInfo,proto3" json:"id_info,omitempty"`
	Token                *Token   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocialDetail) Reset()         { *m = SocialDetail{} }
func (m *SocialDetail) String() string { return proto.CompactTextString(m) }
func (*SocialDetail) ProtoMessage()    {}
func (*SocialDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3b9320c540399847, []int{2}
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

func (m *SocialDetail) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *SocialDetail) GetIdInfo() string {
	if m != nil {
		return m.IdInfo
	}
	return ""
}

func (m *SocialDetail) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Anchor               string   `protobuf:"bytes,2,opt,name=anchor,proto3" json:"anchor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3b9320c540399847, []int{3}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Users) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

type User struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string          `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Phones               []string        `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	Email                string          `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Websites             []string        `protobuf:"bytes,6,rep,name=websites,proto3" json:"websites,omitempty"`
	Socials              []*SocialDetail `protobuf:"bytes,7,rep,name=socials,proto3" json:"socials,omitempty"`
	State                string          `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
	GroupIds             []string        `protobuf:"bytes,9,rep,name=group_ids,json=groupIds,proto3" json:"group_ids,omitempty"`
	Groups               []*GroupUser    `protobuf:"bytes,10,rep,name=groups,proto3" json:"groups,omitempty"`
	Notes                []string        `protobuf:"bytes,11,rep,name=notes,proto3" json:"notes,omitempty"`
	Created              int32           `protobuf:"varint,12,opt,name=created,proto3" json:"created,omitempty"`
	Avatar               string          `protobuf:"bytes,13,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3b9320c540399847, []int{4}
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

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
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

func (m *User) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *User) GetGroupIds() []string {
	if m != nil {
		return m.GroupIds
	}
	return nil
}

func (m *User) GetGroups() []*GroupUser {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *User) GetNotes() []string {
	if m != nil {
		return m.Notes
	}
	return nil
}

func (m *User) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type GroupUser struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PermIds              []string      `protobuf:"bytes,3,rep,name=perm_ids,json=permIds,proto3" json:"perm_ids,omitempty"`
	Perms                []*Permission `protobuf:"bytes,4,rep,name=perms,proto3" json:"perms,omitempty"`
	State                string        `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Created              int32         `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GroupUser) Reset()         { *m = GroupUser{} }
func (m *GroupUser) String() string { return proto.CompactTextString(m) }
func (*GroupUser) ProtoMessage()    {}
func (*GroupUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3b9320c540399847, []int{5}
}
func (m *GroupUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupUser.Unmarshal(m, b)
}
func (m *GroupUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupUser.Marshal(b, m, deterministic)
}
func (dst *GroupUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupUser.Merge(dst, src)
}
func (m *GroupUser) XXX_Size() int {
	return xxx_messageInfo_GroupUser.Size(m)
}
func (m *GroupUser) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupUser.DiscardUnknown(m)
}

var xxx_messageInfo_GroupUser proto.InternalMessageInfo

func (m *GroupUser) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GroupUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GroupUser) GetPermIds() []string {
	if m != nil {
		return m.PermIds
	}
	return nil
}

func (m *GroupUser) GetPerms() []*Permission {
	if m != nil {
		return m.Perms
	}
	return nil
}

func (m *GroupUser) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *GroupUser) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

type Permission struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Created              int32    `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3b9320c540399847, []int{6}
}
func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (dst *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(dst, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Permission) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Permission) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Permission) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Permission) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

type Token struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Created              int32    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Modified             int32    `protobuf:"varint,5,opt,name=modified,proto3" json:"modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_3b9320c540399847, []int{7}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Token) GetModified() int32 {
	if m != nil {
		return m.Modified
	}
	return 0
}

func init() {
	proto.RegisterType((*EmailSubscribers)(nil), "user.EmailSubscribers")
	proto.RegisterType((*EmailSubscriber)(nil), "user.EmailSubscriber")
	proto.RegisterType((*SocialDetail)(nil), "user.SocialDetail")
	proto.RegisterType((*Users)(nil), "user.Users")
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*GroupUser)(nil), "user.GroupUser")
	proto.RegisterType((*Permission)(nil), "user.Permission")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterEnum("user.UserType", UserType_name, UserType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserSubStreamClient is the client API for UserSubStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSubStreamClient interface {
	// email
	DoSendEmail(ctx context.Context, in *EmailSubscribers, opts ...grpc.CallOption) (*EmailSubscribers, error)
}

type userSubStreamClient struct {
	cc *grpc.ClientConn
}

func NewUserSubStreamClient(cc *grpc.ClientConn) UserSubStreamClient {
	return &userSubStreamClient{cc}
}

func (c *userSubStreamClient) DoSendEmail(ctx context.Context, in *EmailSubscribers, opts ...grpc.CallOption) (*EmailSubscribers, error) {
	out := new(EmailSubscribers)
	err := c.cc.Invoke(ctx, "/user.UserSubStream/DoSendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSubStreamServer is the server API for UserSubStream service.
type UserSubStreamServer interface {
	// email
	DoSendEmail(context.Context, *EmailSubscribers) (*EmailSubscribers, error)
}

func RegisterUserSubStreamServer(s *grpc.Server, srv UserSubStreamServer) {
	s.RegisterService(&_UserSubStream_serviceDesc, srv)
}

func _UserSubStream_DoSendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailSubscribers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSubStreamServer).DoSendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSubStream/DoSendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSubStreamServer).DoSendEmail(ctx, req.(*EmailSubscribers))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSubStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserSubStream",
	HandlerType: (*UserSubStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSendEmail",
			Handler:    _UserSubStream_DoSendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

// UserStreamClient is the client API for UserStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserStreamClient interface {
	// user
	LoginOauth(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error)
	RefreshTokenOauth(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	ListUsers(ctx context.Context, in *common.ReqQuery, opts ...grpc.CallOption) (*Users, error)
	UserById(ctx context.Context, in *common.ReqId, opts ...grpc.CallOption) (*User, error)
}

type userStreamClient struct {
	cc *grpc.ClientConn
}

func NewUserStreamClient(cc *grpc.ClientConn) UserStreamClient {
	return &userStreamClient{cc}
}

func (c *userStreamClient) LoginOauth(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/user.UserStream/LoginOauth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStreamClient) RefreshTokenOauth(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/user.UserStream/RefreshTokenOauth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStreamClient) ListUsers(ctx context.Context, in *common.ReqQuery, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user.UserStream/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStreamClient) UserById(ctx context.Context, in *common.ReqId, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserStream/UserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserStreamServer is the server API for UserStream service.
type UserStreamServer interface {
	// user
	LoginOauth(context.Context, *User) (*Token, error)
	RefreshTokenOauth(context.Context, *Token) (*Token, error)
	ListUsers(context.Context, *common.ReqQuery) (*Users, error)
	UserById(context.Context, *common.ReqId) (*User, error)
}

func RegisterUserStreamServer(s *grpc.Server, srv UserStreamServer) {
	s.RegisterService(&_UserStream_serviceDesc, srv)
}

func _UserStream_LoginOauth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStreamServer).LoginOauth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserStream/LoginOauth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStreamServer).LoginOauth(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserStream_RefreshTokenOauth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStreamServer).RefreshTokenOauth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserStream/RefreshTokenOauth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStreamServer).RefreshTokenOauth(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserStream_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.ReqQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStreamServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserStream/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStreamServer).ListUsers(ctx, req.(*common.ReqQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserStream_UserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.ReqId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStreamServer).UserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserStream/UserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStreamServer).UserById(ctx, req.(*common.ReqId))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserStream",
	HandlerType: (*UserStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginOauth",
			Handler:    _UserStream_LoginOauth_Handler,
		},
		{
			MethodName: "RefreshTokenOauth",
			Handler:    _UserStream_RefreshTokenOauth_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserStream_ListUsers_Handler,
		},
		{
			MethodName: "UserById",
			Handler:    _UserStream_UserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_user_3b9320c540399847) }

var fileDescriptor_user_3b9320c540399847 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x6f, 0xda, 0x48,
	0x10, 0x3e, 0x03, 0x06, 0x33, 0x86, 0x84, 0xdb, 0xcb, 0xe5, 0x7c, 0xdc, 0x0b, 0xe1, 0x94, 0x96,
	0x46, 0x6d, 0x52, 0xd1, 0xe7, 0x3e, 0xb4, 0x4a, 0x54, 0x21, 0x45, 0x69, 0x6b, 0xd2, 0x67, 0x64,
	0xd8, 0x25, 0xac, 0x8a, 0x77, 0xdd, 0x5d, 0xbb, 0x15, 0xaf, 0x7d, 0xca, 0xef, 0xe8, 0x7f, 0xe8,
	0x8f, 0xea, 0xbf, 0xa8, 0x66, 0xd7, 0xc0, 0x92, 0x34, 0x52, 0x5f, 0x60, 0xbf, 0x99, 0xf1, 0x37,
	0xdf, 0x7c, 0x6b, 0x0f, 0xec, 0x17, 0x9a, 0xa9, 0x33, 0xfc, 0x39, 0xcd, 0x94, 0xcc, 0x25, 0xa9,
	0xe1, 0xb9, 0xfb, 0xd7, 0x4c, 0xa6, 0xa9, 0x14, 0x67, 0xf6, 0xcf, 0xa6, 0xfa, 0x09, 0x74, 0x2e,
	0xd2, 0x84, 0x2f, 0xc7, 0xc5, 0x54, 0xcf, 0x14, 0x9f, 0x32, 0xa5, 0xc9, 0x73, 0x08, 0x74, 0x31,
	0x9d, 0x2c, 0xb9, 0xce, 0x23, 0xaf, 0x57, 0x1d, 0x84, 0xc3, 0xbf, 0x4f, 0x0d, 0xdb, 0x9d, 0xca,
	0xb8, 0xa1, 0x8b, 0xe9, 0x25, 0xd7, 0x39, 0xf9, 0x0f, 0x9a, 0x9a, 0x09, 0x3a, 0xa1, 0x52, 0xb0,
	0xa8, 0xd2, 0xab, 0x0e, 0x82, 0x38, 0xc0, 0xc0, 0xb9, 0x14, 0xac, 0x9f, 0xc0, 0xfe, 0x9d, 0x07,
	0xc9, 0x1e, 0x54, 0x38, 0x8d, 0xbc, 0x9e, 0x37, 0x68, 0xc6, 0x15, 0x4e, 0x49, 0x07, 0xaa, 0x05,
	0xa7, 0x51, 0xc5, 0x04, 0xf0, 0x48, 0x0e, 0xc0, 0x67, 0xf8, 0x50, 0x54, 0x35, 0x31, 0x0b, 0xc8,
	0x21, 0xd4, 0x75, 0x9e, 0xe4, 0x85, 0x8e, 0x6a, 0x26, 0x5c, 0xa2, 0xfe, 0xad, 0x07, 0xad, 0xb1,
	0x9c, 0xf1, 0x64, 0x79, 0xce, 0x72, 0x2c, 0x24, 0x50, 0x13, 0x49, 0xca, 0xca, 0x16, 0xe6, 0x6c,
	0x9a, 0xa8, 0xe5, 0xa6, 0x89, 0x32, 0x74, 0xc9, 0xe7, 0x24, 0x4f, 0xd4, 0x9a, 0xce, 0x22, 0xf2,
	0x0f, 0x34, 0x38, 0x9d, 0x70, 0x31, 0x97, 0x65, 0xfb, 0x3a, 0xa7, 0x23, 0x31, 0x97, 0xe4, 0x08,
	0xfc, 0x5c, 0x7e, 0x64, 0x22, 0xf2, 0x7b, 0xde, 0x20, 0x1c, 0x86, 0xd6, 0x96, 0x6b, 0x0c, 0xc5,
	0x36, 0xd3, 0x7f, 0x05, 0xfe, 0x07, 0x8d, 0x2e, 0xf6, 0xc0, 0xc7, 0xac, 0x2e, 0x2d, 0x04, 0x5b,
	0x8b, 0xb9, 0xd8, 0x26, 0x4c, 0x7b, 0x31, 0x5b, 0x48, 0x55, 0x6a, 0x2a, 0x51, 0xff, 0x47, 0x05,
	0x6a, 0x58, 0x77, 0xcf, 0xa6, 0xf5, 0x54, 0x15, 0x67, 0xaa, 0x23, 0x68, 0x51, 0xae, 0xb3, 0x65,
	0xb2, 0x9a, 0x98, 0x9c, 0x15, 0x1c, 0x96, 0xb1, 0x2b, 0x2c, 0x39, 0x84, 0x7a, 0xb6, 0x90, 0x82,
	0xa1, 0x6b, 0x55, 0xec, 0x63, 0xd1, 0xd6, 0x63, 0xdf, 0xf5, 0xb8, 0x0b, 0xc1, 0x17, 0x36, 0xd5,
	0x3c, 0x67, 0x3a, 0xaa, 0x9b, 0xfa, 0x0d, 0x26, 0x4f, 0xa1, 0xa1, 0x8d, 0xcd, 0x3a, 0x6a, 0x98,
	0xa9, 0x88, 0x9d, 0xca, 0xf5, 0x3e, 0x5e, 0x97, 0x20, 0x3f, 0xde, 0x0f, 0x8b, 0x02, 0xcb, 0x6f,
	0x00, 0xbe, 0x2b, 0x37, 0x4a, 0x16, 0xd9, 0x84, 0x53, 0x1d, 0x35, 0x6d, 0x03, 0x13, 0x18, 0x51,
	0x4d, 0x1e, 0x43, 0xdd, 0x9c, 0x75, 0x04, 0x86, 0x7f, 0xdf, 0xf2, 0xbf, 0xc1, 0x98, 0xb1, 0xae,
	0x4c, 0x23, 0xb7, 0x90, 0x28, 0x31, 0x34, 0x0c, 0x16, 0x90, 0x08, 0x1a, 0x33, 0xc5, 0x92, 0x9c,
	0xd1, 0xa8, 0xd5, 0xf3, 0x06, 0x7e, 0xbc, 0x86, 0xce, 0x55, 0xb7, 0xdd, 0xab, 0xee, 0x7f, 0xf3,
	0xa0, 0xb9, 0x61, 0xff, 0x2d, 0xc3, 0xff, 0x85, 0x20, 0x63, 0x2a, 0x35, 0xf2, 0xab, 0xa6, 0x79,
	0x03, 0x31, 0xaa, 0x7f, 0x04, 0x3e, 0x1e, 0xad, 0xcf, 0xe1, 0xb0, 0x63, 0xc5, 0xbf, 0x63, 0x2a,
	0xe5, 0x5a, 0x73, 0x29, 0x62, 0x9b, 0xde, 0x1a, 0xe3, 0xbb, 0xc6, 0x38, 0xe2, 0xeb, 0x3b, 0xe2,
	0xfb, 0x5f, 0x3d, 0x80, 0x2d, 0xcb, 0x3d, 0x95, 0x3d, 0x08, 0x29, 0xc3, 0x4f, 0x2b, 0xcb, 0xb9,
	0x14, 0xa5, 0x58, 0x37, 0x84, 0x73, 0xe4, 0xab, 0x6c, 0xfd, 0x72, 0x98, 0xf3, 0x56, 0x44, 0xed,
	0x01, 0x11, 0xfe, 0xae, 0x88, 0x5b, 0x0f, 0x7c, 0xf3, 0xa6, 0xdf, 0xeb, 0x7f, 0xb0, 0xfe, 0x2a,
	0x6c, 0x67, 0x0b, 0xc8, 0xff, 0xd0, 0x56, 0x6c, 0xae, 0x98, 0x5e, 0x4c, 0x6c, 0xd6, 0x36, 0x6f,
	0x95, 0x41, 0x4b, 0xe5, 0xb4, 0xab, 0xed, 0x5e, 0x58, 0x17, 0x82, 0x54, 0x52, 0x3e, 0xe7, 0x1b,
	0x25, 0x1b, 0x7c, 0xf2, 0x04, 0x02, 0xbc, 0xae, 0x6b, 0x1c, 0x63, 0x0f, 0xe0, 0x4a, 0xaa, 0x34,
	0x59, 0x62, 0xa4, 0xf3, 0x07, 0x69, 0x43, 0x73, 0x5c, 0x64, 0x4c, 0x19, 0xe8, 0x0d, 0xaf, 0xa0,
	0x8d, 0xa7, 0x71, 0x31, 0x1d, 0xe7, 0x8a, 0x25, 0x29, 0x79, 0x09, 0xe1, 0xb9, 0x1c, 0x33, 0x41,
	0x2f, 0xec, 0x46, 0xf9, 0xe5, 0x66, 0xd3, 0xdd, 0x07, 0xe2, 0xc3, 0xef, 0x1e, 0x80, 0x21, 0xb4,
	0x6c, 0xc7, 0x00, 0x97, 0xf2, 0x86, 0x8b, 0xb7, 0x49, 0x91, 0x2f, 0x88, 0xf3, 0x8d, 0x77, 0xdd,
	0xdd, 0x40, 0x9e, 0xc1, 0x9f, 0xb1, 0x33, 0xb6, 0xad, 0x76, 0x2b, 0x76, 0xcb, 0x4f, 0xa0, 0x89,
	0x6b, 0xd5, 0xee, 0x91, 0xce, 0x69, 0xb9, 0xb0, 0x63, 0xf6, 0xe9, 0x7d, 0xc1, 0xd4, 0x6a, 0x5d,
	0x6b, 0xd3, 0xc7, 0xd6, 0x8b, 0xd7, 0xab, 0x11, 0x25, 0x6d, 0xa7, 0x74, 0x44, 0xbb, 0x8e, 0x9c,
	0x69, 0xdd, 0xac, 0xfb, 0x17, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xd5, 0xcc, 0xe1, 0x1c,
	0x06, 0x00, 0x00,
}
