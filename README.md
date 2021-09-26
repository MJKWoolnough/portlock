# portlock
--
    import "vimagination.zapto.org/portlock"

Package portlock is a simple mutex for use between processes to protect a shared
### resource

## Usage

#### func  New

```go
func New(addr string) sync.Locker
```
New creates a new Mutex which currently uses a tcp connection to determine the
lock status, and as such requires a tcp address to listen on.

This may change and is not stable.
