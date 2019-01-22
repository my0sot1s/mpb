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

type OauthMessage_OauthType int32

const (
	OauthMessage_facebook OauthMessage_OauthType = 0
	OauthMessage_google   OauthMessage_OauthType = 1
	OauthMessage_self     OauthMessage_OauthType = 2
)

var OauthMessage_OauthType_name = map[int32]string{
	0: "facebook",
	1: "google",
	2: "self",
}
var OauthMessage_OauthType_value = map[string]int32{
	"facebook": 0,
	"google":   1,
	"self":     2,
}

func (x OauthMessage_OauthType) String() string {
	return proto.EnumName(OauthMessage_OauthType_name, int32(x))
}
func (OauthMessage_OauthType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{0, 0}
}

type User_UserState int32

const (
	User_normal User_UserState = 0
	User_super  User_UserState = 1
)

var User_UserState_name = map[int32]string{
	0: "normal",
	1: "super",
}
var User_UserState_value = map[string]int32{
	"normal": 0,
	"super":  1,
}

func (x User_UserState) String() string {
	return proto.EnumName(User_UserState_name, int32(x))
}
func (User_UserState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{6, 0}
}

type GroupUser_UserState int32

const (
	GroupUser_normal  GroupUser_UserState = 0
	GroupUser_manager GroupUser_UserState = 1
)

var GroupUser_UserState_name = map[int32]string{
	0: "normal",
	1: "manager",
}
var GroupUser_UserState_value = map[string]int32{
	"normal":  0,
	"manager": 1,
}

func (x GroupUser_UserState) String() string {
	return proto.EnumName(GroupUser_UserState_name, int32(x))
}
func (GroupUser_UserState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{7, 0}
}

type OauthMessage struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Redirect             string   `protobuf:"bytes,3,opt,name=redirect,proto3" json:"redirect,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OauthMessage) Reset()         { *m = OauthMessage{} }
func (m *OauthMessage) String() string { return proto.CompactTextString(m) }
func (*OauthMessage) ProtoMessage()    {}
func (*OauthMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{0}
}
func (m *OauthMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OauthMessage.Unmarshal(m, b)
}
func (m *OauthMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OauthMessage.Marshal(b, m, deterministic)
}
func (dst *OauthMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OauthMessage.Merge(dst, src)
}
func (m *OauthMessage) XXX_Size() int {
	return xxx_messageInfo_OauthMessage.Size(m)
}
func (m *OauthMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OauthMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OauthMessage proto.InternalMessageInfo

func (m *OauthMessage) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *OauthMessage) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *OauthMessage) GetRedirect() string {
	if m != nil {
		return m.Redirect
	}
	return ""
}

func (m *OauthMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
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
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{1}
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
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{2}
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

type Token struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Modified             int64    `protobuf:"varint,5,opt,name=modified,proto3" json:"modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{3}
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

