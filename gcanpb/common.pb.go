// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package gcanpb

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

type TopicAssignment struct {
	PartitionAssignment  []*PartitionAssignment `protobuf:"bytes,1,rep,name=PartitionAssignment,proto3" json:"PartitionAssignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TopicAssignment) Reset()         { *m = TopicAssignment{} }
func (m *TopicAssignment) String() string { return proto.CompactTextString(m) }
func (*TopicAssignment) ProtoMessage()    {}
func (*TopicAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_f68ffb51160fb1d4, []int{0}
}
func (m *TopicAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicAssignment.Unmarshal(m, b)
}
func (m *TopicAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicAssignment.Marshal(b, m, deterministic)
}
func (dst *TopicAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicAssignment.Merge(dst, src)
}
func (m *TopicAssignment) XXX_Size() int {
	return xxx_messageInfo_TopicAssignment.Size(m)
}
func (m *TopicAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_TopicAssignment proto.InternalMessageInfo

func (m *TopicAssignment) GetPartitionAssignment() []*PartitionAssignment {
	if m != nil {
		return m.PartitionAssignment
	}
	return nil
}

type PartitionAssignment struct {
	Replicas             []uint32 `protobuf:"varint,1,rep,packed,name=Replicas,proto3" json:"Replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionAssignment) Reset()         { *m = PartitionAssignment{} }
func (m *PartitionAssignment) String() string { return proto.CompactTextString(m) }
func (*PartitionAssignment) ProtoMessage()    {}
func (*PartitionAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_f68ffb51160fb1d4, []int{1}
}
func (m *PartitionAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionAssignment.Unmarshal(m, b)
}
func (m *PartitionAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionAssignment.Marshal(b, m, deterministic)
}
func (dst *PartitionAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionAssignment.Merge(dst, src)
}
func (m *PartitionAssignment) XXX_Size() int {
	return xxx_messageInfo_PartitionAssignment.Size(m)
}
func (m *PartitionAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionAssignment proto.InternalMessageInfo

func (m *PartitionAssignment) GetReplicas() []uint32 {
	if m != nil {
		return m.Replicas
	}
	return nil
}

func init() {
	proto.RegisterType((*TopicAssignment)(nil), "gcanpb.TopicAssignment")
	proto.RegisterType((*PartitionAssignment)(nil), "gcanpb.PartitionAssignment")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_f68ffb51160fb1d4) }

var fileDescriptor_common_f68ffb51160fb1d4 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x4f, 0x4e, 0xcc, 0x2b, 0x48,
	0x52, 0x4a, 0xe0, 0xe2, 0x0f, 0xc9, 0x2f, 0xc8, 0x4c, 0x76, 0x2c, 0x2e, 0xce, 0x4c, 0xcf, 0xcb,
	0x4d, 0xcd, 0x2b, 0x11, 0xf2, 0xe5, 0x12, 0x0e, 0x48, 0x2c, 0x2a, 0xc9, 0x2c, 0xc9, 0xcc, 0xcf,
	0x43, 0x08, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x49, 0xeb, 0x41, 0x34, 0xea, 0x61, 0x51,
	0x12, 0x84, 0x4d, 0x9f, 0x92, 0x21, 0x56, 0xe3, 0x84, 0xa4, 0xb8, 0x38, 0x82, 0x52, 0x0b, 0x72,
	0x32, 0x93, 0x13, 0x8b, 0xc1, 0x46, 0xf3, 0x06, 0xc1, 0xf9, 0x49, 0x6c, 0x60, 0x37, 0x1a, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x24, 0x6f, 0x6a, 0xe2, 0xb3, 0x00, 0x00, 0x00,
}