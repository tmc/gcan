package gcanpb

import (
	"io"
	"io/ioutil"
)

func (c MessageCompression) NewReader(r io.Reader) io.ReadCloser {
	switch c {
	case Message_ZLIB:
		return zlibReader(r)
	case Message_SNAPPY:
		return snappyReader(r)
	default:
		return ioutil.NopCloser(r)
	}
}

func (c MessageCompression) NewWriter(w io.Writer) io.WriteCloser {
	switch c {
	case Message_ZLIB:
		return zlibWriter(w)
	case Message_SNAPPY:
		return snappyWriter(w)
	default:
		return NopWCloser(w)
	}
}
