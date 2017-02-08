// Package portlock is a simple mutex for use between processes to protect a shared resource
package portlock

import (
	"io"
	"net"
)

// Mutex is used to unlock the lock
type Mutex struct {
	l io.Closer
}

var readBuf [1]byte

// Lock currently uses a tcp connection to determine the lock status, and as
// such requires a tcp address to listen on.
//
// This may change and is not stable.
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

// Unlock removes the lock
func (m *Mutex) Unlock() error {
	return m.l.Close()
}
