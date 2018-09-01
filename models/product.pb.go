// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

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

type Category_State int32

const (
	Category_ACTIVE         Category_State = 0
	Category_INACTIVE       Category_State = 1
	Category_OUT_OF_PRODUCT Category_State = 2
)

var Category_State_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
	2: "OUT_OF_PRODUCT",
}
var Category_State_value = map[string]int32{
	"ACTIVE":         0,
	"INACTIVE":       1,
	"OUT_OF_PRODUCT": 2,
}

func (x Category_State) String() string {
	return proto.EnumName(Category_State_name, int32(x))
}
func (Category_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{1, 0}
}

type Product_State int32

const (
	Product_ACTIVE   Product_State = 0
	Product_INACTIVE Product_State = 1
	Product_SOLD_OUT Product_State = 2
)

var Product_State_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
	2: "SOLD_OUT",
}
var Product_State_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
	"SOLD_OUT": 2,
}

func (x Product_State) String() string {
	return proto.EnumName(Product_State_name, int32(x))
}
func (Product_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{3, 0}
}

type Tag_State int32

const (
	Tag_ACTIVE   Tag_State = 0
	Tag_INACTIVE Tag_State = 1
)

var Tag_State_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
}
var Tag_State_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
}

func (x Tag_State) String() string {
	return proto.EnumName(Tag_State_name, int32(x))
}
func (Tag_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{5, 0}
}

type Rate_State int32

const (
	Rate_ACTIVE   Rate_State = 0
	Rate_INACTIVE Rate_State = 1
)

var Rate_State_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
}
var Rate_State_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
}

func (x Rate_State) String() string {
	return proto.EnumName(Rate_State_name, int32(x))
}
func (Rate_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{6, 0}
}

type Image struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{0}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Category struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Display              string               `protobuf:"bytes,2,opt,name=display,proto3" json:"display,omitempty"`
	Desscription         string               `protobuf:"bytes,3,opt,name=desscription,proto3" json:"desscription,omitempty"`
	State                Category_State       `protobuf:"varint,4,opt,name=state,proto3,enum=mpb.Category_State" json:"state,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{1}
}
func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (dst *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(dst, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Category) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *Category) GetDesscription() string {
	if m != nil {
		return m.Desscription
	}
	return ""
}

func (m *Category) GetState() Category_State {
	if m != nil {
		return m.State
	}
	return Category_ACTIVE
}

func (m *Category) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type Group struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Display              string               `protobuf:"bytes,2,opt,name=display,proto3" json:"display,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Target               string               `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	BonusPrice           float32              `protobuf:"fixed32,6,opt,name=bonus_price,json=bonusPrice,proto3" json:"bonus_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{2}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (dst *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(dst, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Group) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *Group) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Group) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Group) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Group) GetBonusPrice() float32 {
	if m != nil {
		return m.BonusPrice
	}
	return 0
}

type Product struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Images               []*Image             `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`
	Description          string               `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Categories           []*Category          `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty"`
	Groups               []*Group             `protobuf:"bytes,7,rep,name=groups,proto3" json:"groups,omitempty"`
	State                Product_State        `protobuf:"varint,10,opt,name=state,proto3,enum=mpb.Product_State" json:"state,omitempty"`
	ImportPrice          int32                `protobuf:"varint,11,opt,name=import_price,json=importPrice,proto3" json:"import_price,omitempty"`
	ExportPrice          int32                `protobuf:"varint,12,opt,name=export_price,json=exportPrice,proto3" json:"export_price,omitempty"`
	ImportFrom           *Import              `protobuf:"bytes,13,opt,name=import_from,json=importFrom,proto3" json:"import_from,omitempty"`
	Tag                  []*Tag               `protobuf:"bytes,14,rep,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{3}
}
func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (dst *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(dst, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Product) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Product) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *Product) GetState() Product_State {
	if m != nil {
		return m.State
	}
	return Product_ACTIVE
}

func (m *Product) GetImportPrice() int32 {
	if m != nil {
		return m.ImportPrice
	}
	return 0
}

func (m *Product) GetExportPrice() int32 {
	if m != nil {
		return m.ExportPrice
	}
	return 0
}

func (m *Product) GetImportFrom() *Import {
	if m != nil {
		return m.ImportFrom
	}
	return nil
}

func (m *Product) GetTag() []*Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

type Import struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Display              string   `protobuf:"bytes,2,opt,name=display,proto3" json:"display,omitempty"`
	Desscription         string   `protobuf:"bytes,3,opt,name=desscription,proto3" json:"desscription,omitempty"`
	Note                 []string `protobuf:"bytes,4,rep,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Import) Reset()         { *m = Import{} }
