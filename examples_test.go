package portlock_test

import (
	"fmt"

	"vimagination.zapto.org/portlock"
)

func Example() {
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
