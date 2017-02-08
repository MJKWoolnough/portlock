# portlock
--
    import "github.com/MJKWoolnough/portlock"

Package portlock is a simple mutex for use between processes to protect a shared
### resource

## Usage

#### type Mutex

```go
type Mutex struct {
}
```

Mutex is a mutual exclusion lock that can be used across different processes

#### func  New

```go
func New(addr string) *Mutex
```
New creates a new Mutex which currently uses a tcp connection to determine the
lock status, and as such requires a tcp address to listen on.

This may change and is not stable.

#### func (*Mutex) Lock

```go
func (m *Mutex) Lock()
```
Lock locks the mutex. If it is already locked, by this or another process, then
the call blocks until it is unlocked.

#### func (*Mutex) Unlock

```go
func (m *Mutex) Unlock()
```
Unlock removes the lock. Due to the current implementation, exiting the program
will also unlock the mutex.

It is the intention that this will always be true, but Unlock should be called
before program exit regardless.
