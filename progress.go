package progress

import (
	"io"
)

type ProgressReader struct {
	Input io.Reader
	Size int64
	Finished int64
	OnProgress func(float32)
}

func NewProgressReader(input io.Reader, size int64) *ProgressReader {
	p := new(ProgressReader)
	p.Input = input
	p.Size = size
	p.Finished = 0
	return p
}

func (c *ProgressReader) Read(p []byte) (n int, err error) {
    cBytes, err := c.Input.Read(p)
    c.Finished = c.Finished + int64(cBytes)
	if err == nil && cBytes != 0 && c.OnProgress != nil {
		c.OnProgress(c.Progress())
	}
	return cBytes, err
}

func (c ProgressReader) Progress() (n float32) {
    return float32(c.Finished)/float32(c.Size)
}