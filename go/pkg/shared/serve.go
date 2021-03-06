package shared

import (
	"io"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// rwCloser merges a ReadCloser and a WriteCloser into a ReadWriteCloser.
type rwCloser struct {
	io.ReadCloser
	io.WriteCloser
}

func (rw rwCloser) Close() error {
	err := rw.ReadCloser.Close()
	if err := rw.WriteCloser.Close(); err != nil {
		return err
	}
	return err
}

func Serve() {
	s := rpc.NewServer()
	if err := s.Register(S3Repository{}); err != nil {
		panic(err)
	}
	if err := s.Register(GCSRepository{}); err != nil {
		panic(err)
	}
	if err := s.Register(DiskRepository{}); err != nil {
		panic(err)
	}
	s.ServeCodec(jsonrpc.NewServerCodec(rwCloser{os.Stdin, os.Stdout}))
}
