package proto

// Needs protoc 3.0+
// on MacOS:
// $ brew install protobuf --devel
// $ go get github.com/golang/protobuf/...

//go:generate protoc -I/usr/local/include:.:$GOPATH/src --go_out=plugins=grpc:../gcanpb gcan.proto
