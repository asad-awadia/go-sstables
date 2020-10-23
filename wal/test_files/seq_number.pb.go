// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wal/test_files/seq_number.proto

package test_files

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SequenceNumber struct {
	SequenceNumber       uint64   `protobuf:"varint,1,opt,name=sequenceNumber" json:"sequenceNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SequenceNumber) Reset()         { *m = SequenceNumber{} }
func (m *SequenceNumber) String() string { return proto.CompactTextString(m) }
func (*SequenceNumber) ProtoMessage()    {}
func (*SequenceNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_seq_number_a2750e045d4d51e1, []int{0}
}
func (m *SequenceNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SequenceNumber.Unmarshal(m, b)
}
func (m *SequenceNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SequenceNumber.Marshal(b, m, deterministic)
}
func (dst *SequenceNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SequenceNumber.Merge(dst, src)
}
func (m *SequenceNumber) XXX_Size() int {
	return xxx_messageInfo_SequenceNumber.Size(m)
}
func (m *SequenceNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_SequenceNumber.DiscardUnknown(m)
}

var xxx_messageInfo_SequenceNumber proto.InternalMessageInfo

func (m *SequenceNumber) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*SequenceNumber)(nil), "test_files.SequenceNumber")
}

func init() {
	proto.RegisterFile("wal/test_files/seq_number.proto", fileDescriptor_seq_number_a2750e045d4d51e1)
}

var fileDescriptor_seq_number_a2750e045d4d51e1 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x4f, 0xcc, 0xd1,
	0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x4f, 0xcb, 0xcc, 0x49, 0x2d, 0xd6, 0x2f, 0x4e, 0x2d, 0x8c, 0xcf,
	0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x48, 0x2a,
	0x59, 0x70, 0xf1, 0x05, 0xa7, 0x16, 0x96, 0xa6, 0xe6, 0x25, 0xa7, 0xfa, 0x81, 0xd5, 0x08, 0xa9,
	0x71, 0xf1, 0x15, 0xa3, 0x88, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0xa1, 0x89, 0x26, 0xb1,
	0x81, 0x0d, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x7a, 0x85, 0x0f, 0x6f, 0x00, 0x00,
	0x00,
}
