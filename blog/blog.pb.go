// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/blog.proto

package blog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "common"
import marker "marker"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Blogs struct {
	Blogs                []*Blog  `protobuf:"bytes,1,rep,name=blogs,proto3" json:"blogs,omitempty"`
	Anchor               string   `protobuf:"bytes,2,opt,name=anchor,proto3" json:"anchor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blogs) Reset()         { *m = Blogs{} }
func (m *Blogs) String() string { return proto.CompactTextString(m) }
func (*Blogs) ProtoMessage()    {}
func (*Blogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_ea40874efb4a8716, []int{0}
}
func (m *Blogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blogs.Unmarshal(m, b)
}
func (m *Blogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blogs.Marshal(b, m, deterministic)
}
func (dst *Blogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blogs.Merge(dst, src)
}
func (m *Blogs) XXX_Size() int {
	return xxx_messageInfo_Blogs.Size(m)
}
func (m *Blogs) XXX_DiscardUnknown() {
	xxx_messageInfo_Blogs.DiscardUnknown(m)
}

var xxx_messageInfo_Blogs proto.InternalMessageInfo

func (m *Blogs) GetBlogs() []*Blog {
	if m != nil {
		return m.Blogs
	}
	return nil
}

func (m *Blogs) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

type Blog struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string             `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Author               string             `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	Created              int32              `protobuf:"varint,7,opt,name=created,proto3" json:"created,omitempty"`
	State                int32              `protobuf:"varint,8,opt,name=state,proto3" json:"state,omitempty"`
	Tag                  string             `protobuf:"bytes,9,opt,name=tag,proto3" json:"tag,omitempty"`
	Tags                 []*marker.Tag      `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	Content              string             `protobuf:"bytes,11,opt,name=content,proto3" json:"content,omitempty"`
	Banner               string             `protobuf:"bytes,12,opt,name=banner,proto3" json:"banner,omitempty"`
	Categories           []*marker.Category `protobuf:"bytes,13,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_ea40874efb4a8716, []int{1}
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

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Blog) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Blog) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *Blog) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Blog) GetTags() []*marker.Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Blog) GetBanner() string {
	if m != nil {
		return m.Banner
	}
	return ""
}

func (m *Blog) GetCategories() []*marker.Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type FileDatas struct {
	Files                []*FileData `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Anchor               string      `protobuf:"bytes,2,opt,name=anchor,proto3" json:"anchor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FileDatas) Reset()         { *m = FileDatas{} }
func (m *FileDatas) String() string { return proto.CompactTextString(m) }
func (*FileDatas) ProtoMessage()    {}
func (*FileDatas) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_ea40874efb4a8716, []int{2}
}
func (m *FileDatas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileDatas.Unmarshal(m, b)
}
func (m *FileDatas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileDatas.Marshal(b, m, deterministic)
}
func (dst *FileDatas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileDatas.Merge(dst, src)
}
func (m *FileDatas) XXX_Size() int {
	return xxx_messageInfo_FileDatas.Size(m)
}
func (m *FileDatas) XXX_DiscardUnknown() {
	xxx_messageInfo_FileDatas.DiscardUnknown(m)
}

var xxx_messageInfo_FileDatas proto.InternalMessageInfo

func (m *FileDatas) GetFiles() []*FileData {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *FileDatas) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

type FileData struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Tags                 *marker.Tag         `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Created              int32               `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Media                []*common.MediaInfo `protobuf:"bytes,5,rep,name=media,proto3" json:"media,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FileData) Reset()         { *m = FileData{} }
func (m *FileData) String() string { return proto.CompactTextString(m) }
func (*FileData) ProtoMessage()    {}
func (*FileData) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_ea40874efb4a8716, []int{3}
}
func (m *FileData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileData.Unmarshal(m, b)
}
func (m *FileData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileData.Marshal(b, m, deterministic)
}
func (dst *FileData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileData.Merge(dst, src)
}
func (m *FileData) XXX_Size() int {
	return xxx_messageInfo_FileData.Size(m)
}
func (m *FileData) XXX_DiscardUnknown() {
	xxx_messageInfo_FileData.DiscardUnknown(m)
}

var xxx_messageInfo_FileData proto.InternalMessageInfo

func (m *FileData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FileData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *FileData) GetTags() *marker.Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FileData) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *FileData) GetMedia() []*common.MediaInfo {
	if m != nil {
		return m.Media
	}
	return nil
}

