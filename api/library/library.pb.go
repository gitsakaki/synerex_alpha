// Code generated by protoc-gen-go. DO NOT EDIT.
// source: library/library.proto

package library

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Action int32

const (
	Action_NONE            Action = 0
	Action_LEND_REQUEST    Action = 1
	Action_RETURN_REQUEST  Action = 2
	Action_MODULE_PREPARE  Action = 3
	Action_MODULE_SHIPPING Action = 4
)

var Action_name = map[int32]string{
	0: "NONE",
	1: "LEND_REQUEST",
	2: "RETURN_REQUEST",
	3: "MODULE_PREPARE",
	4: "MODULE_SHIPPING",
}

var Action_value = map[string]int32{
	"NONE":            0,
	"LEND_REQUEST":    1,
	"RETURN_REQUEST":  2,
	"MODULE_PREPARE":  3,
	"MODULE_SHIPPING": 4,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9a985699bc2ca255, []int{0}
}

type LibService struct {
	Action               Action               `protobuf:"varint,1,opt,name=action,proto3,enum=api.library.Action" json:"action,omitempty"`
	Books                []*Book              `protobuf:"bytes,2,rep,name=books,proto3" json:"books,omitempty"`
	LendDate             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=lend_date,json=lendDate,proto3" json:"lend_date,omitempty"`
	ReturnDate           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LibService) Reset()         { *m = LibService{} }
func (m *LibService) String() string { return proto.CompactTextString(m) }
func (*LibService) ProtoMessage()    {}
func (*LibService) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a985699bc2ca255, []int{0}
}

func (m *LibService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LibService.Unmarshal(m, b)
}
func (m *LibService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LibService.Marshal(b, m, deterministic)
}
func (m *LibService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LibService.Merge(m, src)
}
func (m *LibService) XXX_Size() int {
	return xxx_messageInfo_LibService.Size(m)
}
func (m *LibService) XXX_DiscardUnknown() {
	xxx_messageInfo_LibService.DiscardUnknown(m)
}

var xxx_messageInfo_LibService proto.InternalMessageInfo

func (m *LibService) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_NONE
}

func (m *LibService) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

func (m *LibService) GetLendDate() *timestamp.Timestamp {
	if m != nil {
		return m.LendDate
	}
	return nil
}

func (m *LibService) GetReturnDate() *timestamp.Timestamp {
	if m != nil {
		return m.ReturnDate
	}
	return nil
}

type Book struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Volume               string   `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Authors              []string `protobuf:"bytes,3,rep,name=authors,proto3" json:"authors,omitempty"`
	Publisher            string   `protobuf:"bytes,4,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Isbn                 string   `protobuf:"bytes,5,opt,name=isbn,proto3" json:"isbn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a985699bc2ca255, []int{1}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *Book) GetAuthors() []string {
	if m != nil {
		return m.Authors
	}
	return nil
}

