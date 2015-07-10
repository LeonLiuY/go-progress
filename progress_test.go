package progress

import (
	"testing"
	"bytes"
	"io"
)

func TestProgressReader(t *testing.T) {
	pr := NewProgressReader(nil, 100)
	buf := new(bytes.Buffer)
	io.Copy(buf, pr)
}