// Package portlock is a simple mutex for use between processes to protect a shared resource
package portlock // import "vimagination.zapto.org/portlock"

import (
	"io"
	"net"
	"sync"
)

// Mutex is a mutual exclusion lock that can be used across different processes
type mutex struct {
	addr string

	mu sync.Mutex
	l  io.Closer
}

var readBuf [1]byte

// Type Locker combines the sync.Locker interface with the TryLock method
type Locker interface {
	sync.Locker
	TryLock() bool
}

// New creates a new Mutex which currently uses a tcp connection to determine
// the lock status, and as such requires a tcp address to listen on.
//
// This may change and is not stable.
func New(addr string) Locker {
	return &mutex{addr: addr}
}

// Lock locks the mutex. If it is already locked, by this or another process,
// then the call blocks until it is unlocked.
func (m *mutex) Lock() {
	for !m.TryLock() {
		c, err := net.Dial("tcp", m.addr)
		if err == nil {
			c.Read(readBuf[:])
		}
	}
}

// TryLock attempts to lock the Mutex, returning true on a success.
func (m *mutex) TryLock() bool {
	l, err := net.Listen("tcp", m.addr)
	if err == nil {
		m.mu.Lock()
		m.l = l
		m.mu.Unlock()
		return true
	} else if oe, ok := err.(*net.OpError); ok && isOpen(oe.Err) {
		return false
	} else {
		panic(err)
	}
}

// Unlock removes the lock. Due to the current implementation, exiting the
// program will also unlock the mutex.
//
// It is the intention that this will always be true, but Unlock should be
// called before program exit regardless.
func (m *mutex) Unlock() {
	m.mu.Lock()
	m.l.Close()
	m.l = nil
	m.mu.Unlock()
}