func (m *Book) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *Book) GetIsbn() string {
	if m != nil {
		return m.Isbn
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.library.Action", Action_name, Action_value)
	proto.RegisterType((*LibService)(nil), "api.library.LibService")
	proto.RegisterType((*Book)(nil), "api.library.Book")
}

func init() { proto.RegisterFile("library/library.proto", fileDescriptor_9a985699bc2ca255) }

var fileDescriptor_9a985699bc2ca255 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x49, 0x93, 0x96, 0xcd, 0x14, 0x2d, 0x61, 0x16, 0x50, 0xb4, 0x42, 0xa2, 0xda, 0x0b,
	0x15, 0x20, 0x47, 0x2a, 0x07, 0x0e, 0x9c, 0x76, 0x55, 0x0b, 0x56, 0x2a, 0xd9, 0xe2, 0xb6, 0x17,
	0x2e, 0x95, 0xdd, 0x35, 0xad, 0xb5, 0x49, 0x1c, 0x39, 0xce, 0x8a, 0xde, 0xf8, 0x98, 0x7c, 0x1c,
	0x14, 0x27, 0xe1, 0xcf, 0x89, 0xd3, 0x78, 0x7e, 0xef, 0x8d, 0xf5, 0x3c, 0x32, 0x3c, 0xcb, 0x94,
	0x30, 0xdc, 0x1c, 0x93, 0xae, 0x92, 0xd2, 0x68, 0xab, 0x71, 0xcc, 0x4b, 0x45, 0x3a, 0x74, 0xfe,
	0x72, 0xaf, 0xf5, 0x3e, 0x93, 0x89, 0x93, 0x44, 0xfd, 0x2d, 0xb1, 0x2a, 0x97, 0x95, 0xe5, 0x79,
	0xd9, 0xba, 0x2f, 0x7e, 0x7a, 0x00, 0x0b, 0x25, 0x56, 0xd2, 0xdc, 0xab, 0x9d, 0xc4, 0x37, 0x30,
	0xe2, 0x3b, 0xab, 0x74, 0x11, 0x7b, 0x13, 0x6f, 0x7a, 0x3a, 0x3b, 0x23, 0x7f, 0xdd, 0x46, 0x2e,
	0x9d, 0xc4, 0x3a, 0x0b, 0xbe, 0x82, 0xa1, 0xd0, 0xfa, 0xae, 0x8a, 0x07, 0x13, 0x7f, 0x3a, 0x9e,
	0x3d, 0xf9, 0xc7, 0x7b, 0xa5, 0xf5, 0x1d, 0x6b, 0x75, 0x7c, 0x0f, 0x61, 0x26, 0x8b, 0xdb, 0xed,
	0x2d, 0xb7, 0x32, 0xf6, 0x27, 0xde, 0x74, 0x3c, 0x3b, 0x27, 0x6d, 0x32, 0xd2, 0x27, 0x23, 0xeb,
	0x3e, 0x19, 0x3b, 0x69, 0xcc, 0x73, 0x6e, 0x25, 0x7e, 0x80, 0xb1, 0x91, 0xb6, 0x36, 0x45, 0x3b,
	0x1a, 0xfc, 0x77, 0x14, 0x5a, 0x7b, 0x33, 0x7c, 0xf1, 0xc3, 0x83, 0xa0, 0x49, 0x81, 0x4f, 0x61,
	0x68, 0x95, 0xcd, 0xa4, 0x7b, 0x53, 0xc8, 0xda, 0x06, 0x9f, 0xc3, 0xe8, 0x5e, 0x67, 0x75, 0x2e,
	0xe3, 0x81, 0xc3, 0x5d, 0x87, 0x31, 0x3c, 0xe4, 0xb5, 0x3d, 0x68, 0x53, 0xc5, 0xfe, 0xc4, 0x9f,
	0x86, 0xac, 0x6f, 0xf1, 0x05, 0x84, 0x65, 0x2d, 0x32, 0x55, 0x1d, 0xa4, 0x71, 0x59, 0x42, 0xf6,
	0x07, 0x20, 0x42, 0xa0, 0x2a, 0x51, 0xc4, 0x43, 0x27, 0xb8, 0xf3, 0x6b, 0x0e, 0xa3, 0x76, 0x67,
	0x78, 0x02, 0x41, 0x7a, 0x93, 0xd2, 0xe8, 0x01, 0x46, 0xf0, 0x68, 0x41, 0xd3, 0xf9, 0x96, 0xd1,
	0x2f, 0x1b, 0xba, 0x5a, 0x47, 0x1e, 0x22, 0x9c, 0x32, 0xba, 0xde, 0xb0, 0xf4, 0x37, 0x1b, 0x34,
	0xec, 0xf3, 0xcd, 0x7c, 0xb3, 0xa0, 0xdb, 0x25, 0xa3, 0xcb, 0x4b, 0x46, 0x23, 0x1f, 0xcf, 0xe0,
	0x71, 0xc7, 0x56, 0x9f, 0xae, 0x97, 0xcb, 0xeb, 0xf4, 0x63, 0x14, 0x5c, 0x91, 0xaf, 0x6f, 0xf7,
	0xca, 0x1e, 0x6a, 0x41, 0x76, 0x3a, 0x4f, 0xaa, 0x63, 0x21, 0x8d, 0xfc, 0xde, 0xd7, 0x2d, 0xcf,
	0xca, 0x03, 0x4f, 0x78, 0xa9, 0xfa, 0x4f, 0x22, 0x46, 0x6e, 0x6b, 0xef, 0x7e, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x44, 0xce, 0x9c, 0xc3, 0x3e, 0x02, 0x00, 0x00,
}
