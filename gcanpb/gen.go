package gcanpb

// Needs protoc 3.0+
// on MacOS:
// $ brew install protobuf --devel
// $ go get github.com/gogo/protobuf/...

//go:generate protoc -I/usr/local/include:.:$GOPATH/src --go_out=plugins=grpc:. gcan.proto
//go:generate ./.genreadme.sh