func (m *Token) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Token) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
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
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{4}
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
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{5}
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
	Created              int64           `protobuf:"varint,12,opt,name=created,proto3" json:"created,omitempty"`
	Avatar               string          `protobuf:"bytes,13,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{6}
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

func (m *User) GetCreated() int64 {
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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State                string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Created              int64    `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupUser) Reset()         { *m = GroupUser{} }
func (m *GroupUser) String() string { return proto.CompactTextString(m) }
func (*GroupUser) ProtoMessage()    {}
func (*GroupUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_8f697d8a0b3b7bd1, []int{7}
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

func (m *GroupUser) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *GroupUser) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func init() {
	proto.RegisterType((*OauthMessage)(nil), "user.OauthMessage")
	proto.RegisterType((*EmailSubscribers)(nil), "user.EmailSubscribers")
	proto.RegisterType((*EmailSubscriber)(nil), "user.EmailSubscriber")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*SocialDetail)(nil), "user.SocialDetail")
	proto.RegisterType((*Users)(nil), "user.Users")
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*GroupUser)(nil), "user.GroupUser")
	proto.RegisterEnum("user.OauthMessage_OauthType", OauthMessage_OauthType_name, OauthMessage_OauthType_value)
	proto.RegisterEnum("user.User_UserState", User_UserState_name, User_UserState_value)
	proto.RegisterEnum("user.GroupUser_UserState", GroupUser_UserState_name, GroupUser_UserState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserStreamClient is the client API for UserStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserStreamClient interface {
	// user
	LoginOauth(ctx context.Context, in *OauthMessage, opts ...grpc.CallOption) (*Token, error)
	RefreshTokenOauth(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	ListUsers(ctx context.Context, in *common.ReqQuery, opts ...grpc.CallOption) (*Users, error)
	UserById(ctx context.Context, in *common.ReqId, opts ...grpc.CallOption) (*User, error)
	// follow
	DoSendEmail(ctx context.Context, in *EmailSubscribers, opts ...grpc.CallOption) (*EmailSubscribers, error)
}

type userStreamClient struct {
	cc *grpc.ClientConn
}

func NewUserStreamClient(cc *grpc.ClientConn) UserStreamClient {
	return &userStreamClient{cc}
}

func (c *userStreamClient) LoginOauth(ctx context.Context, in *OauthMessage, opts ...grpc.CallOption) (*Token, error) {
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

func (c *userStreamClient) DoSendEmail(ctx context.Context, in *EmailSubscribers, opts ...grpc.CallOption) (*EmailSubscribers, error) {
	out := new(EmailSubscribers)
	err := c.cc.Invoke(ctx, "/user.UserStream/DoSendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserStreamServer is the server API for UserStream service.
type UserStreamServer interface {
	// user
	LoginOauth(context.Context, *OauthMessage) (*Token, error)
	RefreshTokenOauth(context.Context, *Token) (*Token, error)
	ListUsers(context.Context, *common.ReqQuery) (*Users, error)
	UserById(context.Context, *common.ReqId) (*User, error)
	// follow
	DoSendEmail(context.Context, *EmailSubscribers) (*EmailSubscribers, error)
}

func RegisterUserStreamServer(s *grpc.Server, srv UserStreamServer) {
	s.RegisterService(&_UserStream_serviceDesc, srv)
}

func _UserStream_LoginOauth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OauthMessage)
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
		return srv.(UserStreamServer).LoginOauth(ctx, req.(*OauthMessage))
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

func _UserStream_DoSendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailSubscribers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStreamServer).DoSendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserStream/DoSendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStreamServer).DoSendEmail(ctx, req.(*EmailSubscribers))
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
		{
			MethodName: "DoSendEmail",
			Handler:    _UserStream_DoSendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_user_8f697d8a0b3b7bd1) }

var fileDescriptor_user_8f697d8a0b3b7bd1 = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6f, 0xeb, 0x44,
	0x10, 0x7e, 0x4e, 0x62, 0xc7, 0x9e, 0xa4, 0xd4, 0x2c, 0x8f, 0x87, 0x15, 0x2e, 0x79, 0x06, 0x44,
	0x85, 0x78, 0x2d, 0x2a, 0x67, 0x0e, 0xa0, 0x3e, 0xa1, 0x48, 0x0f, 0x10, 0x4e, 0x39, 0x47, 0x1b,
	0xef, 0x24, 0x59, 0xd5, 0xf6, 0x86, 0x5d, 0x1b, 0x94, 0x3f, 0x00, 0x89, 0x0b, 0x67, 0xfe, 0x59,
	0x0e, 0x68, 0x76, 0x37, 0xa9, 0xdb, 0x52, 0x89, 0x4b, 0x3b, 0x3f, 0x3e, 0xcf, 0x7c, 0xdf, 0xce,
	0x64, 0xe0, 0xbc, 0x33, 0xa8, 0xaf, 0xe8, 0xcf, 0xe5, 0x5e, 0xab, 0x56, 0xb1, 0x11, 0xd9, 0xb3,
	0x0f, 0x4a, 0x55, 0xd7, 0xaa, 0xb9, 0x72, 0xff, 0x5c, 0x2a, 0xff, 0x3b, 0x80, 0xe9, 0x4f, 0xbc,
	0x6b, 0x77, 0x3f, 0xa0, 0x31, 0x7c, 0x8b, 0x8c, 0xc1, 0xa8, 0x54, 0x02, 0xb3, 0x60, 0x1e, 0x5c,
	0x24, 0x85, 0xb5, 0x29, 0xb6, 0x53, 0xa6, 0xcd, 0x06, 0x2e, 0x46, 0x36, 0x9b, 0x41, 0xac, 0x51,
	0x48, 0x8d, 0x65, 0x9b, 0x0d, 0x6d, 0xfc, 0xe4, 0x13, 0xbe, 0x3d, 0xec, 0x31, 0x1b, 0x39, 0x3c,
	0xd9, 0xf9, 0x15, 0x24, 0xb6, 0xcf, 0xed, 0x61, 0x8f, 0x6c, 0x0a, 0xf1, 0x86, 0x97, 0xb8, 0x56,
	0xea, 0x2e, 0x7d, 0xc1, 0x00, 0xa2, 0xad, 0x52, 0xdb, 0x0a, 0xd3, 0x80, 0xc5, 0x30, 0x32, 0x58,
	0x6d, 0xd2, 0x41, 0xce, 0x21, 0x7d, 0x5b, 0x73, 0x59, 0x2d, 0xbb, 0xb5, 0x29, 0xb5, 0x5c, 0xa3,
	0x36, 0xec, 0x2b, 0x88, 0x4d, 0xb7, 0x5e, 0x55, 0xd2, 0xb4, 0x59, 0x30, 0x1f, 0x5e, 0x4c, 0xae,
	0x3f, 0xbc, 0xb4, 0x3a, 0x1f, 0x21, 0x8b, 0xb1, 0xe9, 0xd6, 0xef, 0xa4, 0x69, 0xd9, 0xc7, 0x90,
	0x18, 0x6c, 0xc4, 0x4a, 0xa8, 0x06, 0xb3, 0xc1, 0x7c, 0x78, 0x11, 0x17, 0x31, 0x05, 0x6e, 0x54,
	0x83, 0x39, 0x87, 0xf3, 0x47, 0x1f, 0xb2, 0xf7, 0x60, 0x20, 0x85, 0x17, 0x3f, 0x90, 0x82, 0xa5,
	0x30, 0xec, 0xa4, 0xf0, 0xca, 0xc9, 0x64, 0x2f, 0x21, 0x44, 0xfa, 0xc8, 0xab, 0x76, 0x0e, 0x7b,
	0x05, 0x91, 0x69, 0x79, 0xdb, 0x19, 0x2f, 0xda, 0x7b, 0xf9, 0x9f, 0x01, 0x84, 0xb7, 0xea, 0x0e,
	0x9b, 0x27, 0x95, 0x5f, 0x42, 0xd8, 0x52, 0xc2, 0xd7, 0x76, 0x0e, 0xfb, 0x04, 0xce, 0x34, 0x6e,
	0x34, 0x9a, 0xdd, 0xca, 0x65, 0x5d, 0x97, 0xa9, 0x0f, 0xba, 0x52, 0x19, 0x8c, 0x4b, 0x8d, 0xbc,
	0x45, 0x61, 0xbb, 0x0d, 0x8b, 0xa3, 0x4b, 0x53, 0xa9, 0x95, 0x90, 0x1b, 0x89, 0x22, 0x0b, 0x6d,
	0xea, 0xe4, 0x13, 0x95, 0xe9, 0x52, 0x95, 0x92, 0x57, 0x37, 0xd8, 0x12, 0x67, 0x06, 0xa3, 0x86,
	0xd7, 0xa7, 0x51, 0x93, 0x6d, 0xf5, 0xea, 0xea, 0xa4, 0x57, 0x5b, 0x65, 0xfc, 0x37, 0xde, 0x72,
	0x7d, 0x54, 0xe6, 0x3c, 0xf6, 0x11, 0x8c, 0xa5, 0x58, 0xc9, 0x66, 0xa3, 0x3c, 0xc7, 0x48, 0x8a,
	0x45, 0xb3, 0x51, 0xec, 0xf5, 0x51, 0x18, 0x11, 0x98, 0x5c, 0x4f, 0xdc, 0x84, 0x2c, 0x73, 0xaf,
	0x32, 0xff, 0x16, 0xc2, 0x5f, 0x0c, 0x0d, 0x74, 0x0e, 0x21, 0x65, 0x8d, 0x9f, 0x26, 0x38, 0x2c,
	0xe5, 0x0a, 0x97, 0xb0, 0xed, 0x9b, 0x72, 0xa7, 0xb4, 0xe7, 0xe4, 0xbd, 0xfc, 0xaf, 0x21, 0x8c,
	0x08, 0xf7, 0xe4, 0x5d, 0x8f, 0xaa, 0x06, 0x3d, 0x55, 0xaf, 0x61, 0x2a, 0xa4, 0xd9, 0x57, 0xfc,
	0xb0, 0xb2, 0x39, 0x47, 0x78, 0xe2, 0x63, 0x3f, 0x12, 0xe4, 0x15, 0x44, 0xfb, 0x9d, 0x6a, 0x90,
	0x06, 0x38, 0xa4, 0x3e, 0xce, 0xbb, 0x1f, 0x77, 0xd8, 0x1f, 0xf7, 0x0c, 0xe2, 0xdf, 0x71, 0x6d,
	0x64, 0x8b, 0x26, 0x8b, 0x2c, 0xfe, 0xe4, 0xb3, 0x2f, 0x61, 0x6c, 0xec, 0x33, 0x9b, 0x6c, 0x6c,
	0x55, 0x31, 0xa7, 0xaa, 0xff, 0xf6, 0xc5, 0x11, 0x42, 0xf5, 0x69, 0x55, 0x30, 0x8b, 0x5d, 0x7d,
	0xeb, 0xd0, 0xda, 0x6e, 0xb5, 0xea, 0xf6, 0x2b, 0x29, 0x4c, 0x96, 0xb8, 0x06, 0x36, 0xb0, 0x10,
	0x86, 0x7d, 0x0e, 0x91, 0xb5, 0x4d, 0x06, 0xb6, 0xfe, 0xb9, 0xab, 0xff, 0x3d, 0xc5, 0xec, 0xd3,
	0xf9, 0x34, 0xd5, 0x6e, 0x14, 0x51, 0x9c, 0xd8, 0x0a, 0xce, 0xe9, 0x6f, 0xcf, 0xf4, 0xe1, 0xf6,
	0xdc, 0x8f, 0xfa, 0xac, 0x3f, 0xea, 0x3c, 0x87, 0x84, 0xea, 0x2e, 0x2d, 0x35, 0x80, 0xa8, 0x51,
	0xba, 0xe6, 0x55, 0xfa, 0x82, 0x25, 0x10, 0x9a, 0x6e, 0x8f, 0x3a, 0x0d, 0xf2, 0x3f, 0x02, 0x48,
	0x4e, 0x0c, 0xfe, 0xd7, 0x50, 0x4e, 0xca, 0xc3, 0xbe, 0xf2, 0x1e, 0xbb, 0xe8, 0x01, 0xbb, 0xfc,
	0xd3, 0xe7, 0x58, 0x4c, 0x60, 0x5c, 0xf3, 0x86, 0x6f, 0x89, 0xc7, 0xf5, 0x3f, 0x01, 0x80, 0x83,
	0x69, 0xe4, 0x35, 0x7b, 0x03, 0xf0, 0x4e, 0x6d, 0x65, 0x63, 0x6f, 0x0f, 0xf3, 0x93, 0xe8, 0x1f,
	0xbc, 0x59, 0x7f, 0x3f, 0xd9, 0x1b, 0x78, 0xbf, 0xe8, 0xfd, 0xd2, 0xdc, 0x57, 0x7d, 0xc4, 0x43,
	0xf8, 0x17, 0x90, 0xd0, 0x95, 0x71, 0xbb, 0x9c, 0x5e, 0xfa, 0xcb, 0x5a, 0xe0, 0xaf, 0x3f, 0x77,
	0xa8, 0x0f, 0x47, 0xac, 0x4b, 0x7f, 0x06, 0x31, 0x19, 0xdf, 0x1d, 0x16, 0x82, 0x9d, 0xf5, 0xa0,
	0x0b, 0x31, 0xeb, 0xad, 0x3d, 0xfb, 0x06, 0x26, 0x37, 0x6a, 0x89, 0x8d, 0x78, 0xeb, 0xee, 0xca,
	0x7f, 0xde, 0x37, 0x33, 0x7b, 0x26, 0xbe, 0x8e, 0xec, 0x59, 0xff, 0xfa, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x11, 0xaa, 0xf3, 0xba, 0x04, 0x06, 0x00, 0x00,
}
