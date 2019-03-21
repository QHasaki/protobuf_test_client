// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/status.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Code int32

const (
	Code_OK             Code = 0
	Code_INTERNAL       Code = 1
	Code_DATA_LOSE      Code = 2
	Code_NO_TOKEN       Code = 3
	Code_UNAUTHORIZATED Code = 4
)

var Code_name = map[int32]string{
	0: "OK",
	1: "INTERNAL",
	2: "DATA_LOSE",
	3: "NO_TOKEN",
	4: "UNAUTHORIZATED",
}

var Code_value = map[string]int32{
	"OK":             0,
	"INTERNAL":       1,
	"DATA_LOSE":      2,
	"NO_TOKEN":       3,
	"UNAUTHORIZATED": 4,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17c96f8603cffea9, []int{0}
}

type Status struct {
	Code                 Code     `protobuf:"varint,1,opt,name=code,proto3,enum=protocol.Code" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_17c96f8603cffea9, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("protocol.Code", Code_name, Code_value)
	proto.RegisterType((*Status)(nil), "protocol.Status")
}

func init() { proto.RegisterFile("proto/status.proto", fileDescriptor_17c96f8603cffea9) }

var fileDescriptor_17c96f8603cffea9 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0x31, 0x4f, 0x87, 0x30,
	0x10, 0x47, 0x2d, 0x12, 0xfc, 0x73, 0x0a, 0x69, 0x6e, 0x62, 0x24, 0x4c, 0xc4, 0x01, 0x13, 0x1d,
	0x9d, 0xaa, 0xd4, 0x48, 0xc0, 0x36, 0x29, 0x65, 0x61, 0x21, 0x08, 0x0d, 0x8b, 0xa6, 0xc6, 0xe2,
	0xf7, 0x37, 0x34, 0x61, 0xba, 0x7b, 0x79, 0x6f, 0xf8, 0x01, 0xfe, 0xfc, 0xda, 0xdd, 0x3e, 0xb8,
	0x7d, 0xde, 0xff, 0x5c, 0xe5, 0x01, 0x2f, 0xfe, 0x2c, 0xf6, 0xab, 0x78, 0x83, 0xa8, 0xf7, 0x06,
	0x0b, 0x08, 0x17, 0xbb, 0x9a, 0x8c, 0xe4, 0xa4, 0x4c, 0x1f, 0xd3, 0xea, 0x4c, 0xaa, 0x57, 0xbb,
	0x1a, 0xe5, 0x1d, 0x66, 0x70, 0xf3, 0x6d, 0x9c, 0x9b, 0x37, 0x93, 0x05, 0x39, 0x29, 0x63, 0x75,
	0xe2, 0xfd, 0x07, 0x84, 0x47, 0x87, 0x11, 0x04, 0xb2, 0xa5, 0x57, 0x78, 0x07, 0x97, 0x46, 0x68,
	0xae, 0x04, 0xeb, 0x28, 0xc1, 0x04, 0xe2, 0x9a, 0x69, 0x36, 0x75, 0xb2, 0xe7, 0x34, 0x38, 0xa4,
	0x90, 0x93, 0x96, 0x2d, 0x17, 0xf4, 0x1a, 0x11, 0xd2, 0x41, 0xb0, 0x41, 0xbf, 0x4b, 0xd5, 0x8c,
	0x4c, 0xf3, 0x9a, 0x86, 0x2f, 0xc9, 0x78, 0xbb, 0xd9, 0xe7, 0x73, 0xc2, 0x67, 0xe4, 0xbf, 0xa7,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xfa, 0x05, 0xa7, 0xcc, 0x00, 0x00, 0x00,
}
