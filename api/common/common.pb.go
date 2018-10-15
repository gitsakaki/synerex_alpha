// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common // import "api/common"

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

type OtherTime int32

const (
	OtherTime_NONE       OtherTime = 0
	OtherTime_AS_SONN_AS OtherTime = 1
)

var OtherTime_name = map[int32]string{
	0: "NONE",
	1: "AS_SONN_AS",
}
var OtherTime_value = map[string]int32{
	"NONE":       0,
	"AS_SONN_AS": 1,
}

func (x OtherTime) String() string {
	return proto.EnumName(OtherTime_name, int32(x))
}
func (OtherTime) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_6fe0c46842ff11f1, []int{0}
}

type Place struct {
	// Types that are valid to be assigned to Value:
	//	*Place_Point
	//	*Place_Areas
	Value                isPlace_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Place) Reset()         { *m = Place{} }
func (m *Place) String() string { return proto.CompactTextString(m) }
func (*Place) ProtoMessage()    {}
func (*Place) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_6fe0c46842ff11f1, []int{0}
}
func (m *Place) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Place.Unmarshal(m, b)
}
func (m *Place) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Place.Marshal(b, m, deterministic)
}
func (dst *Place) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Place.Merge(dst, src)
}
func (m *Place) XXX_Size() int {
	return xxx_messageInfo_Place.Size(m)
}
func (m *Place) XXX_DiscardUnknown() {
	xxx_messageInfo_Place.DiscardUnknown(m)
}

var xxx_messageInfo_Place proto.InternalMessageInfo

type isPlace_Value interface {
	isPlace_Value()
}

type Place_Point struct {
	Point *Point `protobuf:"bytes,1,opt,name=point,proto3,oneof"`
}

type Place_Areas struct {
	Areas *Areas `protobuf:"bytes,2,opt,name=areas,proto3,oneof"`
}

func (*Place_Point) isPlace_Value() {}

func (*Place_Areas) isPlace_Value() {}

func (m *Place) GetValue() isPlace_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Place) GetPoint() *Point {
	if x, ok := m.GetValue().(*Place_Point); ok {
		return x.Point
	}
	return nil
}

func (m *Place) GetAreas() *Areas {
	if x, ok := m.GetValue().(*Place_Areas); ok {
		return x.Areas
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Place) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Place_OneofMarshaler, _Place_OneofUnmarshaler, _Place_OneofSizer, []interface{}{
		(*Place_Point)(nil),
		(*Place_Areas)(nil),
	}
}

func _Place_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Place)
	// value
	switch x := m.Value.(type) {
	case *Place_Point:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Point); err != nil {
			return err
		}
	case *Place_Areas:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Areas); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Place.Value has unexpected type %T", x)
	}
	return nil
}

func _Place_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Place)
	switch tag {
	case 1: // value.point
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Point)
		err := b.DecodeMessage(msg)
		m.Value = &Place_Point{msg}
		return true, err
	case 2: // value.areas
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Areas)
		err := b.DecodeMessage(msg)
		m.Value = &Place_Areas{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Place_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Place)
	// value
	switch x := m.Value.(type) {
	case *Place_Point:
		s := proto.Size(x.Point)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Place_Areas:
		s := proto.Size(x.Areas)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Point struct {
	Latitude             float64  `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_6fe0c46842ff11f1, []int{1}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (dst *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(dst, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Point) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Areas struct {
	Values               []*Area  `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Areas) Reset()         { *m = Areas{} }
func (m *Areas) String() string { return proto.CompactTextString(m) }
func (*Areas) ProtoMessage()    {}
func (*Areas) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_6fe0c46842ff11f1, []int{2}
}
func (m *Areas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Areas.Unmarshal(m, b)
}
func (m *Areas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Areas.Marshal(b, m, deterministic)
}
func (dst *Areas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Areas.Merge(dst, src)
}
func (m *Areas) XXX_Size() int {
	return xxx_messageInfo_Areas.Size(m)
}
func (m *Areas) XXX_DiscardUnknown() {
	xxx_messageInfo_Areas.DiscardUnknown(m)
}

var xxx_messageInfo_Areas proto.InternalMessageInfo

func (m *Areas) GetValues() []*Area {
	if m != nil {
		return m.Values
	}
	return nil
}

type Area struct {
	Points               []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Area) Reset()         { *m = Area{} }
