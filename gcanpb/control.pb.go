// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control.proto

package gcanpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateTopicRequest struct {
	Topic      string `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Partitions uint32 `protobuf:"varint,2,opt,name=Partitions,proto3" json:"Partitions,omitempty"`
	Replicas   uint32 `protobuf:"varint,3,opt,name=Replicas,proto3" json:"Replicas,omitempty"`
	// Optional
	ReplicaAssignment    *TopicAssignment `protobuf:"bytes,4,opt,name=ReplicaAssignment,proto3" json:"ReplicaAssignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateTopicRequest) Reset()         { *m = CreateTopicRequest{} }
func (m *CreateTopicRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTopicRequest) ProtoMessage()    {}
func (*CreateTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_control_23b3926d92f35cc6, []int{0}
}
func (m *CreateTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicRequest.Unmarshal(m, b)
}
func (m *CreateTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicRequest.Merge(dst, src)
}
func (m *CreateTopicRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTopicRequest.Size(m)
}
func (m *CreateTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicRequest proto.InternalMessageInfo

func (m *CreateTopicRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *CreateTopicRequest) GetPartitions() uint32 {
	if m != nil {
		return m.Partitions
	}
	return 0
}

func (m *CreateTopicRequest) GetReplicas() uint32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *CreateTopicRequest) GetReplicaAssignment() *TopicAssignment {
	if m != nil {
		return m.ReplicaAssignment
	}
	return nil
}

type CreateTopicResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTopicResponse) Reset()         { *m = CreateTopicResponse{} }
func (m *CreateTopicResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTopicResponse) ProtoMessage()    {}
func (*CreateTopicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_control_23b3926d92f35cc6, []int{1}
}
func (m *CreateTopicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicResponse.Unmarshal(m, b)
}
func (m *CreateTopicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTopicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicResponse.Merge(dst, src)
}
func (m *CreateTopicResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTopicResponse.Size(m)
}
func (m *CreateTopicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateTopicRequest)(nil), "gcanpb.CreateTopicRequest")
	proto.RegisterType((*CreateTopicResponse)(nil), "gcanpb.CreateTopicResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopicClient is the client API for Topic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicClient interface {
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error)
}

type topicClient struct {
	cc *grpc.ClientConn
}

func NewTopicClient(cc *grpc.ClientConn) TopicClient {
	return &topicClient{cc}
}

func (c *topicClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*CreateTopicResponse, error) {
	out := new(CreateTopicResponse)
	err := c.cc.Invoke(ctx, "/gcanpb.Topic/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServer is the server API for Topic service.
type TopicServer interface {
	CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error)
}

func RegisterTopicServer(s *grpc.Server, srv TopicServer) {
	s.RegisterService(&_Topic_serviceDesc, srv)
}

func _Topic_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gcanpb.Topic/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Topic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gcanpb.Topic",
	HandlerType: (*TopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _Topic_CreateTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}

func init() { proto.RegisterFile("control.proto", fileDescriptor_control_23b3926d92f35cc6) }

var fileDescriptor_control_23b3926d92f35cc6 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4a, 0xc6, 0x30,
	0x10, 0xc5, 0x8d, 0x7f, 0x3e, 0x74, 0x3e, 0xbb, 0x70, 0x54, 0x0c, 0x11, 0xa4, 0x74, 0xd5, 0x55,
	0x17, 0xf5, 0x04, 0x22, 0x82, 0x4b, 0x0d, 0x5e, 0x20, 0x0d, 0xa1, 0x04, 0xda, 0x4c, 0x4c, 0xe2,
	0xb1, 0xbc, 0xa3, 0x90, 0x54, 0xad, 0xd4, 0x5d, 0xde, 0xfb, 0x91, 0x37, 0x6f, 0x06, 0x2a, 0x4d,
	0x2e, 0x05, 0x9a, 0x3a, 0x1f, 0x28, 0x11, 0xee, 0x46, 0xad, 0x9c, 0x1f, 0xc4, 0xb9, 0xa6, 0x79,
	0x26, 0x57, 0xdc, 0xe6, 0x93, 0x01, 0x3e, 0x06, 0xa3, 0x92, 0x79, 0x23, 0x6f, 0xb5, 0x34, 0xef,
	0x1f, 0x26, 0x26, 0xbc, 0x82, 0x93, 0xac, 0x39, 0xab, 0x59, 0x7b, 0x26, 0x8b, 0xc0, 0x3b, 0x80,
	0x17, 0x15, 0x92, 0x4d, 0x96, 0x5c, 0xe4, 0x87, 0x35, 0x6b, 0x2b, 0xb9, 0x72, 0x50, 0xc0, 0xa9,
	0x34, 0x7e, 0xb2, 0x5a, 0x45, 0x7e, 0x94, 0xe9, 0x8f, 0xc6, 0x27, 0xb8, 0x58, 0xde, 0x0f, 0x31,
	0xda, 0xd1, 0xcd, 0xc6, 0x25, 0x7e, 0x5c, 0xb3, 0x76, 0xdf, 0xdf, 0x74, 0xa5, 0x5a, 0x97, 0xa7,
	0xfc, 0x62, 0xb9, 0xfd, 0xd1, 0x5c, 0xc3, 0xe5, 0x9f, 0xba, 0xd1, 0x93, 0x8b, 0xa6, 0x7f, 0x5d,
	0xfa, 0xe2, 0x33, 0xec, 0x57, 0x1c, 0xc5, 0x77, 0xf4, 0x76, 0x47, 0x71, 0xfb, 0x2f, 0x2b, 0x81,
	0xcd, 0xc1, 0xb0, 0xcb, 0x07, 0xba, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x79, 0x4e, 0x2d, 0x1e,
	0x47, 0x01, 0x00, 0x00,
}