package progress

import (
	"testing"
	"io"
)

func TestProgressReader(t *testing.T) {
	data := []byte{1,2}
	pr, pw := io.Pipe()
	go func(){
		pw.Write(data)
		pw.Close()
	}()
	pgr := NewProgressReader(pr, 2)
	buf := make([]byte, 1)
	pgr.Read(buf)
	if pgr.Progress() != 0.5 {
		t.Error("Expected 0.5 progress, got ", pgr.Progress())
	}
}