func (m *Area) String() string { return proto.CompactTextString(m) }
func (*Area) ProtoMessage()    {}
func (*Area) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_6fe0c46842ff11f1, []int{3}
}
func (m *Area) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Area.Unmarshal(m, b)
}
func (m *Area) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Area.Marshal(b, m, deterministic)
}
func (dst *Area) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Area.Merge(dst, src)
}
func (m *Area) XXX_Size() int {
	return xxx_messageInfo_Area.Size(m)
}
func (m *Area) XXX_DiscardUnknown() {
	xxx_messageInfo_Area.DiscardUnknown(m)
}

var xxx_messageInfo_Area proto.InternalMessageInfo

func (m *Area) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type Time struct {
	// Types that are valid to be assigned to Value:
	//	*Time_Timestamp
	//	*Time_Periods
	//	*Time_Other
	Value                isTime_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_6fe0c46842ff11f1, []int{4}
}
func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}
func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}
func (dst *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(dst, src)
}
func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}
func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

type isTime_Value interface {
	isTime_Value()
}

type Time_Timestamp struct {
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3,oneof"`
}

type Time_Periods struct {
	Periods *Period `protobuf:"bytes,2,opt,name=periods,proto3,oneof"`
}

type Time_Other struct {
	Other OtherTime `protobuf:"varint,3,opt,name=other,proto3,enum=api.common.OtherTime,oneof"`
}

func (*Time_Timestamp) isTime_Value() {}

func (*Time_Periods) isTime_Value() {}

func (*Time_Other) isTime_Value() {}

func (m *Time) GetValue() isTime_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Time) GetTimestamp() *timestamp.Timestamp {
	if x, ok := m.GetValue().(*Time_Timestamp); ok {
		return x.Timestamp
	}
	return nil
}

func (m *Time) GetPeriods() *Period {
	if x, ok := m.GetValue().(*Time_Periods); ok {
		return x.Periods
	}
	return nil
}

func (m *Time) GetOther() OtherTime {
	if x, ok := m.GetValue().(*Time_Other); ok {
		return x.Other
	}
	return OtherTime_NONE
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Time) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Time_OneofMarshaler, _Time_OneofUnmarshaler, _Time_OneofSizer, []interface{}{
		(*Time_Timestamp)(nil),
		(*Time_Periods)(nil),
		(*Time_Other)(nil),
	}
}

func _Time_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Time)
	// value
	switch x := m.Value.(type) {
	case *Time_Timestamp:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Timestamp); err != nil {
			return err
		}
	case *Time_Periods:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Periods); err != nil {
			return err
		}
	case *Time_Other:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Other))
	case nil:
	default:
		return fmt.Errorf("Time.Value has unexpected type %T", x)
	}
	return nil
}

func _Time_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Time)
	switch tag {
	case 1: // value.timestamp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(timestamp.Timestamp)
		err := b.DecodeMessage(msg)
		m.Value = &Time_Timestamp{msg}
		return true, err
	case 2: // value.periods
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Period)
		err := b.DecodeMessage(msg)
		m.Value = &Time_Periods{msg}
		return true, err
	case 3: // value.other
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Time_Other{OtherTime(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Time_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Time)
	// value
	switch x := m.Value.(type) {
	case *Time_Timestamp:
		s := proto.Size(x.Timestamp)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Time_Periods:
		s := proto.Size(x.Periods)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Time_Other:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Other))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Periods struct {
	Values               []*Period `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Periods) Reset()         { *m = Periods{} }
func (m *Periods) String() string { return proto.CompactTextString(m) }
func (*Periods) ProtoMessage()    {}
func (*Periods) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_6fe0c46842ff11f1, []int{5}
}
func (m *Periods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Periods.Unmarshal(m, b)
}
func (m *Periods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Periods.Marshal(b, m, deterministic)
}
func (dst *Periods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Periods.Merge(dst, src)
}
func (m *Periods) XXX_Size() int {
	return xxx_messageInfo_Periods.Size(m)
}
func (m *Periods) XXX_DiscardUnknown() {
	xxx_messageInfo_Periods.DiscardUnknown(m)
}

var xxx_messageInfo_Periods proto.InternalMessageInfo

func (m *Periods) GetValues() []*Period {
	if m != nil {
		return m.Values
	}
	return nil
}

type Period struct {
	From                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Options              []*RepeatOption      `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_6fe0c46842ff11f1, []int{6}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Period.Unmarshal(m, b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Period.Marshal(b, m, deterministic)
}
func (dst *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(dst, src)
}
func (m *Period) XXX_Size() int {
	return xxx_messageInfo_Period.Size(m)
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetFrom() *timestamp.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Period) GetTo() *timestamp.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Period) GetOptions() []*RepeatOption {
	if m != nil {
		return m.Options
	}
	return nil
}

