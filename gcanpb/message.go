package gcanpb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
)

func (m *Message) IsCompressed() bool {
	return m.Compression != Message_NONE
}

func (m *Message) CheckIntegrity() error {
	if m.ComputeCRC() != m.CRC {
		return Err_CRCChecksumFailed
	}
	return nil
}

func (m *Message) ComputeCRC() uint32 {
	var crc uint32
	buf := make([]byte, 4) // sync.Pool candidate

	binary.BigEndian.PutUint32(buf, uint32(m.Compression))
	crc = crc32.Update(crc, crc32.IEEETable, buf)
	crc = crc32.Update(crc, crc32.IEEETable, m.Key)
	crc = crc32.Update(crc, crc32.IEEETable, m.Value)
	return crc
}

func (m *Message) Decompress() (io.ReadCloser, error) {
	if !m.IsCompressed() {
		return nil, fmt.Errorf("gcan: Message is not compressed")
	}
	return m.Compression.NewReader(bytes.NewReader(m.Value)), nil
}
