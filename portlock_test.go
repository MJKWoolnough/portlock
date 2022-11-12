package portlock

import (
	"runtime"
	"sync"
	"testing"
)

var _ sync.Locker = &mutex{}

func HammerMutex(m Locker, loops int, cdone chan bool) {
	for i := 0; i < loops; i++ {
		if i%3 == 0 {
			if m.TryLock() {
				m.Unlock()
			}
			continue
		}
		m.Lock()
		m.Unlock()
	}
	cdone <- true
}

func TestMutex(t *testing.T) {
	if n := runtime.SetMutexProfileFraction(1); n != 0 {
		t.Logf("got mutexrate %d expected 0", n)
	}
	defer runtime.SetMutexProfileFraction(0)

	m := New(":9999")

	m.Lock()
	if m.TryLock() {
		t.Fatalf("TryLock succeeded with mutex locked")
	}
	m.Unlock()
	if !m.TryLock() {
		t.Fatalf("TryLock failed with mutex unlocked")
	}
	m.Unlock()

	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go HammerMutex(m, 1000, c)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}
