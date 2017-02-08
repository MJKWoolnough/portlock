package portlock

import (
	"io"
	"net"
)

type Mutex struct {
	l io.Closer
}

var readBuf [1]byte

func Lock(addr string) (*Mutex, error) {
	for {
		l, err := net.Listen("tcp", addr)
		if err == nil {
			return &Mutex{l}, nil
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

func (m *Mutex) Unlock() error {
	return m.l.Close()
}
