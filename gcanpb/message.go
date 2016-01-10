package gcanpb

import (
	"encoding/binary"
	"hash/crc32"
)

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