func (m *Import) String() string { return proto.CompactTextString(m) }
func (*Import) ProtoMessage()    {}
func (*Import) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{4}
}
func (m *Import) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Import.Unmarshal(m, b)
}
func (m *Import) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Import.Marshal(b, m, deterministic)
}
func (dst *Import) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Import.Merge(dst, src)
}
func (m *Import) XXX_Size() int {
	return xxx_messageInfo_Import.Size(m)
}
func (m *Import) XXX_DiscardUnknown() {
	xxx_messageInfo_Import.DiscardUnknown(m)
}

var xxx_messageInfo_Import proto.InternalMessageInfo

func (m *Import) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Import) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *Import) GetDesscription() string {
	if m != nil {
		return m.Desscription
	}
	return ""
}

func (m *Import) GetNote() []string {
	if m != nil {
		return m.Note
	}
	return nil
}

type Tag struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Display              string    `protobuf:"bytes,2,opt,name=display,proto3" json:"display,omitempty"`
	Note                 []string  `protobuf:"bytes,4,rep,name=note,proto3" json:"note,omitempty"`
	State                Tag_State `protobuf:"varint,5,opt,name=state,proto3,enum=mpb.Tag_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{5}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (dst *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(dst, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tag) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *Tag) GetNote() []string {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *Tag) GetState() Tag_State {
	if m != nil {
		return m.State
	}
	return Tag_ACTIVE
}

type Rate struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	By                   string               `protobuf:"bytes,3,opt,name=by,proto3" json:"by,omitempty"`
	Score                float32              `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Rate) Reset()         { *m = Rate{} }
func (m *Rate) String() string { return proto.CompactTextString(m) }
func (*Rate) ProtoMessage()    {}
func (*Rate) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{6}
}
func (m *Rate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rate.Unmarshal(m, b)
}
func (m *Rate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rate.Marshal(b, m, deterministic)
}
func (dst *Rate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rate.Merge(dst, src)
}
func (m *Rate) XXX_Size() int {
	return xxx_messageInfo_Rate.Size(m)
}
func (m *Rate) XXX_DiscardUnknown() {
	xxx_messageInfo_Rate.DiscardUnknown(m)
}

var xxx_messageInfo_Rate proto.InternalMessageInfo

func (m *Rate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Rate) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Rate) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

func (m *Rate) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Rate) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type ProductsDetail struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	GroupId              string   `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	ImportPrice          int32    `protobuf:"varint,4,opt,name=importPrice,proto3" json:"importPrice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductsDetail) Reset()         { *m = ProductsDetail{} }
func (m *ProductsDetail) String() string { return proto.CompactTextString(m) }
func (*ProductsDetail) ProtoMessage()    {}
func (*ProductsDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{7}
}
func (m *ProductsDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductsDetail.Unmarshal(m, b)
}
func (m *ProductsDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductsDetail.Marshal(b, m, deterministic)
}
func (dst *ProductsDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductsDetail.Merge(dst, src)
}
func (m *ProductsDetail) XXX_Size() int {
	return xxx_messageInfo_ProductsDetail.Size(m)
}
func (m *ProductsDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductsDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ProductsDetail proto.InternalMessageInfo

func (m *ProductsDetail) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *ProductsDetail) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *ProductsDetail) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ProductsDetail) GetImportPrice() int32 {
	if m != nil {
		return m.ImportPrice
	}
	return 0
}

