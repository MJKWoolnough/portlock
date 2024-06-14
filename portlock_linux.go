package portlock

import (
	"errors"
	"os"
	"syscall"
)

func isOpen(err error) bool {
	var se *os.SyscallError

	return errors.As(err, &se) && errors.Is(se.Err, syscall.EADDRINUSE)
}
