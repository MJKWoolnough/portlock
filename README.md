# portlock

[![CI](https://github.com/MJKWoolnough/portlock/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/portlock/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/portlock.svg)](https://pkg.go.dev/vimagination.zapto.org/portlock)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/portlock)](https://goreportcard.com/report/vimagination.zapto.org/portlock)

--
    import "vimagination.zapto.org/portlock"

Package portlock is a simple mutex for use between processes to protect a shared resource.

## Highlights

 - Mutex type implements `sync.Locker`.
 - Adds a `TryLock()` method for non-blocking lock attempts.
 - Can lock across processes, not just goroutines within a process.

## Usage

```go
package main

import (
	"fmt"

	"vimagination.zapto.org/portlock"
)

func main() {
	mu := portlock.New("127.0.0.1:9999")

	if !mu.TryLock() {
		fmt.Println("Failed to lock mutex")

		return
	}
	fmt.Println("Mutex locked")

	mu.Unlock()
	fmt.Println("Mutex unlocked")

	// Output:
	// Mutex locked
	// Mutex unlocked
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/portlock