type ImportProduct struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Products             []*ProductsDetail    `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	TotalPrice           int32                `protobuf:"varint,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Discount             float32              `protobuf:"fixed32,5,opt,name=discount,proto3" json:"discount,omitempty"`
	RealPay              int32                `protobuf:"varint,6,opt,name=real_pay,json=realPay,proto3" json:"real_pay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ImportProduct) Reset()         { *m = ImportProduct{} }
func (m *ImportProduct) String() string { return proto.CompactTextString(m) }
func (*ImportProduct) ProtoMessage()    {}
func (*ImportProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_f7e6dc8c1a043a00, []int{8}
}
func (m *ImportProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportProduct.Unmarshal(m, b)
}
func (m *ImportProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportProduct.Marshal(b, m, deterministic)
}
func (dst *ImportProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportProduct.Merge(dst, src)
}
func (m *ImportProduct) XXX_Size() int {
	return xxx_messageInfo_ImportProduct.Size(m)
}
func (m *ImportProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportProduct.DiscardUnknown(m)
}

var xxx_messageInfo_ImportProduct proto.InternalMessageInfo

func (m *ImportProduct) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ImportProduct) GetProducts() []*ProductsDetail {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *ImportProduct) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *ImportProduct) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *ImportProduct) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *ImportProduct) GetRealPay() int32 {
	if m != nil {
		return m.RealPay
	}
	return 0
}

func init() {
	proto.RegisterType((*Image)(nil), "mpb.Image")
	proto.RegisterType((*Category)(nil), "mpb.Category")
	proto.RegisterType((*Group)(nil), "mpb.Group")
	proto.RegisterType((*Product)(nil), "mpb.Product")
	proto.RegisterType((*Import)(nil), "mpb.Import")
	proto.RegisterType((*Tag)(nil), "mpb.Tag")
	proto.RegisterType((*Rate)(nil), "mpb.Rate")
	proto.RegisterType((*ProductsDetail)(nil), "mpb.ProductsDetail")
	proto.RegisterType((*ImportProduct)(nil), "mpb.ImportProduct")
	proto.RegisterEnum("mpb.Category_State", Category_State_name, Category_State_value)
	proto.RegisterEnum("mpb.Product_State", Product_State_name, Product_State_value)
	proto.RegisterEnum("mpb.Tag_State", Tag_State_name, Tag_State_value)
	proto.RegisterEnum("mpb.Rate_State", Rate_State_name, Rate_State_value)
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_product_f7e6dc8c1a043a00) }

var fileDescriptor_product_f7e6dc8c1a043a00 = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6a, 0xdb, 0x4a,
	0x10, 0x3e, 0x92, 0x2c, 0xd9, 0x19, 0xff, 0x10, 0x36, 0x87, 0x83, 0x62, 0x38, 0xc4, 0x11, 0xbd,
	0x70, 0xa1, 0xb5, 0x21, 0x6d, 0x1f, 0xa0, 0x24, 0x4d, 0x31, 0x94, 0xda, 0x6c, 0x9c, 0xde, 0x9a,
	0xb5, 0xb4, 0x11, 0x2a, 0x96, 0x57, 0x68, 0xd7, 0x50, 0xdf, 0xf6, 0xba, 0xef, 0xd2, 0x77, 0xe8,
	0x83, 0xf4, 0xb6, 0xb7, 0x7d, 0x84, 0xb2, 0xb3, 0xab, 0x20, 0x27, 0x81, 0xfc, 0xd0, 0x3b, 0xcd,
	0x37, 0x9f, 0x76, 0xe7, 0x9b, 0xf9, 0x66, 0xa1, 0x5b, 0x94, 0x22, 0xd9, 0xc4, 0x6a, 0x54, 0x94,
	0x42, 0x09, 0xe2, 0xe5, 0xc5, 0xb2, 0x7f, 0x94, 0x0a, 0x91, 0xae, 0xf8, 0x18, 0xa1, 0xe5, 0xe6,
	0x6a, 0xac, 0xb2, 0x9c, 0x4b, 0xc5, 0xf2, 0xc2, 0xb0, 0xa2, 0x43, 0xf0, 0x27, 0x39, 0x4b, 0x39,
	0xd9, 0x07, 0x6f, 0x53, 0xae, 0x42, 0x67, 0xe0, 0x0c, 0xf7, 0xa8, 0xfe, 0x8c, 0x7e, 0x3b, 0xd0,
	0x3a, 0x65, 0x8a, 0xa7, 0xa2, 0xdc, 0x92, 0x1e, 0xb8, 0x59, 0x62, 0xb3, 0x6e, 0x96, 0x90, 0x10,
	0x9a, 0x49, 0x26, 0x8b, 0x15, 0xdb, 0x86, 0x2e, 0x82, 0x55, 0x48, 0x22, 0xe8, 0x24, 0x5c, 0xca,
	0xb8, 0xcc, 0x0a, 0x95, 0x89, 0x75, 0xe8, 0x61, 0x7a, 0x07, 0x23, 0xcf, 0xc1, 0x97, 0x8a, 0x29,
	0x1e, 0x36, 0x06, 0xce, 0xb0, 0x77, 0x72, 0x30, 0xca, 0x8b, 0xe5, 0xa8, 0xba, 0x6b, 0x74, 0xa1,
	0x53, 0xd4, 0x30, 0xc8, 0x6b, 0x68, 0xc6, 0x25, 0x67, 0x8a, 0x27, 0xa1, 0x3f, 0x70, 0x86, 0xed,
	0x93, 0xfe, 0xc8, 0x68, 0x1a, 0x55, 0x9a, 0x46, 0xf3, 0x4a, 0x13, 0xad, 0xa8, 0xd1, 0x1b, 0xf0,
	0xf1, 0x14, 0x02, 0x10, 0xbc, 0x3d, 0x9d, 0x4f, 0x3e, 0xbd, 0xdb, 0xff, 0x87, 0x74, 0xa0, 0x35,
	0xf9, 0x68, 0x23, 0x87, 0x10, 0xe8, 0x4d, 0x2f, 0xe7, 0x8b, 0xe9, 0xf9, 0x62, 0x46, 0xa7, 0x67,
	0x97, 0xa7, 0xf3, 0x7d, 0x37, 0xfa, 0xe1, 0x80, 0xff, 0xbe, 0x14, 0x9b, 0xe2, 0x11, 0x7a, 0x07,
	0xd0, 0x4e, 0xf8, 0x4d, 0xb9, 0x75, 0x88, 0xfc, 0x07, 0x81, 0x62, 0x65, 0xca, 0x15, 0xca, 0xdd,
	0xa3, 0x36, 0x7a, 0x9a, 0x34, 0x72, 0x04, 0xed, 0xa5, 0x58, 0x6f, 0xe4, 0xa2, 0x28, 0xb3, 0x98,
	0x87, 0xc1, 0xc0, 0x19, 0xba, 0x14, 0x10, 0x9a, 0x69, 0x24, 0xfa, 0xe5, 0x41, 0x73, 0x66, 0xac,
	0x70, 0x4b, 0x06, 0x81, 0xc6, 0x9a, 0xe5, 0xdc, 0x6a, 0xc0, 0x6f, 0x12, 0x41, 0x90, 0x69, 0x0b,
	0xc8, 0xd0, 0x1b, 0x78, 0xc3, 0xf6, 0x09, 0xe0, 0x34, 0xd0, 0x15, 0xd4, 0x66, 0x6e, 0x8a, 0x6c,
	0xdc, 0x16, 0xf9, 0x34, 0x31, 0x2f, 0x01, 0x62, 0x33, 0xf6, 0x8c, 0xcb, 0x30, 0xc0, 0xfb, 0xbb,
	0x3b, 0x6e, 0xa0, 0x35, 0x82, 0x2e, 0x35, 0xd5, 0xe3, 0x91, 0x61, 0xb3, 0x56, 0x2a, 0x4e, 0x8c,
	0xda, 0x0c, 0x19, 0x56, 0xde, 0x02, 0xf4, 0x16, 0x41, 0x8a, 0xed, 0xc7, 0xae, 0xb5, 0x8e, 0xa1,
	0x93, 0xe5, 0x85, 0x28, 0x95, 0x6d, 0x65, 0x7b, 0xe0, 0x0c, 0x7d, 0xda, 0x36, 0x18, 0xf6, 0x52,
	0x53, 0xf8, 0x97, 0x1a, 0xa5, 0x63, 0x28, 0x06, 0x33, 0x94, 0x17, 0x60, 0xff, 0x58, 0x5c, 0x95,
	0x22, 0x0f, 0xbb, 0x28, 0xbe, 0x6d, 0x7b, 0xa8, 0x71, 0x0a, 0x26, 0x7f, 0x5e, 0x8a, 0x9c, 0xf4,
	0xc1, 0x53, 0x2c, 0x0d, 0x7b, 0x58, 0x7e, 0x0b, 0x59, 0x73, 0x96, 0x52, 0x0d, 0x46, 0xe3, 0xfb,
	0x4d, 0xdb, 0x81, 0xd6, 0xc5, 0xf4, 0xc3, 0xd9, 0x62, 0x7a, 0xa9, 0xed, 0xfa, 0x19, 0x02, 0x73,
	0xc5, 0x5f, 0x5e, 0x4f, 0xed, 0x12, 0x81, 0xdb, 0xe9, 0xa1, 0x4b, 0x84, 0xe2, 0xd1, 0x37, 0x07,
	0xbc, 0x39, 0x4b, 0x1f, 0x71, 0xd3, 0x1d, 0xa7, 0x90, 0x67, 0xd5, 0x70, 0x7c, 0x1c, 0x4e, 0xaf,
	0x6a, 0xc0, 0xce, 0x60, 0xa2, 0xe3, 0x7b, 0x1b, 0x11, 0x7d, 0x77, 0xa0, 0x41, 0x35, 0xe5, 0x8e,
	0x7a, 0x72, 0x2e, 0x25, 0x4b, 0x2b, 0x93, 0x57, 0xa1, 0x66, 0x2e, 0xb7, 0x56, 0xaf, 0xbb, 0xdc,
	0x92, 0x7f, 0xc1, 0x97, 0xb1, 0x28, 0x4d, 0x2d, 0x2e, 0x35, 0x41, 0xdd, 0xc7, 0xc1, 0xc3, 0xdf,
	0x9b, 0x07, 0x54, 0xfc, 0xd5, 0x81, 0x9e, 0xb5, 0xa1, 0x3c, 0xe3, 0x8a, 0x65, 0x2b, 0xf2, 0x3f,
	0x80, 0x7d, 0xb3, 0x17, 0xd7, 0x1a, 0xf6, 0x2c, 0x32, 0x49, 0xc8, 0x21, 0xb4, 0xd0, 0xd3, 0x3a,
	0x69, 0xb5, 0x60, 0x3c, 0x49, 0x74, 0xed, 0xb1, 0xd8, 0xac, 0x15, 0xca, 0xf1, 0xa9, 0x09, 0xf4,
	0x96, 0xd6, 0xcc, 0x8b, 0x5b, 0xba, 0xeb, 0xe7, 0xe8, 0xa7, 0x03, 0xdd, 0x89, 0x8d, 0xef, 0x7e,
	0x21, 0xc6, 0xd0, 0xb2, 0x15, 0xc8, 0xd0, 0x45, 0x97, 0x1e, 0xd4, 0x37, 0xc8, 0x96, 0x4e, 0xaf,
	0x49, 0xf5, 0x86, 0x79, 0x8f, 0x7a, 0xc5, 0x94, 0x50, 0x6c, 0x65, 0xf7, 0xca, 0x94, 0x0a, 0x08,
	0x99, 0xb5, 0xea, 0x43, 0x2b, 0xc9, 0xa4, 0x11, 0x69, 0x06, 0x74, 0x1d, 0xeb, 0xc6, 0x94, 0x5c,
	0xff, 0xcb, 0xb6, 0x38, 0x24, 0x9f, 0x36, 0x75, 0x3c, 0x63, 0xdb, 0x65, 0x80, 0x97, 0xbe, 0xfa,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x78, 0x14, 0x0a, 0xca, 0x0d, 0x07, 0x00, 0x00,
}
