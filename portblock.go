package portblock

import (
	"io"
	"net"
)

type Lock struct {
	l io.Closer
}

var readBuf [1]byte

func Lock(addr string) (*Lock, error) {
	for {
		l, err := net.Listen("tcp", addr)
		if err == nil {
			return &Lock{l}, nil
		} else if oe, ok := err.(*net.OpError); ok && isOpen(oe.Err) {
			c, err := net.Dial("tcp", addr)
			if err == nil {
				c.Read(readBuf[:])
			}
		} else {
			return nil, err
		}
	}
}

func (l *Lock) Unlock() error {
	return l.l.Close()
}
