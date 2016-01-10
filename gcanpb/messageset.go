package gcanpb

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func (m *MessageSet) CheckIntegrity() error {
	for _, m := range m.Messages {
		if err := m.CheckIntegrity(); err != nil {
			return err
		}
	}
	return nil
}

func (m *MessageSet) IsCompressedSet() bool {
	if m.Messages == nil || len(m.Messages) == 0 {
		return false
	}
	return len(m.Messages) == 1 && m.Messages[0].IsCompressed()
}

func (m *MessageSet) DecompressSet() (*MessageSet, error) {
	if !m.IsCompressedSet() {
		return nil, fmt.Errorf("gcan: MessageSet does not appear to be compressed")
	}
	decompressor, err := m.Messages[0].Decompress()
	if err != nil {
		return nil, err
	}
	defer decompressor.Close()
	result := &MessageSet{}
	buf, err := ioutil.ReadAll(decompressor)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(buf, result)
	return result, err
}

func (m *MessageSet) Compress(codec MessageCompression) (*MessageSet, error) {
	compressed := new(bytes.Buffer) // sync.Pool candidate
	compressor := codec.NewWriter(compressed)

	data, err := proto.Marshal(m) // sync.Pool candidate
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(compressor, bytes.NewReader(data))
	compressor.Close()
	if err != nil {
		return nil, err
	}
	ms := &MessageSet{
		Messages: []*Message{
			{Value: compressed.Bytes()},
		},
	}
	// recompute CRC
	ms.Messages[0].Compression = codec
	ms.Messages[0].CRC = ms.Messages[0].ComputeCRC()
	return ms, nil
}
