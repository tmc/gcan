
# gcanpb
    import "github.com/tmc/gcan/gcanpb"

Package gcanpb is a generated protocol buffer package.

It is generated from these files:


	gcan.proto

It has these top-level messages:


	MessageSet
	Message
	SendRequest
	SendResponse
	ReceiveRequest
	ReceiveResponse





## Variables
``` go
var Err_name = map[int32]string{
    0: "Unknown",
    1: "CRCChecksumFailed",
}
```
``` go
var Err_value = map[string]int32{
    "Unknown":           0,
    "CRCChecksumFailed": 1,
}
```
``` go
var MessageCompression_name = map[int32]string{
    0: "NONE",
    1: "ZLIB",
    2: "SNAPPY",
}
```
``` go
var MessageCompression_value = map[string]int32{
    "NONE":   0,
    "ZLIB":   1,
    "SNAPPY": 2,
}
```

## func NopWCloser
``` go
func NopWCloser(w io.Writer) io.WriteCloser
```
NopWCloser returns a WriteCloser with a no-op Close method wrapping
the provided Writer w.


## func RegisterConsumerServer
``` go
func RegisterConsumerServer(s *grpc.Server, srv ConsumerServer)
```

## func RegisterProducerServer
``` go
func RegisterProducerServer(s *grpc.Server, srv ProducerServer)
```


## type ConsumerClient
``` go
type ConsumerClient interface {
    Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (Consumer_ReceiveClient, error)
}
```








### func NewConsumerClient
``` go
func NewConsumerClient(cc *grpc.ClientConn) ConsumerClient
```



## type ConsumerServer
``` go
type ConsumerServer interface {
    Receive(*ReceiveRequest, Consumer_ReceiveServer) error
}
```










## type Consumer_ReceiveClient
``` go
type Consumer_ReceiveClient interface {
    Recv() (*ReceiveResponse, error)
    grpc.ClientStream
}
```










## type Consumer_ReceiveServer
``` go
type Consumer_ReceiveServer interface {
    Send(*ReceiveResponse) error
    grpc.ServerStream
}
```










## type Err
``` go
type Err int32
```


``` go
const (
    Err_Unknown           Err = 0
    Err_CRCChecksumFailed Err = 1
)
```








### func (Err) EnumDescriptor
``` go
func (Err) EnumDescriptor() ([]byte, []int)
```


### func (Err) Error
``` go
func (x Err) Error() string
```


### func (Err) String
``` go
func (x Err) String() string
```


## type Message
``` go
type Message struct {
    // CRC32 of all following fields
    CRC         uint32             `protobuf:"varint,1,opt,name=CRC" json:"CRC,omitempty"`
    Compression MessageCompression `protobuf:"varint,2,opt,name=Compression,enum=gcanpb.MessageCompression" json:"Compression,omitempty"`
    Key         []byte             `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
    Value       []byte             `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
}
```
Message is an individual key-value pair.











### func (\*Message) CheckIntegrity
``` go
func (m *Message) CheckIntegrity() error
```


### func (\*Message) ComputeCRC
``` go
func (m *Message) ComputeCRC() uint32
```


### func (\*Message) Decompress
``` go
func (m *Message) Decompress() (io.ReadCloser, error)
```


### func (\*Message) Descriptor
``` go
func (*Message) Descriptor() ([]byte, []int)
```


### func (\*Message) IsCompressed
``` go
func (m *Message) IsCompressed() bool
```


### func (\*Message) ProtoMessage
``` go
func (*Message) ProtoMessage()
```


### func (\*Message) Reset
``` go
func (m *Message) Reset()
```


### func (\*Message) String
``` go
func (m *Message) String() string
```


## type MessageCompression
``` go
type MessageCompression int32
```


``` go
const (
    Message_NONE   MessageCompression = 0
    Message_ZLIB   MessageCompression = 1
    Message_SNAPPY MessageCompression = 2
)
```








### func (MessageCompression) EnumDescriptor
``` go
func (MessageCompression) EnumDescriptor() ([]byte, []int)
```


### func (MessageCompression) NewReader
``` go
func (c MessageCompression) NewReader(r io.Reader) io.ReadCloser
```


### func (MessageCompression) NewWriter
``` go
func (c MessageCompression) NewWriter(w io.Writer) io.WriteCloser
```


### func (MessageCompression) String
``` go
func (x MessageCompression) String() string
```


## type MessageSet
``` go
type MessageSet struct {
    // Offset is not populated by producers.
    Offset   int64      `protobuf:"varint,1,opt,name=Offset" json:"Offset,omitempty"`
    Messages []*Message `protobuf:"bytes,2,rep,name=Messages" json:"Messages,omitempty"`
}
```
MessageSet is the shared encoding that is used both on disk and over the wire.











### func (\*MessageSet) CheckIntegrity
``` go
func (m *MessageSet) CheckIntegrity() error
```


### func (\*MessageSet) Compress
``` go
func (m *MessageSet) Compress(codec MessageCompression) (*MessageSet, error)
```


