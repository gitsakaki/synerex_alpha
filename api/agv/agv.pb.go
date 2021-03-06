// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agv/agv.proto

package agv

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	common "github.com/synerex/synerex_alpha/api/common"
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

type AgvService struct {
	OperatorId           int32              `protobuf:"varint,1,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	DepartPlace          *common.Place      `protobuf:"bytes,4,opt,name=depart_place,json=departPlace,proto3" json:"depart_place,omitempty"`
	ArrivePlace          *common.Place      `protobuf:"bytes,5,opt,name=arrive_place,json=arrivePlace,proto3" json:"arrive_place,omitempty"`
	DepartTime           *common.Time       `protobuf:"bytes,6,opt,name=depart_time,json=departTime,proto3" json:"depart_time,omitempty"`
	ArriveTime           *common.Time       `protobuf:"bytes,7,opt,name=arrive_time,json=arriveTime,proto3" json:"arrive_time,omitempty"`
	AmountTime           *duration.Duration `protobuf:"bytes,8,opt,name=amount_time,json=amountTime,proto3" json:"amount_time,omitempty"`
	DelayTime            *duration.Duration `protobuf:"bytes,13,opt,name=delay_time,json=delayTime,proto3" json:"delay_time,omitempty"`
	Points               []*common.Point    `protobuf:"bytes,15,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AgvService) Reset()         { *m = AgvService{} }
func (m *AgvService) String() string { return proto.CompactTextString(m) }
func (*AgvService) ProtoMessage()    {}
func (*AgvService) Descriptor() ([]byte, []int) {
	return fileDescriptor_51c30a9f8766ea5c, []int{0}
}

func (m *AgvService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgvService.Unmarshal(m, b)
}
func (m *AgvService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgvService.Marshal(b, m, deterministic)
}
func (m *AgvService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgvService.Merge(m, src)
}
func (m *AgvService) XXX_Size() int {
	return xxx_messageInfo_AgvService.Size(m)
}
func (m *AgvService) XXX_DiscardUnknown() {
	xxx_messageInfo_AgvService.DiscardUnknown(m)
}

var xxx_messageInfo_AgvService proto.InternalMessageInfo

func (m *AgvService) GetOperatorId() int32 {
	if m != nil {
		return m.OperatorId
	}
	return 0
}

func (m *AgvService) GetDepartPlace() *common.Place {
	if m != nil {
		return m.DepartPlace
	}
	return nil
}

func (m *AgvService) GetArrivePlace() *common.Place {
	if m != nil {
		return m.ArrivePlace
	}
	return nil
}

func (m *AgvService) GetDepartTime() *common.Time {
	if m != nil {
		return m.DepartTime
	}
	return nil
}

func (m *AgvService) GetArriveTime() *common.Time {
	if m != nil {
		return m.ArriveTime
	}
	return nil
}

func (m *AgvService) GetAmountTime() *duration.Duration {
	if m != nil {
		return m.AmountTime
	}
	return nil
}

func (m *AgvService) GetDelayTime() *duration.Duration {
	if m != nil {
		return m.DelayTime
	}
	return nil
}

func (m *AgvService) GetPoints() []*common.Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*AgvService)(nil), "api.agv.AgvService")
}

func init() { proto.RegisterFile("agv/agv.proto", fileDescriptor_51c30a9f8766ea5c) }

var fileDescriptor_51c30a9f8766ea5c = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x29, 0xb5, 0xad, 0x4e, 0x2c, 0x6a, 0xbd, 0xc4, 0x1e, 0xb4, 0x78, 0xaa, 0x97, 0x2c,
	0xd6, 0x1e, 0xc4, 0x9b, 0xe2, 0xc5, 0x9b, 0x44, 0x4f, 0x5e, 0xc2, 0xb4, 0x59, 0xb7, 0x43, 0x93,
	0xec, 0xb2, 0xdd, 0x2c, 0xf6, 0x6d, 0x7c, 0x54, 0xd9, 0x3f, 0x05, 0x41, 0x8a, 0xa7, 0x49, 0x66,
	0xbe, 0x6f, 0x26, 0xfc, 0x08, 0x0c, 0x51, 0x58, 0x86, 0xc2, 0x66, 0x4a, 0x4b, 0x23, 0x47, 0x03,
	0x54, 0x94, 0xa1, 0xb0, 0xe3, 0x4b, 0x21, 0xa5, 0xa8, 0x38, 0xf3, 0xed, 0x45, 0xfb, 0xc9, 0xca,
	0x56, 0xa3, 0x21, 0xd9, 0x04, 0x70, 0x7c, 0xbe, 0x94, 0x75, 0x2d, 0x1b, 0x16, 0x4a, 0x68, 0x5e,
	0x7f, 0x77, 0x01, 0x1e, 0x85, 0x7d, 0xe3, 0xda, 0xd2, 0x92, 0x8f, 0xae, 0x20, 0x91, 0x8a, 0x6b,
	0x34, 0x52, 0x17, 0x54, 0xa6, 0x9d, 0x49, 0x67, 0xda, 0xcb, 0x61, 0xd7, 0x7a, 0x29, 0x47, 0x73,
	0x38, 0x2e, 0xb9, 0x42, 0x6d, 0x0a, 0x55, 0xe1, 0x92, 0xa7, 0x07, 0x93, 0xce, 0x34, 0x99, 0x9d,
	0x65, 0xee, 0x23, 0xe2, 0xe2, 0x57, 0x37, 0xc8, 0x93, 0x80, 0xf9, 0x17, 0x67, 0xa1, 0xd6, 0x64,
	0x79, 0xb4, 0x7a, 0x7b, 0xad, 0x80, 0x05, 0xeb, 0x16, 0xe2, 0x92, 0xc2, 0x50, 0xcd, 0xd3, 0xbe,
	0x97, 0x4e, 0x7f, 0x4b, 0xef, 0x54, 0xf3, 0x1c, 0x02, 0xe4, 0x9e, 0x9d, 0x12, 0x0f, 0x79, 0x65,
	0xb0, 0x4f, 0x09, 0x90, 0x57, 0x1e, 0x20, 0xc1, 0x5a, 0xb6, 0x4d, 0xbc, 0x72, 0xe8, 0x95, 0x8b,
	0x2c, 0x84, 0x99, 0xed, 0xc2, 0xcc, 0x9e, 0x63, 0x98, 0x39, 0x04, 0xda, 0xbb, 0xf7, 0x00, 0x25,
	0xaf, 0x70, 0x1b, 0xd4, 0xe1, 0x7f, 0xea, 0x91, 0x87, 0xbd, 0x79, 0x03, 0x7d, 0x25, 0xa9, 0x31,
	0x9b, 0xf4, 0x64, 0xd2, 0xfd, 0x93, 0x85, 0x9b, 0xe4, 0x11, 0x78, 0x9a, 0x7f, 0xcc, 0x04, 0x99,
	0x55, 0xbb, 0x70, 0x63, 0x26, 0xc8, 0x6c, 0x70, 0x8d, 0x6b, 0x62, 0x9b, 0x6d, 0xc3, 0x35, 0xff,
	0xda, 0xd5, 0x02, 0x2b, 0xb5, 0x42, 0x86, 0x8a, 0xdc, 0xcf, 0xb1, 0xe8, 0xfb, 0xf3, 0x77, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x96, 0xba, 0x95, 0x2e, 0x2e, 0x02, 0x00, 0x00,
}
