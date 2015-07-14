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
	prg := float32(0.0)
	pgr := NewProgressReader(pr, 2)
	pgr.OnProgress = func(p float32) {
		prg = p
	}
	buf := make([]byte, 1)
	pgr.Read(buf)
	if pgr.Progress() != 0.5 {
		t.Error("Expected 0.5 progress, got ", pgr.Progress())
	}
	if prg != 0.5 {
		t.Error("Expected callback 0.5 progress, got ", prg)
	}
}