### func (\*MessageSet) DecompressSet
``` go
func (m *MessageSet) DecompressSet() (*MessageSet, error)
```


### func (\*MessageSet) Descriptor
``` go
func (*MessageSet) Descriptor() ([]byte, []int)
```


### func (\*MessageSet) GetMessages
``` go
func (m *MessageSet) GetMessages() []*Message
```


### func (\*MessageSet) IsCompressedSet
``` go
func (m *MessageSet) IsCompressedSet() bool
```


### func (\*MessageSet) ProtoMessage
``` go
func (*MessageSet) ProtoMessage()
```


### func (\*MessageSet) Reset
``` go
func (m *MessageSet) Reset()
```


### func (\*MessageSet) String
``` go
func (m *MessageSet) String() string
```


## type ProducerClient
``` go
type ProducerClient interface {
    Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
    SendStream(ctx context.Context, opts ...grpc.CallOption) (Producer_SendStreamClient, error)
}
```








### func NewProducerClient
``` go
func NewProducerClient(cc *grpc.ClientConn) ProducerClient
```



## type ProducerServer
``` go
type ProducerServer interface {
    Send(context.Context, *SendRequest) (*SendResponse, error)
    SendStream(Producer_SendStreamServer) error
}
```










## type Producer_SendStreamClient
``` go
type Producer_SendStreamClient interface {
    Send(*SendRequest) error
    Recv() (*SendResponse, error)
    grpc.ClientStream
}
```










## type Producer_SendStreamServer
``` go
type Producer_SendStreamServer interface {
    Send(*SendResponse) error
    Recv() (*SendRequest, error)
    grpc.ServerStream
}
```










## type ReceiveRequest
``` go
type ReceiveRequest struct {
    Topic     string `protobuf:"bytes,1,opt,name=Topic" json:"Topic,omitempty"`
    Partition uint32 `protobuf:"varint,2,opt,name=Partition" json:"Partition,omitempty"`
    Offset    uint64 `protobuf:"varint,3,opt,name=Offset" json:"Offset,omitempty"`
}
```










### func (\*ReceiveRequest) Descriptor
``` go
func (*ReceiveRequest) Descriptor() ([]byte, []int)
```


### func (\*ReceiveRequest) ProtoMessage
``` go
func (*ReceiveRequest) ProtoMessage()
```


### func (\*ReceiveRequest) Reset
``` go
func (m *ReceiveRequest) Reset()
```


### func (\*ReceiveRequest) String
``` go
func (m *ReceiveRequest) String() string
```


## type ReceiveResponse
``` go
type ReceiveResponse struct {
    Topic      string      `protobuf:"bytes,1,opt,name=Topic" json:"Topic,omitempty"`
    Partition  uint32      `protobuf:"varint,2,opt,name=Partition" json:"Partition,omitempty"`
    MessageSet *MessageSet `protobuf:"bytes,3,opt,name=MessageSet" json:"MessageSet,omitempty"`
}
```










### func (\*ReceiveResponse) Descriptor
``` go
func (*ReceiveResponse) Descriptor() ([]byte, []int)
```


### func (\*ReceiveResponse) GetMessageSet
``` go
func (m *ReceiveResponse) GetMessageSet() *MessageSet
```


### func (\*ReceiveResponse) ProtoMessage
``` go
func (*ReceiveResponse) ProtoMessage()
```


### func (\*ReceiveResponse) Reset
``` go
func (m *ReceiveResponse) Reset()
```


### func (\*ReceiveResponse) String
``` go
func (m *ReceiveResponse) String() string
```


## type SendRequest
``` go
type SendRequest struct {
    Topic      string      `protobuf:"bytes,1,opt,name=Topic" json:"Topic,omitempty"`
    MessageSet *MessageSet `protobuf:"bytes,2,opt,name=MessageSet" json:"MessageSet,omitempty"`
}
```










### func (\*SendRequest) Descriptor
``` go
func (*SendRequest) Descriptor() ([]byte, []int)
```


### func (\*SendRequest) GetMessageSet
``` go
func (m *SendRequest) GetMessageSet() *MessageSet
```


### func (\*SendRequest) ProtoMessage
``` go
func (*SendRequest) ProtoMessage()
```


### func (\*SendRequest) Reset
``` go
func (m *SendRequest) Reset()
```


### func (\*SendRequest) String
``` go
func (m *SendRequest) String() string
```


## type SendResponse
``` go
type SendResponse struct {
}
```










### func (\*SendResponse) Descriptor
``` go
func (*SendResponse) Descriptor() ([]byte, []int)
```


### func (\*SendResponse) ProtoMessage
``` go
func (*SendResponse) ProtoMessage()
```


### func (\*SendResponse) Reset
``` go
func (m *SendResponse) Reset()
```


### func (\*SendResponse) String
``` go
func (m *SendResponse) String() string
```


## type Server
``` go
type Server interface {
    ProducerServer
    ConsumerServer
}
```













