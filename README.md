# portlock
--
    import "vimagination.zapto.org/portlock"

Package portlock is a simple mutex for use between processes to protect a shared
resource.

## Usage

#### type Locker

```go
type Locker interface {
	sync.Locker
	TryLock() bool
}
```

Type Locker combines the sync.Locker interface with the TryLock method.

#### func  New

```go
func New(addr string) Locker
```
New creates a new Mutex which currently uses a tcp connection to determine the
lock status, and as such requires a tcp address to listen on.

This may change and is not stable.
