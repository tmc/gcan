package gcanpb

import "io"

type nopWCloser struct {
	io.Writer
}

func (nopWCloser) Close() error { return nil }

// NopWCloser returns a WriteCloser with a no-op Close method wrapping
// the provided Writer w.
func NopWCloser(w io.Writer) io.WriteCloser {
	return nopWCloser{w}
}
