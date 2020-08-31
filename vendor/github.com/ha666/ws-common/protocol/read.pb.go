// Code generated by protoc-gen-go. DO NOT EDIT.
// source: read.proto

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

type Read struct {
	ReadVal              string   `protobuf:"bytes,1,opt,name=read_val,json=readVal,proto3" json:"read_val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Read) Reset()         { *m = Read{} }
func (m *Read) String() string { return proto.CompactTextString(m) }
func (*Read) ProtoMessage()    {}
func (*Read) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{0}
}

func (m *Read) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Read.Unmarshal(m, b)
}
func (m *Read) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Read.Marshal(b, m, deterministic)
}
func (m *Read) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Read.Merge(m, src)
}
func (m *Read) XXX_Size() int {
	return xxx_messageInfo_Read.Size(m)
}
func (m *Read) XXX_DiscardUnknown() {
	xxx_messageInfo_Read.DiscardUnknown(m)
}

var xxx_messageInfo_Read proto.InternalMessageInfo

func (m *Read) GetReadVal() string {
	if m != nil {
		return m.ReadVal
	}
	return ""
}

func init() {
	proto.RegisterType((*Read)(nil), "protocol.Read")
}

func init() { proto.RegisterFile("read.proto", fileDescriptor_7b10ec61df6818dd) }

var fileDescriptor_7b10ec61df6818dd = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4a, 0x4d, 0x4c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x8a, 0x5c,
	0x2c, 0x41, 0xa9, 0x89, 0x29, 0x42, 0x92, 0x5c, 0x1c, 0x20, 0xf9, 0xf8, 0xb2, 0xc4, 0x1c, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x76, 0x10, 0x3f, 0x2c, 0x31, 0x27, 0x89, 0x0d, 0xac, 0xd8,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x93, 0xe1, 0x02, 0x41, 0x00, 0x00, 0x00,
}
