package mysqldumper

import (
	"bufio"
	"io"
)

type GzipWriter struct {
	wr *bufio.Writer
}

func (s *GzipWriter) Write(data string) error {
	_, err := s.wr.Write([]byte(data))

	return err
}

func (s *GzipWriter) Flush() error {
	return s.wr.Flush()
}

func NewGzipWriter(w io.Writer) *GzipWriter {
	return &GzipWriter{wr: bufio.NewWriter(w)}
}
