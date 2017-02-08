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

Mutex is used to unlock the lock

#### func  Lock

```go
func Lock(addr string) (*Mutex, error)
```
Lock currently uses a tcp connection to determine the lock status, and as such
requires a tcp address to listen on.

This may change and is not stable.

#### func (*Mutex) Unlock

```go
func (m *Mutex) Unlock() error
```
Unlock removes the lock
