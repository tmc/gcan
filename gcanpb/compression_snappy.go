package gcanpb

import (
	"io"
	"io/ioutil"

	"github.com/golang/snappy"
)

func snappyReader(r io.Reader) io.ReadCloser {
	return ioutil.NopCloser(snappy.NewReader(r))
}

func snappyWriter(w io.Writer) io.WriteCloser {
	return NopWCloser(snappy.NewWriter(w))
}
