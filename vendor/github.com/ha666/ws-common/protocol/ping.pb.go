// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping.proto

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

type Ping struct {
	PingVal              string   `protobuf:"bytes,1,opt,name=ping_val,json=pingVal,proto3" json:"ping_val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetPingVal() string {
	if m != nil {
		return m.PingVal
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "protocol.Ping")
}

func init() { proto.RegisterFile("ping.proto", fileDescriptor_6d51d96c3ad891f5) }

var fileDescriptor_6d51d96c3ad891f5 = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0xcc, 0x4b,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x8a, 0x5c,
	0x2c, 0x01, 0x99, 0x79, 0xe9, 0x42, 0x92, 0x5c, 0x1c, 0x20, 0xf9, 0xf8, 0xb2, 0xc4, 0x1c, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x76, 0x10, 0x3f, 0x2c, 0x31, 0x27, 0x89, 0x0d, 0xac, 0xd8,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xc8, 0x19, 0x75, 0x41, 0x00, 0x00, 0x00,
}