type RepeatOption struct {
	Weeks                []bool   `protobuf:"varint,1,rep,packed,name=weeks,proto3" json:"weeks,omitempty"`
	Weekdays             []bool   `protobuf:"varint,2,rep,packed,name=weekdays,proto3" json:"weekdays,omitempty"`
	Times                []uint32 `protobuf:"varint,3,rep,packed,name=times,proto3" json:"times,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepeatOption) Reset()         { *m = RepeatOption{} }
func (m *RepeatOption) String() string { return proto.CompactTextString(m) }
func (*RepeatOption) ProtoMessage()    {}
func (*RepeatOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_6fe0c46842ff11f1, []int{7}
}
func (m *RepeatOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatOption.Unmarshal(m, b)
}
func (m *RepeatOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatOption.Marshal(b, m, deterministic)
}
func (dst *RepeatOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatOption.Merge(dst, src)
}
func (m *RepeatOption) XXX_Size() int {
	return xxx_messageInfo_RepeatOption.Size(m)
}
func (m *RepeatOption) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatOption.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatOption proto.InternalMessageInfo

func (m *RepeatOption) GetWeeks() []bool {
	if m != nil {
		return m.Weeks
	}
	return nil
}

func (m *RepeatOption) GetWeekdays() []bool {
	if m != nil {
		return m.Weekdays
	}
	return nil
}

func (m *RepeatOption) GetTimes() []uint32 {
	if m != nil {
		return m.Times
	}
	return nil
}

func init() {
	proto.RegisterType((*Place)(nil), "api.common.Place")
	proto.RegisterType((*Point)(nil), "api.common.Point")
	proto.RegisterType((*Areas)(nil), "api.common.Areas")
	proto.RegisterType((*Area)(nil), "api.common.Area")
	proto.RegisterType((*Time)(nil), "api.common.Time")
	proto.RegisterType((*Periods)(nil), "api.common.Periods")
	proto.RegisterType((*Period)(nil), "api.common.Period")
	proto.RegisterType((*RepeatOption)(nil), "api.common.RepeatOption")
	proto.RegisterEnum("api.common.OtherTime", OtherTime_name, OtherTime_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_common_6fe0c46842ff11f1) }

var fileDescriptor_common_6fe0c46842ff11f1 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x3b, 0x6d, 0xd2, 0x3f, 0x77, 0xd7, 0xa5, 0x5e, 0x15, 0x42, 0x11, 0x2c, 0x01, 0xa1,
	0x5b, 0x70, 0xca, 0x56, 0x7c, 0xf1, 0xad, 0x0b, 0x42, 0x9f, 0xda, 0x32, 0x5d, 0x7d, 0xf0, 0x65,
	0x99, 0xdd, 0x4e, 0x6b, 0x30, 0xc9, 0x84, 0x64, 0xaa, 0xf8, 0x4d, 0xfc, 0x0c, 0x7e, 0x4a, 0x99,
	0x3b, 0xf9, 0x53, 0x77, 0x45, 0x9f, 0x9a, 0x3b, 0xe7, 0x97, 0x39, 0x73, 0x4e, 0x33, 0xf0, 0xec,
	0x5e, 0x27, 0x89, 0x4e, 0x67, 0xee, 0x87, 0x67, 0xb9, 0x36, 0x1a, 0x41, 0x66, 0x11, 0x77, 0x2b,
	0xa3, 0x57, 0x07, 0xad, 0x0f, 0xb1, 0x9a, 0x91, 0x72, 0x77, 0xdc, 0xcf, 0x4c, 0x94, 0xa8, 0xc2,
	0xc8, 0x24, 0x73, 0x70, 0xb8, 0x07, 0x7f, 0x13, 0xcb, 0x7b, 0x85, 0x97, 0xe0, 0x67, 0x3a, 0x4a,
	0x4d, 0xc0, 0xc6, 0x6c, 0x72, 0x36, 0x7f, 0xca, 0x9b, 0x5d, 0xf8, 0xc6, 0x0a, 0xcb, 0x96, 0x70,
	0x84, 0x45, 0x65, 0xae, 0x64, 0x11, 0xb4, 0x1f, 0xa3, 0x0b, 0x2b, 0x58, 0x94, 0x88, 0xeb, 0x1e,
	0xf8, 0xdf, 0x64, 0x7c, 0x54, 0xe1, 0x47, 0xf0, 0x69, 0x17, 0x1c, 0x41, 0x3f, 0x96, 0x26, 0x32,
	0xc7, 0x9d, 0x22, 0x2b, 0x26, 0xea, 0x19, 0x5f, 0xc2, 0x20, 0xd6, 0xe9, 0xc1, 0x89, 0x6d, 0x12,
	0x9b, 0x05, 0x44, 0xf0, 0x52, 0x99, 0xa8, 0xa0, 0x33, 0x66, 0x93, 0x81, 0xa0, 0xe7, 0xf0, 0x0a,
	0x7c, 0x72, 0xc4, 0x09, 0x74, 0xc9, 0xa8, 0x08, 0xd8, 0xb8, 0x33, 0x39, 0x9b, 0x0f, 0x1f, 0x1e,
	0x4a, 0x94, 0x7a, 0x78, 0x05, 0x9e, 0x9d, 0xf1, 0x12, 0xba, 0x14, 0xa7, 0x7a, 0xe3, 0x71, 0x62,
	0x51, 0x02, 0xe1, 0x2f, 0x06, 0xde, 0x4d, 0x94, 0x28, 0x7c, 0x0f, 0x83, 0xba, 0xc0, 0xb2, 0xa8,
	0x11, 0x77, 0x15, 0xf3, 0xaa, 0x62, 0x7e, 0x53, 0x11, 0xcb, 0x96, 0x68, 0x70, 0xe4, 0xd0, 0xcb,
	0x54, 0x1e, 0xe9, 0x5d, 0xd5, 0x1b, 0xfe, 0x61, 0x48, 0xd2, 0xb2, 0x25, 0x2a, 0x08, 0xdf, 0x80,
	0xaf, 0xcd, 0x17, 0x95, 0x53, 0xde, 0x8b, 0xf9, 0x8b, 0x53, 0x7a, 0x6d, 0x05, 0xeb, 0x63, 0x9b,
	0x26, 0xaa, 0x69, 0xfa, 0x1d, 0xf4, 0x36, 0xe5, 0x16, 0xd3, 0x07, 0xa5, 0xfc, 0xc5, 0xb1, 0xae,
	0xe5, 0x27, 0x83, 0xae, 0x5b, 0x42, 0x0e, 0xde, 0x3e, 0xd7, 0xc9, 0xff, 0x03, 0x0a, 0xe2, 0x70,
	0x0a, 0x6d, 0xa3, 0xcb, 0x50, 0xff, 0xa2, 0xdb, 0x46, 0xe3, 0x1c, 0x7a, 0x3a, 0x33, 0x91, 0x4e,
	0x8b, 0xa0, 0x43, 0x67, 0x0a, 0x4e, 0xcf, 0x24, 0x54, 0xa6, 0xa4, 0x59, 0x13, 0x20, 0x2a, 0x30,
	0xfc, 0x04, 0xe7, 0xa7, 0x02, 0x3e, 0x07, 0xff, 0xbb, 0x52, 0x5f, 0x5d, 0xaa, 0xbe, 0x70, 0x83,
	0xfd, 0xb0, 0xec, 0xc3, 0x4e, 0xfe, 0xb0, 0x05, 0x5b, 0xa1, 0x9e, 0xed, 0x1b, 0xf4, 0x47, 0x90,
	0xe7, 0x13, 0xe1, 0x86, 0xe9, 0x6b, 0x18, 0xd4, 0x45, 0x62, 0x1f, 0xbc, 0xd5, 0x7a, 0xf5, 0x61,
	0xd8, 0xc2, 0x0b, 0x80, 0xc5, 0xf6, 0x76, 0xbb, 0x5e, 0xad, 0x6e, 0x17, 0xdb, 0x21, 0xbb, 0x3e,
	0xff, 0x6c, 0x6f, 0x54, 0x79, 0xc7, 0xee, 0xba, 0x14, 0xec, 0xed, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x73, 0xcf, 0xfd, 0x7f, 0x7b, 0x03, 0x00, 0x00,
}
