// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth/auth.proto

It has these top-level messages:
	Done
	SignatureReq
	Signature
	TokenRequest
	Token
*/
package auth

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

type Done struct {
}

func (m *Done) Reset()                    { *m = Done{} }
func (m *Done) String() string            { return proto.CompactTextString(m) }
func (*Done) ProtoMessage()               {}
func (*Done) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SignatureReq struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
}

func (m *SignatureReq) Reset()                    { *m = SignatureReq{} }
func (m *SignatureReq) String() string            { return proto.CompactTextString(m) }
func (*SignatureReq) ProtoMessage()               {}
func (*SignatureReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignatureReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SignatureReq) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

type Signature struct {
	Key     string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Created int64  `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Signature) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Signature) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type TokenRequest struct {
	Payload             map[string]string `protobuf:"bytes,1,rep,name=payload" json:"payload,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AccessTokenExpired  int64             `protobuf:"varint,2,opt,name=access_token_expired,json=accessTokenExpired" json:"access_token_expired,omitempty"`
	RefreshTokenExpired int64             `protobuf:"varint,3,opt,name=refresh_token_expired,json=refreshTokenExpired" json:"refresh_token_expired,omitempty"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TokenRequest) GetPayload() map[string]string {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TokenRequest) GetAccessTokenExpired() int64 {
	if m != nil {
		return m.AccessTokenExpired
	}
	return 0
}

func (m *TokenRequest) GetRefreshTokenExpired() int64 {
	if m != nil {
		return m.RefreshTokenExpired
	}
	return 0
}

type Token struct {
	AccessToken         string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	Created             string `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	AccessTokenExpired  int64  `protobuf:"varint,3,opt,name=access_token_expired,json=accessTokenExpired" json:"access_token_expired,omitempty"`
	RefreshToken        string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	RefreshTokenExpired int64  `protobuf:"varint,5,opt,name=refresh_token_expired,json=refreshTokenExpired" json:"refresh_token_expired,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Token) GetAccessTokenExpired() int64 {
	if m != nil {
		return m.AccessTokenExpired
	}
	return 0
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetRefreshTokenExpired() int64 {
	if m != nil {
		return m.RefreshTokenExpired
	}
	return 0
}

func init() {
	proto.RegisterType((*Done)(nil), "auth.Done")
	proto.RegisterType((*SignatureReq)(nil), "auth.SignatureReq")
	proto.RegisterType((*Signature)(nil), "auth.Signature")
	proto.RegisterType((*TokenRequest)(nil), "auth.TokenRequest")
	proto.RegisterType((*Token)(nil), "auth.Token")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthProvider service

type AuthProviderClient interface {
	MakeToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Token, error)
	RefreshToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	CreatSignature(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Signature, error)
}

type authProviderClient struct {
	cc *grpc.ClientConn
}

func NewAuthProviderClient(cc *grpc.ClientConn) AuthProviderClient {
	return &authProviderClient{cc}
}

func (c *authProviderClient) MakeToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/auth.AuthProvider/MakeToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authProviderClient) RefreshToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/auth.AuthProvider/RefreshToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authProviderClient) CreatSignature(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Signature, error) {
	out := new(Signature)
	err := grpc.Invoke(ctx, "/auth.AuthProvider/CreatSignature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthProvider service

type AuthProviderServer interface {
	MakeToken(context.Context, *TokenRequest) (*Token, error)
	RefreshToken(context.Context, *Token) (*Token, error)
	CreatSignature(context.Context, *Token) (*Signature, error)
}

func RegisterAuthProviderServer(s *grpc.Server, srv AuthProviderServer) {
	s.RegisterService(&_AuthProvider_serviceDesc, srv)
}

func _AuthProvider_MakeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthProviderServer).MakeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthProvider/MakeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthProviderServer).MakeToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthProvider_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthProviderServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthProvider/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthProviderServer).RefreshToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthProvider_CreatSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthProviderServer).CreatSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthProvider/CreatSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthProviderServer).CreatSignature(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthProvider",
	HandlerType: (*AuthProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeToken",
			Handler:    _AuthProvider_MakeToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthProvider_RefreshToken_Handler,
		},
		{
			MethodName: "CreatSignature",
			Handler:    _AuthProvider_CreatSignature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

func init() { proto.RegisterFile("auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xcd, 0xd2, 0x02, 0x61, 0xa8, 0x62, 0x56, 0x4c, 0x1a, 0x2e, 0x62, 0xbd, 0xf4, 0x60, 0xaa,
	0xc1, 0xc4, 0x0f, 0x6e, 0x46, 0x39, 0x9a, 0x90, 0xea, 0x9d, 0xac, 0x74, 0xb0, 0x04, 0x6c, 0xcb,
	0x76, 0x4b, 0xe4, 0x8f, 0xf8, 0xb7, 0xfc, 0x21, 0xfe, 0x09, 0xd3, 0xdd, 0x16, 0x16, 0x4d, 0x2f,
	0x64, 0x66, 0xde, 0x9b, 0xc7, 0x7b, 0xd3, 0x85, 0x0e, 0xcb, 0x44, 0x78, 0x99, 0xff, 0x78, 0x09,
	0x8f, 0x45, 0x4c, 0xcd, 0xbc, 0x76, 0x1a, 0x60, 0x3e, 0xc5, 0x11, 0x3a, 0x37, 0x60, 0xbd, 0xcc,
	0xdf, 0x23, 0x26, 0x32, 0x8e, 0x3e, 0xae, 0x28, 0x05, 0x33, 0x61, 0x22, 0xb4, 0x49, 0x9f, 0xb8,
	0x2d, 0x5f, 0xd6, 0xf9, 0x6c, 0xc6, 0xe3, 0x0f, 0xbb, 0xa6, 0x66, 0x79, 0xed, 0xdc, 0x42, 0x6b,
	0xbb, 0x47, 0x8f, 0xc0, 0x58, 0xe0, 0xa6, 0xc0, 0xf3, 0x92, 0xda, 0xd0, 0x9c, 0x72, 0x64, 0x02,
	0x03, 0xdb, 0xe8, 0x13, 0xd7, 0xf0, 0xcb, 0xd6, 0xf9, 0x21, 0x60, 0xbd, 0xc6, 0x0b, 0x8c, 0x7c,
	0x5c, 0x65, 0x98, 0x0a, 0x7a, 0x0f, 0xcd, 0x84, 0x6d, 0x96, 0x31, 0x0b, 0x6c, 0xd2, 0x37, 0xdc,
	0xf6, 0xe0, 0xd4, 0x93, 0x6e, 0x75, 0x92, 0x37, 0x56, 0x8c, 0x51, 0x24, 0xf8, 0xc6, 0x2f, 0xf9,
	0xf4, 0x0a, 0xba, 0x6c, 0x3a, 0xc5, 0x34, 0x9d, 0x88, 0x9c, 0x3c, 0xc1, 0xcf, 0x64, 0xce, 0x31,
	0x90, 0x46, 0x0c, 0x9f, 0x2a, 0x4c, 0xea, 0x8c, 0x14, 0x42, 0x07, 0x70, 0xc2, 0x71, 0xc6, 0x31,
	0x0d, 0xff, 0xac, 0x28, 0x97, 0xc7, 0x05, 0xa8, 0xef, 0xf4, 0x86, 0x60, 0xe9, 0x7f, 0x5f, 0xa6,
	0x25, 0xbb, 0xb4, 0x5d, 0xa8, 0xaf, 0xd9, 0x32, 0xc3, 0xe2, 0x02, 0xaa, 0x19, 0xd6, 0xee, 0x88,
	0xf3, 0x4d, 0xa0, 0x2e, 0xc5, 0xe8, 0x19, 0x58, 0xba, 0xd7, 0x62, 0xbd, 0xad, 0x79, 0xd4, 0x8f,
	0xa6, 0x84, 0xca, 0xb6, 0x32, 0xa8, 0x51, 0x19, 0xf4, 0x1c, 0x0e, 0xf6, 0x82, 0xda, 0xa6, 0x54,
	0xb4, 0xf4, 0x80, 0xd5, 0xd7, 0xa8, 0x57, 0x5e, 0x63, 0xf0, 0x45, 0xc0, 0x7a, 0xc8, 0x44, 0x38,
	0xe6, 0xf1, 0x7a, 0x1e, 0x20, 0xa7, 0x17, 0xd0, 0x7a, 0x66, 0x0b, 0x54, 0x8a, 0xf4, 0xff, 0xb7,
	0xeb, 0xb5, 0xb5, 0x19, 0x75, 0xc1, 0xf2, 0x75, 0x0b, 0x3a, 0xb8, 0xcf, 0xf4, 0xe0, 0xf0, 0x31,
	0x8f, 0xbf, 0x7b, 0x66, 0x7b, 0xdc, 0x8e, 0x6a, 0xb6, 0xe8, 0x5b, 0x43, 0x3e, 0xef, 0xeb, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x89, 0x0f, 0x72, 0xf1, 0x02, 0x00, 0x00,
}
