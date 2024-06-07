package portlock

import (
	"os"
	"syscall"
)

func isOpen(err error) bool {
	if se, ok := err.(*os.SyscallError); ok {
		return se.Err == syscall.EADDRINUSE
	}

	return false
}
