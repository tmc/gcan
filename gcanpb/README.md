
# gcanpb
    import "github.com/tmc/gcan/gcanpb"

Package gcanpb is a generated protocol buffer package.

It is generated from these files:


	gcan.proto

It has these top-level messages:


	SendRequest
	SendResponse
	ReceiveRequest
	Response






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
    Recv() (*Response, error)
    grpc.ClientStream
}
```










## type Consumer_ReceiveServer
``` go
type Consumer_ReceiveServer interface {
    Send(*Response) error
    grpc.ServerStream
}
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
    Partition int32  `protobuf:"varint,2,opt,name=Partition" json:"Partition,omitempty"`
    Offset    int64  `protobuf:"varint,3,opt,name=Offset" json:"Offset,omitempty"`
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


## type Response
``` go
type Response struct {
    Key   string `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
    Value []byte `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
}
```










### func (\*Response) Descriptor
``` go
func (*Response) Descriptor() ([]byte, []int)
```


### func (\*Response) ProtoMessage
``` go
func (*Response) ProtoMessage()
```


### func (\*Response) Reset
``` go
func (m *Response) Reset()
```


### func (\*Response) String
``` go
func (m *Response) String() string
```


## type SendRequest
``` go
type SendRequest struct {
    Topic string `protobuf:"bytes,1,opt,name=Topic" json:"Topic,omitempty"`
    Key   string `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
    Value []byte `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
}
```










### func (\*SendRequest) Descriptor
``` go
func (*SendRequest) Descriptor() ([]byte, []int)
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





