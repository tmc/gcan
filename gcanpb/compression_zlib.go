package gcanpb

import (
	"compress/zlib"
	"io"
)

func zlibReader(r io.Reader) io.ReadCloser {
	rdr, err := zlib.NewReader(r)
	if err != nil {
		panic(err)
	}
	return rdr
}

func zlibWriter(w io.Writer) io.WriteCloser {
	return zlib.NewWriter(w)
}
