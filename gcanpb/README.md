

# gcanpb
`import "github.com/tmc/gcan/gcanpb"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func NopWCloser(w io.Writer) io.WriteCloser](#NopWCloser)
* [func RegisterConsumerServer(s *grpc.Server, srv ConsumerServer)](#RegisterConsumerServer)
* [func RegisterProducerServer(s *grpc.Server, srv ProducerServer)](#RegisterProducerServer)
* [type ConsumerClient](#ConsumerClient)
  * [func NewConsumerClient(cc *grpc.ClientConn) ConsumerClient](#NewConsumerClient)
* [type ConsumerServer](#ConsumerServer)
* [type Consumer_ReceiveClient](#Consumer_ReceiveClient)
* [type Consumer_ReceiveServer](#Consumer_ReceiveServer)
* [type Err](#Err)
  * [func (Err) EnumDescriptor() ([]byte, []int)](#Err.EnumDescriptor)
  * [func (x Err) Error() string](#Err.Error)
  * [func (x Err) String() string](#Err.String)
* [type Message](#Message)
  * [func (m *Message) CheckIntegrity() error](#Message.CheckIntegrity)
  * [func (m *Message) ComputeCRC() uint32](#Message.ComputeCRC)
  * [func (m *Message) Decompress() (io.ReadCloser, error)](#Message.Decompress)
  * [func (*Message) Descriptor() ([]byte, []int)](#Message.Descriptor)
  * [func (m *Message) GetCRC() uint32](#Message.GetCRC)
  * [func (m *Message) GetCompression() MessageCompression](#Message.GetCompression)
  * [func (m *Message) GetKey() []byte](#Message.GetKey)
  * [func (m *Message) GetValue() []byte](#Message.GetValue)
  * [func (m *Message) IsCompressed() bool](#Message.IsCompressed)
  * [func (*Message) ProtoMessage()](#Message.ProtoMessage)
  * [func (m *Message) Reset()](#Message.Reset)
  * [func (m *Message) String() string](#Message.String)
  * [func (m *Message) XXX_DiscardUnknown()](#Message.XXX_DiscardUnknown)
  * [func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)](#Message.XXX_Marshal)
  * [func (dst *Message) XXX_Merge(src proto.Message)](#Message.XXX_Merge)
  * [func (m *Message) XXX_Size() int](#Message.XXX_Size)
  * [func (m *Message) XXX_Unmarshal(b []byte) error](#Message.XXX_Unmarshal)
* [type MessageCompression](#MessageCompression)
  * [func (MessageCompression) EnumDescriptor() ([]byte, []int)](#MessageCompression.EnumDescriptor)
  * [func (c MessageCompression) NewReader(r io.Reader) io.ReadCloser](#MessageCompression.NewReader)
  * [func (c MessageCompression) NewWriter(w io.Writer) io.WriteCloser](#MessageCompression.NewWriter)
  * [func (x MessageCompression) String() string](#MessageCompression.String)
* [type MessageSet](#MessageSet)
  * [func (m *MessageSet) CheckIntegrity() error](#MessageSet.CheckIntegrity)
  * [func (m *MessageSet) Compress(codec MessageCompression) (*MessageSet, error)](#MessageSet.Compress)
  * [func (m *MessageSet) DecompressSet() (*MessageSet, error)](#MessageSet.DecompressSet)
  * [func (*MessageSet) Descriptor() ([]byte, []int)](#MessageSet.Descriptor)
  * [func (m *MessageSet) GetMessages() []*Message](#MessageSet.GetMessages)
  * [func (m *MessageSet) GetOffset() int64](#MessageSet.GetOffset)
  * [func (m *MessageSet) IsCompressedSet() bool](#MessageSet.IsCompressedSet)
  * [func (*MessageSet) ProtoMessage()](#MessageSet.ProtoMessage)
  * [func (m *MessageSet) Reset()](#MessageSet.Reset)
  * [func (m *MessageSet) String() string](#MessageSet.String)
  * [func (m *MessageSet) XXX_DiscardUnknown()](#MessageSet.XXX_DiscardUnknown)
  * [func (m *MessageSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)](#MessageSet.XXX_Marshal)
  * [func (dst *MessageSet) XXX_Merge(src proto.Message)](#MessageSet.XXX_Merge)
  * [func (m *MessageSet) XXX_Size() int](#MessageSet.XXX_Size)
  * [func (m *MessageSet) XXX_Unmarshal(b []byte) error](#MessageSet.XXX_Unmarshal)
* [type ProducerClient](#ProducerClient)
  * [func NewProducerClient(cc *grpc.ClientConn) ProducerClient](#NewProducerClient)
* [type ProducerServer](#ProducerServer)
* [type Producer_SendStreamClient](#Producer_SendStreamClient)
* [type Producer_SendStreamServer](#Producer_SendStreamServer)
* [type ReceiveRequest](#ReceiveRequest)
  * [func (*ReceiveRequest) Descriptor() ([]byte, []int)](#ReceiveRequest.Descriptor)
  * [func (m *ReceiveRequest) GetOffset() uint64](#ReceiveRequest.GetOffset)
  * [func (m *ReceiveRequest) GetPartition() uint32](#ReceiveRequest.GetPartition)
  * [func (m *ReceiveRequest) GetTopic() string](#ReceiveRequest.GetTopic)
  * [func (*ReceiveRequest) ProtoMessage()](#ReceiveRequest.ProtoMessage)
  * [func (m *ReceiveRequest) Reset()](#ReceiveRequest.Reset)
  * [func (m *ReceiveRequest) String() string](#ReceiveRequest.String)
  * [func (m *ReceiveRequest) XXX_DiscardUnknown()](#ReceiveRequest.XXX_DiscardUnknown)
  * [func (m *ReceiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)](#ReceiveRequest.XXX_Marshal)
  * [func (dst *ReceiveRequest) XXX_Merge(src proto.Message)](#ReceiveRequest.XXX_Merge)
  * [func (m *ReceiveRequest) XXX_Size() int](#ReceiveRequest.XXX_Size)
  * [func (m *ReceiveRequest) XXX_Unmarshal(b []byte) error](#ReceiveRequest.XXX_Unmarshal)
* [type ReceiveResponse](#ReceiveResponse)
  * [func (*ReceiveResponse) Descriptor() ([]byte, []int)](#ReceiveResponse.Descriptor)
  * [func (m *ReceiveResponse) GetMessageSet() *MessageSet](#ReceiveResponse.GetMessageSet)
  * [func (m *ReceiveResponse) GetPartition() uint32](#ReceiveResponse.GetPartition)
  * [func (m *ReceiveResponse) GetTopic() string](#ReceiveResponse.GetTopic)
  * [func (*ReceiveResponse) ProtoMessage()](#ReceiveResponse.ProtoMessage)
  * [func (m *ReceiveResponse) Reset()](#ReceiveResponse.Reset)
  * [func (m *ReceiveResponse) String() string](#ReceiveResponse.String)
  * [func (m *ReceiveResponse) XXX_DiscardUnknown()](#ReceiveResponse.XXX_DiscardUnknown)
  * [func (m *ReceiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)](#ReceiveResponse.XXX_Marshal)
  * [func (dst *ReceiveResponse) XXX_Merge(src proto.Message)](#ReceiveResponse.XXX_Merge)
  * [func (m *ReceiveResponse) XXX_Size() int](#ReceiveResponse.XXX_Size)
  * [func (m *ReceiveResponse) XXX_Unmarshal(b []byte) error](#ReceiveResponse.XXX_Unmarshal)
* [type SendRequest](#SendRequest)
  * [func (*SendRequest) Descriptor() ([]byte, []int)](#SendRequest.Descriptor)
  * [func (m *SendRequest) GetMessageSet() *MessageSet](#SendRequest.GetMessageSet)
  * [func (m *SendRequest) GetTopic() string](#SendRequest.GetTopic)
  * [func (*SendRequest) ProtoMessage()](#SendRequest.ProtoMessage)
  * [func (m *SendRequest) Reset()](#SendRequest.Reset)
  * [func (m *SendRequest) String() string](#SendRequest.String)
  * [func (m *SendRequest) XXX_DiscardUnknown()](#SendRequest.XXX_DiscardUnknown)
  * [func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)](#SendRequest.XXX_Marshal)
  * [func (dst *SendRequest) XXX_Merge(src proto.Message)](#SendRequest.XXX_Merge)
  * [func (m *SendRequest) XXX_Size() int](#SendRequest.XXX_Size)
  * [func (m *SendRequest) XXX_Unmarshal(b []byte) error](#SendRequest.XXX_Unmarshal)
* [type SendResponse](#SendResponse)
  * [func (*SendResponse) Descriptor() ([]byte, []int)](#SendResponse.Descriptor)
  * [func (*SendResponse) ProtoMessage()](#SendResponse.ProtoMessage)
  * [func (m *SendResponse) Reset()](#SendResponse.Reset)
  * [func (m *SendResponse) String() string](#SendResponse.String)
  * [func (m *SendResponse) XXX_DiscardUnknown()](#SendResponse.XXX_DiscardUnknown)
  * [func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)](#SendResponse.XXX_Marshal)
  * [func (dst *SendResponse) XXX_Merge(src proto.Message)](#SendResponse.XXX_Merge)
  * [func (m *SendResponse) XXX_Size() int](#SendResponse.XXX_Size)
  * [func (m *SendResponse) XXX_Unmarshal(b []byte) error](#SendResponse.XXX_Unmarshal)
* [type Server](#Server)


#### <a name="pkg-files">Package files</a>
[compression.go](/src/github.com/tmc/gcan/gcanpb/compression.go) [compression_snappy.go](/src/github.com/tmc/gcan/gcanpb/compression_snappy.go) [compression_zlib.go](/src/github.com/tmc/gcan/gcanpb/compression_zlib.go) [err.go](/src/github.com/tmc/gcan/gcanpb/err.go) [gcan.pb.go](/src/github.com/tmc/gcan/gcanpb/gcan.pb.go) [gen.go](/src/github.com/tmc/gcan/gcanpb/gen.go) [ioutil.go](/src/github.com/tmc/gcan/gcanpb/ioutil.go) [message.go](/src/github.com/tmc/gcan/gcanpb/message.go) [messageset.go](/src/github.com/tmc/gcan/gcanpb/messageset.go) [server.go](/src/github.com/tmc/gcan/gcanpb/server.go) 



## <a name="pkg-variables">Variables</a>
``` go
var Err_name = map[int32]string{
    0: "OK",
    1: "Unknown",
    2: "CRCChecksumFailed",
}
```
``` go
var Err_value = map[string]int32{
    "OK":                0,
    "Unknown":           1,
    "CRCChecksumFailed": 2,
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


## <a name="NopWCloser">func</a> [NopWCloser](/src/target/ioutil.go?s=213:256#L13)
``` go
func NopWCloser(w io.Writer) io.WriteCloser
```
NopWCloser returns a WriteCloser with a no-op Close method wrapping
the provided Writer w.



## <a name="RegisterConsumerServer">func</a> [RegisterConsumerServer](/src/target/gcan.pb.go?s=17076:17139#L575)
``` go
func RegisterConsumerServer(s *grpc.Server, srv ConsumerServer)
```


## <a name="RegisterProducerServer">func</a> [RegisterProducerServer](/src/target/gcan.pb.go?s=13862:13925#L455)
``` go
func RegisterProducerServer(s *grpc.Server, srv ProducerServer)
```



## <a name="ConsumerClient">type</a> [ConsumerClient](/src/target/gcan.pb.go?s=15822:15962#L526)
``` go
type ConsumerClient interface {
    Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (Consumer_ReceiveClient, error)
}
```
ConsumerClient is the client API for Consumer service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to <a href="https://godoc.org/google.golang.org/grpc#ClientConn.NewStream">https://godoc.org/google.golang.org/grpc#ClientConn.NewStream</a>.







### <a name="NewConsumerClient">func</a> [NewConsumerClient](/src/target/gcan.pb.go?s=16017:16075#L534)
``` go
func NewConsumerClient(cc *grpc.ClientConn) ConsumerClient
```




## <a name="ConsumerServer">type</a> [ConsumerServer](/src/target/gcan.pb.go?s=16985:17074#L571)
``` go
type ConsumerServer interface {
    Receive(*ReceiveRequest, Consumer_ReceiveServer) error
}
```
ConsumerServer is the server API for Consumer service.










## <a name="Consumer_ReceiveClient">type</a> [Consumer_ReceiveClient](/src/target/gcan.pb.go?s=16589:16683#L553)
``` go
type Consumer_ReceiveClient interface {
    Recv() (*ReceiveResponse, error)
    grpc.ClientStream
}
```









## <a name="Consumer_ReceiveServer">type</a> [Consumer_ReceiveServer](/src/target/gcan.pb.go?s=17435:17525#L587)
``` go
type Consumer_ReceiveServer interface {
    Send(*ReceiveResponse) error
    grpc.ServerStream
}
```









## <a name="Err">type</a> [Err](/src/target/gcan.pb.go?s=704:718#L26)
``` go
type Err int32
```

``` go
const (
    Err_OK                Err = 0
    Err_Unknown           Err = 1
    Err_CRCChecksumFailed Err = 2
)
```









### <a name="Err.EnumDescriptor">func</a> (Err) [EnumDescriptor](/src/target/gcan.pb.go?s=1097:1140#L48)
``` go
func (Err) EnumDescriptor() ([]byte, []int)
```



### <a name="Err.Error">func</a> (Err) [Error](/src/target/err.go?s=16:43#L3)
``` go
func (x Err) Error() string
```



### <a name="Err.String">func</a> (Err) [String](/src/target/gcan.pb.go?s=1021:1049#L45)
``` go
func (x Err) String() string
```



## <a name="Message">type</a> [Message](/src/target/gcan.pb.go?s=3364:4042#L127)
``` go
type Message struct {
    // CRC32 of all following fields
    CRC                  uint32             `protobuf:"varint,1,opt,name=CRC,proto3" json:"CRC,omitempty"`
    Compression          MessageCompression `protobuf:"varint,2,opt,name=Compression,proto3,enum=gcanpb.MessageCompression" json:"Compression,omitempty"`
    Key                  []byte             `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
    Value                []byte             `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
    XXX_NoUnkeyedLiteral struct{}           `json:"-"`
    XXX_unrecognized     []byte             `json:"-"`
    XXX_sizecache        int32              `json:"-"`
}
```
Message is an individual key-value pair.










### <a name="Message.CheckIntegrity">func</a> (\*Message) [CheckIntegrity](/src/target/message.go?s=164:204#L15)
``` go
func (m *Message) CheckIntegrity() error
```



### <a name="Message.ComputeCRC">func</a> (\*Message) [ComputeCRC](/src/target/message.go?s=286:323#L22)
``` go
func (m *Message) ComputeCRC() uint32
```



### <a name="Message.Decompress">func</a> (\*Message) [Decompress](/src/target/message.go?s=608:661#L33)
``` go
func (m *Message) Decompress() (io.ReadCloser, error)
```



### <a name="Message.Descriptor">func</a> (\*Message) [Descriptor](/src/target/gcan.pb.go?s=4206:4250#L141)
``` go
func (*Message) Descriptor() ([]byte, []int)
```



### <a name="Message.GetCRC">func</a> (\*Message) [GetCRC](/src/target/gcan.pb.go?s=4864:4897#L162)
``` go
func (m *Message) GetCRC() uint32
```



### <a name="Message.GetCompression">func</a> (\*Message) [GetCompression](/src/target/gcan.pb.go?s=4946:4999#L169)
``` go
func (m *Message) GetCompression() MessageCompression
```



### <a name="Message.GetKey">func</a> (\*Message) [GetKey](/src/target/gcan.pb.go?s=5067:5100#L176)
``` go
func (m *Message) GetKey() []byte
```



### <a name="Message.GetValue">func</a> (\*Message) [GetValue](/src/target/gcan.pb.go?s=5151:5186#L183)
``` go
func (m *Message) GetValue() []byte
```



### <a name="Message.IsCompressed">func</a> (\*Message) [IsCompressed](/src/target/message.go?s=83:120#L11)
``` go
func (m *Message) IsCompressed() bool
```



### <a name="Message.ProtoMessage">func</a> (\*Message) [ProtoMessage](/src/target/gcan.pb.go?s=4169:4199#L140)
``` go
func (*Message) ProtoMessage()
```



### <a name="Message.Reset">func</a> (\*Message) [Reset](/src/target/gcan.pb.go?s=4044:4069#L138)
``` go
func (m *Message) Reset()
```



### <a name="Message.String">func</a> (\*Message) [String](/src/target/gcan.pb.go?s=4097:4130#L139)
``` go
func (m *Message) String() string
```



### <a name="Message.XXX_DiscardUnknown">func</a> (\*Message) [XXX_DiscardUnknown](/src/target/gcan.pb.go?s=4722:4760#L156)
``` go
func (m *Message) XXX_DiscardUnknown()
```



### <a name="Message.XXX_Marshal">func</a> (\*Message) [XXX_Marshal](/src/target/gcan.pb.go?s=4410:4485#L147)
``` go
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```



### <a name="Message.XXX_Merge">func</a> (\*Message) [XXX_Merge](/src/target/gcan.pb.go?s=4551:4599#L150)
``` go
func (dst *Message) XXX_Merge(src proto.Message)
```



### <a name="Message.XXX_Size">func</a> (\*Message) [XXX_Size](/src/target/gcan.pb.go?s=4645:4677#L153)
``` go
func (m *Message) XXX_Size() int
```



### <a name="Message.XXX_Unmarshal">func</a> (\*Message) [XXX_Unmarshal](/src/target/gcan.pb.go?s=4310:4357#L144)
``` go
func (m *Message) XXX_Unmarshal(b []byte) error
```



## <a name="MessageCompression">type</a> [MessageCompression](/src/target/gcan.pb.go?s=1201:1230#L52)
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









### <a name="MessageCompression.EnumDescriptor">func</a> (MessageCompression) [EnumDescriptor](/src/target/gcan.pb.go?s=1648:1706#L74)
``` go
func (MessageCompression) EnumDescriptor() ([]byte, []int)
```



### <a name="MessageCompression.NewReader">func</a> (MessageCompression) [NewReader](/src/target/compression.go?s=47:111#L8)
``` go
func (c MessageCompression) NewReader(r io.Reader) io.ReadCloser
```



### <a name="MessageCompression.NewWriter">func</a> (MessageCompression) [NewWriter](/src/target/compression.go?s=261:326#L19)
``` go
func (c MessageCompression) NewWriter(w io.Writer) io.WriteCloser
```



### <a name="MessageCompression.String">func</a> (MessageCompression) [String](/src/target/gcan.pb.go?s=1542:1585#L71)
``` go
func (x MessageCompression) String() string
```



## <a name="MessageSet">type</a> [MessageSet](/src/target/gcan.pb.go?s=1852:2259#L79)
``` go
type MessageSet struct {
    // Offset is not populated by producers.
    Offset               int64      `protobuf:"varint,1,opt,name=Offset,proto3" json:"Offset,omitempty"`
    Messages             []*Message `protobuf:"bytes,2,rep,name=Messages,proto3" json:"Messages,omitempty"`
    XXX_NoUnkeyedLiteral struct{}   `json:"-"`
    XXX_unrecognized     []byte     `json:"-"`
    XXX_sizecache        int32      `json:"-"`
}
```
MessageSet is the shared encoding that is used both on disk and over the wire.










### <a name="MessageSet.CheckIntegrity">func</a> (\*MessageSet) [CheckIntegrity](/src/target/messageset.go?s=100:143#L12)
``` go
func (m *MessageSet) CheckIntegrity() error
```



### <a name="MessageSet.Compress">func</a> (\*MessageSet) [Compress](/src/target/messageset.go?s=882:958#L46)
``` go
func (m *MessageSet) Compress(codec MessageCompression) (*MessageSet, error)
```



### <a name="MessageSet.DecompressSet">func</a> (\*MessageSet) [DecompressSet](/src/target/messageset.go?s=435:492#L28)
``` go
func (m *MessageSet) DecompressSet() (*MessageSet, error)
```



### <a name="MessageSet.Descriptor">func</a> (\*MessageSet) [Descriptor](/src/target/gcan.pb.go?s=2435:2482#L91)
``` go
func (*MessageSet) Descriptor() ([]byte, []int)
```



### <a name="MessageSet.GetMessages">func</a> (\*MessageSet) [GetMessages](/src/target/gcan.pb.go?s=3219:3264#L119)
``` go
func (m *MessageSet) GetMessages() []*Message
```



### <a name="MessageSet.GetOffset">func</a> (\*MessageSet) [GetOffset](/src/target/gcan.pb.go?s=3129:3167#L112)
``` go
func (m *MessageSet) GetOffset() int64
```



### <a name="MessageSet.IsCompressedSet">func</a> (\*MessageSet) [IsCompressedSet](/src/target/messageset.go?s=259:302#L21)
``` go
func (m *MessageSet) IsCompressedSet() bool
```



### <a name="MessageSet.ProtoMessage">func</a> (\*MessageSet) [ProtoMessage](/src/target/gcan.pb.go?s=2395:2428#L90)
``` go
func (*MessageSet) ProtoMessage()
```



### <a name="MessageSet.Reset">func</a> (\*MessageSet) [Reset](/src/target/gcan.pb.go?s=2261:2289#L88)
``` go
func (m *MessageSet) Reset()
```



### <a name="MessageSet.String">func</a> (\*MessageSet) [String](/src/target/gcan.pb.go?s=2320:2356#L89)
``` go
func (m *MessageSet) String() string
```



### <a name="MessageSet.XXX_DiscardUnknown">func</a> (\*MessageSet) [XXX_DiscardUnknown](/src/target/gcan.pb.go?s=2978:3019#L106)
``` go
func (m *MessageSet) XXX_DiscardUnknown()
```



### <a name="MessageSet.XXX_Marshal">func</a> (\*MessageSet) [XXX_Marshal](/src/target/gcan.pb.go?s=2648:2726#L97)
``` go
func (m *MessageSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```



### <a name="MessageSet.XXX_Merge">func</a> (\*MessageSet) [XXX_Merge](/src/target/gcan.pb.go?s=2795:2846#L100)
``` go
func (dst *MessageSet) XXX_Merge(src proto.Message)
```



### <a name="MessageSet.XXX_Size">func</a> (\*MessageSet) [XXX_Size](/src/target/gcan.pb.go?s=2895:2930#L103)
``` go
func (m *MessageSet) XXX_Size() int
```



### <a name="MessageSet.XXX_Unmarshal">func</a> (\*MessageSet) [XXX_Unmarshal](/src/target/gcan.pb.go?s=2542:2592#L94)
``` go
func (m *MessageSet) XXX_Unmarshal(b []byte) error
```



## <a name="ProducerClient">type</a> [ProducerClient](/src/target/gcan.pb.go?s=12238:12456#L396)
``` go
type ProducerClient interface {
    Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
    SendStream(ctx context.Context, opts ...grpc.CallOption) (Producer_SendStreamClient, error)
}
```
ProducerClient is the client API for Producer service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to <a href="https://godoc.org/google.golang.org/grpc#ClientConn.NewStream">https://godoc.org/google.golang.org/grpc#ClientConn.NewStream</a>.







### <a name="NewProducerClient">func</a> [NewProducerClient](/src/target/gcan.pb.go?s=12511:12569#L405)
``` go
func NewProducerClient(cc *grpc.ClientConn) ProducerClient
```




## <a name="ProducerServer">type</a> [ProducerServer](/src/target/gcan.pb.go?s=13722:13860#L450)
``` go
type ProducerServer interface {
    Send(context.Context, *SendRequest) (*SendResponse, error)
    SendStream(Producer_SendStreamServer) error
}
```
ProducerServer is the server API for Producer service.










## <a name="Producer_SendStreamClient">type</a> [Producer_SendStreamClient](/src/target/gcan.pb.go?s=13199:13319#L427)
``` go
type Producer_SendStreamClient interface {
    Send(*SendRequest) error
    Recv() (*SendResponse, error)
    grpc.ClientStream
}
```









## <a name="Producer_SendStreamServer">type</a> [Producer_SendStreamServer](/src/target/gcan.pb.go?s=14732:14852#L481)
``` go
type Producer_SendStreamServer interface {
    Send(*SendResponse) error
    Recv() (*SendRequest, error)
    grpc.ServerStream
}
```









## <a name="ReceiveRequest">type</a> [ReceiveRequest](/src/target/gcan.pb.go?s=7751:8210#L266)
``` go
type ReceiveRequest struct {
    Topic                string   `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
    Partition            uint32   `protobuf:"varint,2,opt,name=Partition,proto3" json:"Partition,omitempty"`
    Offset               uint64   `protobuf:"varint,3,opt,name=Offset,proto3" json:"Offset,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}
```









### <a name="ReceiveRequest.Descriptor">func</a> (\*ReceiveRequest) [Descriptor](/src/target/gcan.pb.go?s=8402:8453#L278)
``` go
func (*ReceiveRequest) Descriptor() ([]byte, []int)
```



### <a name="ReceiveRequest.GetOffset">func</a> (\*ReceiveRequest) [GetOffset](/src/target/gcan.pb.go?s=9339:9382#L313)
``` go
func (m *ReceiveRequest) GetOffset() uint64
```



### <a name="ReceiveRequest.GetPartition">func</a> (\*ReceiveRequest) [GetPartition](/src/target/gcan.pb.go?s=9238:9284#L306)
``` go
func (m *ReceiveRequest) GetPartition() uint32
```



### <a name="ReceiveRequest.GetTopic">func</a> (\*ReceiveRequest) [GetTopic](/src/target/gcan.pb.go?s=9144:9186#L299)
``` go
func (m *ReceiveRequest) GetTopic() string
```



### <a name="ReceiveRequest.ProtoMessage">func</a> (\*ReceiveRequest) [ProtoMessage](/src/target/gcan.pb.go?s=8358:8395#L277)
``` go
func (*ReceiveRequest) ProtoMessage()
```



### <a name="ReceiveRequest.Reset">func</a> (\*ReceiveRequest) [Reset](/src/target/gcan.pb.go?s=8212:8244#L275)
``` go
func (m *ReceiveRequest) Reset()
```



### <a name="ReceiveRequest.String">func</a> (\*ReceiveRequest) [String](/src/target/gcan.pb.go?s=8279:8319#L276)
``` go
func (m *ReceiveRequest) String() string
```



### <a name="ReceiveRequest.XXX_DiscardUnknown">func</a> (\*ReceiveRequest) [XXX_DiscardUnknown](/src/target/gcan.pb.go?s=8981:9026#L293)
``` go
func (m *ReceiveRequest) XXX_DiscardUnknown()
```



### <a name="ReceiveRequest.XXX_Marshal">func</a> (\*ReceiveRequest) [XXX_Marshal](/src/target/gcan.pb.go?s=8627:8709#L284)
``` go
func (m *ReceiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```



### <a name="ReceiveRequest.XXX_Merge">func</a> (\*ReceiveRequest) [XXX_Merge](/src/target/gcan.pb.go?s=8782:8837#L287)
``` go
func (dst *ReceiveRequest) XXX_Merge(src proto.Message)
```



### <a name="ReceiveRequest.XXX_Size">func</a> (\*ReceiveRequest) [XXX_Size](/src/target/gcan.pb.go?s=8890:8929#L290)
``` go
func (m *ReceiveRequest) XXX_Size() int
```



### <a name="ReceiveRequest.XXX_Unmarshal">func</a> (\*ReceiveRequest) [XXX_Unmarshal](/src/target/gcan.pb.go?s=8513:8567#L281)
``` go
func (m *ReceiveRequest) XXX_Unmarshal(b []byte) error
```



## <a name="ReceiveResponse">type</a> [ReceiveResponse](/src/target/gcan.pb.go?s=9434:9919#L320)
``` go
type ReceiveResponse struct {
    Topic                string      `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
    Partition            uint32      `protobuf:"varint,2,opt,name=Partition,proto3" json:"Partition,omitempty"`
    MessageSet           *MessageSet `protobuf:"bytes,3,opt,name=MessageSet,proto3" json:"MessageSet,omitempty"`
    XXX_NoUnkeyedLiteral struct{}    `json:"-"`
    XXX_unrecognized     []byte      `json:"-"`
    XXX_sizecache        int32       `json:"-"`
}
```









### <a name="ReceiveResponse.Descriptor">func</a> (\*ReceiveResponse) [Descriptor](/src/target/gcan.pb.go?s=10115:10167#L332)
``` go
func (*ReceiveResponse) Descriptor() ([]byte, []int)
```



### <a name="ReceiveResponse.GetMessageSet">func</a> (\*ReceiveResponse) [GetMessageSet](/src/target/gcan.pb.go?s=11066:11119#L367)
``` go
func (m *ReceiveResponse) GetMessageSet() *MessageSet
```



### <a name="ReceiveResponse.GetPartition">func</a> (\*ReceiveResponse) [GetPartition](/src/target/gcan.pb.go?s=10964:11011#L360)
``` go
func (m *ReceiveResponse) GetPartition() uint32
```



### <a name="ReceiveResponse.GetTopic">func</a> (\*ReceiveResponse) [GetTopic](/src/target/gcan.pb.go?s=10869:10912#L353)
``` go
func (m *ReceiveResponse) GetTopic() string
```



### <a name="ReceiveResponse.ProtoMessage">func</a> (\*ReceiveResponse) [ProtoMessage](/src/target/gcan.pb.go?s=10070:10108#L331)
``` go
func (*ReceiveResponse) ProtoMessage()
```



### <a name="ReceiveResponse.Reset">func</a> (\*ReceiveResponse) [Reset](/src/target/gcan.pb.go?s=9921:9954#L329)
``` go
func (m *ReceiveResponse) Reset()
```



### <a name="ReceiveResponse.String">func</a> (\*ReceiveResponse) [String](/src/target/gcan.pb.go?s=9990:10031#L330)
``` go
func (m *ReceiveResponse) String() string
```



### <a name="ReceiveResponse.XXX_DiscardUnknown">func</a> (\*ReceiveResponse) [XXX_DiscardUnknown](/src/target/gcan.pb.go?s=10703:10749#L347)
``` go
func (m *ReceiveResponse) XXX_DiscardUnknown()
```



### <a name="ReceiveResponse.XXX_Marshal">func</a> (\*ReceiveResponse) [XXX_Marshal](/src/target/gcan.pb.go?s=10343:10426#L338)
``` go
func (m *ReceiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```



### <a name="ReceiveResponse.XXX_Merge">func</a> (\*ReceiveResponse) [XXX_Merge](/src/target/gcan.pb.go?s=10500:10556#L341)
``` go
func (dst *ReceiveResponse) XXX_Merge(src proto.Message)
```



### <a name="ReceiveResponse.XXX_Size">func</a> (\*ReceiveResponse) [XXX_Size](/src/target/gcan.pb.go?s=10610:10650#L344)
``` go
func (m *ReceiveResponse) XXX_Size() int
```



### <a name="ReceiveResponse.XXX_Unmarshal">func</a> (\*ReceiveResponse) [XXX_Unmarshal](/src/target/gcan.pb.go?s=10227:10282#L335)
``` go
func (m *ReceiveResponse) XXX_Unmarshal(b []byte) error
```



## <a name="SendRequest">type</a> [SendRequest](/src/target/gcan.pb.go?s=5239:5611#L190)
``` go
type SendRequest struct {
    Topic                string      `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
    MessageSet           *MessageSet `protobuf:"bytes,2,opt,name=MessageSet,proto3" json:"MessageSet,omitempty"`
    XXX_NoUnkeyedLiteral struct{}    `json:"-"`
    XXX_unrecognized     []byte      `json:"-"`
    XXX_sizecache        int32       `json:"-"`
}
```









### <a name="SendRequest.Descriptor">func</a> (\*SendRequest) [Descriptor](/src/target/gcan.pb.go?s=5791:5839#L201)
``` go
func (*SendRequest) Descriptor() ([]byte, []int)
```



### <a name="SendRequest.GetMessageSet">func</a> (\*SendRequest) [GetMessageSet](/src/target/gcan.pb.go?s=6588:6637#L229)
``` go
func (m *SendRequest) GetMessageSet() *MessageSet
```



### <a name="SendRequest.GetTopic">func</a> (\*SendRequest) [GetTopic](/src/target/gcan.pb.go?s=6497:6536#L222)
``` go
func (m *SendRequest) GetTopic() string
```



### <a name="SendRequest.ProtoMessage">func</a> (\*SendRequest) [ProtoMessage](/src/target/gcan.pb.go?s=5750:5784#L200)
``` go
func (*SendRequest) ProtoMessage()
```



### <a name="SendRequest.Reset">func</a> (\*SendRequest) [Reset](/src/target/gcan.pb.go?s=5613:5642#L198)
``` go
func (m *SendRequest) Reset()
```



### <a name="SendRequest.String">func</a> (\*SendRequest) [String](/src/target/gcan.pb.go?s=5674:5711#L199)
``` go
func (m *SendRequest) String() string
```



### <a name="SendRequest.XXX_DiscardUnknown">func</a> (\*SendRequest) [XXX_DiscardUnknown](/src/target/gcan.pb.go?s=6343:6385#L216)
``` go
func (m *SendRequest) XXX_DiscardUnknown()
```



### <a name="SendRequest.XXX_Marshal">func</a> (\*SendRequest) [XXX_Marshal](/src/target/gcan.pb.go?s=6007:6086#L207)
``` go
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```



### <a name="SendRequest.XXX_Merge">func</a> (\*SendRequest) [XXX_Merge](/src/target/gcan.pb.go?s=6156:6208#L210)
``` go
func (dst *SendRequest) XXX_Merge(src proto.Message)
```



### <a name="SendRequest.XXX_Size">func</a> (\*SendRequest) [XXX_Size](/src/target/gcan.pb.go?s=6258:6294#L213)
``` go
func (m *SendRequest) XXX_Size() int
```



### <a name="SendRequest.XXX_Unmarshal">func</a> (\*SendRequest) [XXX_Unmarshal](/src/target/gcan.pb.go?s=5899:5950#L204)
``` go
func (m *SendRequest) XXX_Unmarshal(b []byte) error
```



## <a name="SendResponse">type</a> [SendResponse](/src/target/gcan.pb.go?s=6695:6849#L236)
``` go
type SendResponse struct {
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}
```









### <a name="SendResponse.Descriptor">func</a> (\*SendResponse) [Descriptor](/src/target/gcan.pb.go?s=7033:7082#L245)
``` go
func (*SendResponse) Descriptor() ([]byte, []int)
```



### <a name="SendResponse.ProtoMessage">func</a> (\*SendResponse) [ProtoMessage](/src/target/gcan.pb.go?s=6991:7026#L244)
``` go
func (*SendResponse) ProtoMessage()
```



### <a name="SendResponse.Reset">func</a> (\*SendResponse) [Reset](/src/target/gcan.pb.go?s=6851:6881#L242)
``` go
func (m *SendResponse) Reset()
```



### <a name="SendResponse.String">func</a> (\*SendResponse) [String](/src/target/gcan.pb.go?s=6914:6952#L243)
``` go
func (m *SendResponse) String() string
```



### <a name="SendResponse.XXX_DiscardUnknown">func</a> (\*SendResponse) [XXX_DiscardUnknown](/src/target/gcan.pb.go?s=7594:7637#L260)
``` go
func (m *SendResponse) XXX_DiscardUnknown()
```



### <a name="SendResponse.XXX_Marshal">func</a> (\*SendResponse) [XXX_Marshal](/src/target/gcan.pb.go?s=7252:7332#L251)
``` go
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```



### <a name="SendResponse.XXX_Merge">func</a> (\*SendResponse) [XXX_Merge](/src/target/gcan.pb.go?s=7403:7456#L254)
``` go
func (dst *SendResponse) XXX_Merge(src proto.Message)
```



### <a name="SendResponse.XXX_Size">func</a> (\*SendResponse) [XXX_Size](/src/target/gcan.pb.go?s=7507:7544#L257)
``` go
func (m *SendResponse) XXX_Size() int
```



### <a name="SendResponse.XXX_Unmarshal">func</a> (\*SendResponse) [XXX_Unmarshal](/src/target/gcan.pb.go?s=7142:7194#L248)
``` go
func (m *SendResponse) XXX_Unmarshal(b []byte) error
```



## <a name="Server">type</a> [Server](/src/target/server.go?s=16:73#L3)
``` go
type Server interface {
    ProducerServer
    ConsumerServer
}
```