type Notes struct {
	Notes                []*Note  `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	Anchor               string   `protobuf:"bytes,2,opt,name=anchor,proto3" json:"anchor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notes) Reset()         { *m = Notes{} }
func (m *Notes) String() string { return proto.CompactTextString(m) }
func (*Notes) ProtoMessage()    {}
func (*Notes) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_ea40874efb4a8716, []int{4}
}
func (m *Notes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notes.Unmarshal(m, b)
}
func (m *Notes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notes.Marshal(b, m, deterministic)
}
func (dst *Notes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notes.Merge(dst, src)
}
func (m *Notes) XXX_Size() int {
	return xxx_messageInfo_Notes.Size(m)
}
func (m *Notes) XXX_DiscardUnknown() {
	xxx_messageInfo_Notes.DiscardUnknown(m)
}

var xxx_messageInfo_Notes proto.InternalMessageInfo

func (m *Notes) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

func (m *Notes) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

type Note struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string              `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Created              int32               `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Title                string              `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	PosX                 int32               `protobuf:"varint,5,opt,name=posX,proto3" json:"posX,omitempty"`
	PosY                 int32               `protobuf:"varint,6,opt,name=posY,proto3" json:"posY,omitempty"`
	Tags                 []*marker.Tag       `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	Media                []*common.MediaInfo `protobuf:"bytes,7,rep,name=media,proto3" json:"media,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_ea40874efb4a8716, []int{5}
}
func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (dst *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(dst, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Note) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Note) GetCreated() int32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Note) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Note) GetPosX() int32 {
	if m != nil {
		return m.PosX
	}
	return 0
}

func (m *Note) GetPosY() int32 {
	if m != nil {
		return m.PosY
	}
	return 0
}

func (m *Note) GetTags() []*marker.Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Note) GetMedia() []*common.MediaInfo {
	if m != nil {
		return m.Media
	}
	return nil
}

func init() {
	proto.RegisterType((*Blogs)(nil), "blog.Blogs")
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*FileDatas)(nil), "blog.FileDatas")
	proto.RegisterType((*FileData)(nil), "blog.FileData")
	proto.RegisterType((*Notes)(nil), "blog.Notes")
	proto.RegisterType((*Note)(nil), "blog.Note")
}

func init() { proto.RegisterFile("blog/blog.proto", fileDescriptor_blog_ea40874efb4a8716) }

var fileDescriptor_blog_ea40874efb4a8716 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x4d, 0x4b, 0x0b, 0xcb, 0x45, 0x58, 0x9c, 0x35, 0x66, 0xd2, 0x17, 0x09, 0x31, 0x91, 0x6c,
	0x14, 0x14, 0xbf, 0x40, 0x5d, 0x4d, 0x48, 0x76, 0x4d, 0xac, 0x9a, 0xb8, 0x8f, 0x03, 0xcc, 0xd6,
	0x66, 0x4b, 0x07, 0x67, 0x86, 0x87, 0xfd, 0x06, 0x5f, 0xfd, 0x2f, 0x13, 0xbf, 0xc8, 0xcc, 0xbd,
	0x2d, 0x94, 0x25, 0x5d, 0x5e, 0xe8, 0x9c, 0x7b, 0xee, 0xdc, 0x9e, 0x39, 0x3d, 0x03, 0x9c, 0xce,
	0x33, 0x95, 0x4c, 0xdc, 0xcf, 0x78, 0xad, 0x95, 0x55, 0x2c, 0x70, 0xeb, 0xe8, 0x6c, 0xa1, 0x56,
	0x2b, 0x95, 0x4f, 0xe8, 0x41, 0x54, 0x74, 0xb6, 0x12, 0xfa, 0x56, 0xea, 0x09, 0x3d, 0xa8, 0x38,
	0x7c, 0x07, 0xe1, 0xfb, 0x4c, 0x25, 0x86, 0x0d, 0x20, 0x74, 0x5b, 0x0d, 0xf7, 0x06, 0x8d, 0x51,
	0x67, 0x0a, 0x63, 0x1c, 0xea, 0xb8, 0x98, 0x08, 0xf6, 0x14, 0x9a, 0x22, 0x5f, 0xfc, 0x54, 0x9a,
	0xfb, 0x03, 0x6f, 0xd4, 0x8e, 0x0b, 0x34, 0xfc, 0xed, 0x43, 0xe0, 0xfa, 0x58, 0x0f, 0xfc, 0x74,
	0xc9, 0x3d, 0x24, 0xfd, 0x74, 0xc9, 0x9e, 0x40, 0x68, 0x53, 0x9b, 0x49, 0xde, 0xc0, 0x12, 0x01,
	0x1c, 0xb3, 0xb1, 0x6e, 0x4c, 0xb3, 0x18, 0x83, 0x88, 0x71, 0x68, 0x2d, 0xb4, 0x14, 0x56, 0x2e,
	0x79, 0x6b, 0xe0, 0x8d, 0xc2, 0xb8, 0x84, 0x6e, 0x8e, 0xb1, 0xc2, 0x4a, 0x7e, 0x82, 0x75, 0x02,
	0xac, 0x0f, 0x0d, 0x2b, 0x12, 0xde, 0xc6, 0x21, 0x6e, 0xc9, 0x9e, 0x41, 0x60, 0x45, 0x62, 0x38,
	0xe0, 0x09, 0x3a, 0xe3, 0xe2, 0xa0, 0xdf, 0x44, 0x12, 0x23, 0x81, 0xaf, 0x50, 0xb9, 0x95, 0xb9,
	0xe5, 0x1d, 0xdc, 0x56, 0x42, 0x27, 0x6a, 0x2e, 0xf2, 0x5c, 0x6a, 0xfe, 0x88, 0x44, 0x11, 0x62,
	0xaf, 0x01, 0x16, 0xc2, 0xca, 0x44, 0xe9, 0x54, 0x1a, 0xde, 0xc5, 0xc1, 0xfd, 0x72, 0xf0, 0x07,
	0x62, 0xee, 0xe2, 0x4a, 0xcf, 0x70, 0x06, 0xed, 0x4f, 0x69, 0x26, 0x2f, 0x84, 0x15, 0x86, 0x3d,
	0x87, 0xf0, 0x26, 0xcd, 0x64, 0x69, 0x6a, 0x8f, 0x4c, 0x2d, 0xf9, 0x98, 0xc8, 0x5a, 0x63, 0xff,
	0x78, 0x70, 0x52, 0xf6, 0xd6, 0x9b, 0xeb, 0x57, 0xcd, 0x2d, 0x2d, 0x70, 0x8e, 0xd7, 0x5a, 0x50,
	0xb8, 0x1c, 0xec, 0xbb, 0xfc, 0x02, 0xc2, 0x95, 0x5c, 0xa6, 0x82, 0x87, 0xa8, 0xf5, 0xf1, 0xb8,
	0x08, 0xcf, 0x95, 0x2b, 0xce, 0xf2, 0x1b, 0x15, 0x13, 0xef, 0x22, 0xf3, 0x59, 0x59, 0x89, 0x91,
	0xc9, 0xdd, 0x62, 0x3f, 0x32, 0x8e, 0x8b, 0x89, 0xa8, 0x3d, 0xd9, 0x3f, 0x0f, 0x02, 0xd7, 0x77,
	0x70, 0xaa, 0xca, 0x17, 0xf2, 0xf7, 0xbf, 0x50, 0x45, 0x78, 0xe3, 0x20, 0x1e, 0xe4, 0x44, 0x50,
	0x75, 0x82, 0x41, 0xb0, 0x56, 0xe6, 0x07, 0x0f, 0xb1, 0x19, 0xd7, 0x45, 0xed, 0x1a, 0x83, 0x47,
	0xb5, 0xeb, 0xe3, 0xa1, 0xd9, 0xfa, 0xd2, 0x7a, 0xd8, 0x97, 0xe9, 0xdf, 0x06, 0x80, 0xbb, 0x07,
	0x5f, 0xad, 0x96, 0x62, 0xc5, 0xce, 0xa1, 0x7d, 0x99, 0x1a, 0x4b, 0xb7, 0xab, 0x5f, 0xee, 0x8a,
	0xe5, 0xaf, 0x2f, 0x1b, 0xa9, 0xef, 0xa2, 0xce, 0xee, 0x82, 0xb9, 0x9c, 0xc0, 0xf7, 0xb5, 0x91,
	0x1a, 0xbb, 0x59, 0xe5, 0xee, 0x45, 0x95, 0x35, 0x7b, 0x05, 0x70, 0x21, 0x33, 0x69, 0x25, 0xa2,
	0x6e, 0x65, 0xe4, 0x6c, 0x19, 0x6d, 0xdf, 0xf0, 0x51, 0x6b, 0xa5, 0xaf, 0x4c, 0xc2, 0xa6, 0xd0,
	0x75, 0x02, 0x76, 0x69, 0x3c, 0x14, 0x71, 0xba, 0x1f, 0x48, 0xc3, 0x5e, 0x02, 0xcc, 0x72, 0x27,
	0xc4, 0x95, 0xd8, 0xbd, 0xbc, 0x46, 0xf7, 0x30, 0x7b, 0x03, 0x3d, 0x12, 0xb4, 0xad, 0x1c, 0x15,
	0x55, 0xb8, 0x42, 0x01, 0xaa, 0x75, 0x85, 0xe8, 0x73, 0x80, 0x4b, 0xb5, 0xb8, 0xdd, 0xac, 0x31,
	0x2a, 0x87, 0xcd, 0x95, 0xc0, 0xed, 0x1c, 0x44, 0x54, 0x61, 0xf6, 0xba, 0xb6, 0x0e, 0x22, 0x3a,
	0x26, 0x76, 0xde, 0xc4, 0xff, 0xc8, 0xb7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x5e, 0x94,
	0x73, 0x66, 0x05, 0x00, 0x00,
}
