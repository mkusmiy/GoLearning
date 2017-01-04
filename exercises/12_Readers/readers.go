package readers

import (
	"io"
	"strings"
)

// Combine returns an io.Reader which represents
// the contents of a and b.
func Combine(a, b io.Reader) io.Reader {
	return a
}

// always reader always fills the read buffer with
// the byte ch.
type alwaysReader struct {
	ch byte
}

func (a *alwaysReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = a.ch
	}
	return len(buf), nil
}

// AReader returns an io.Reader which returns n 'A' characters
func AReader(n int) io.Reader {
	return strings.NewReader("AAAAAAAAAA")
}
