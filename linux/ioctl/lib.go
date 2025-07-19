//go:build linux

package ioctl

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Any performs an ioctl system call on the given file descriptor.
// It wraps the raw [unix.SYS_IOCTL] syscall, passing req as the ioctl
// request code. The arg parameter is an optional pointer to a value of
// type T. If arg is non-nil, its address is sent to the kernel, allowing
// data to be read into or written from *arg. If arg is nil, a zero pointer
// is passed, which is valid for no-data ioctls (e.g [IO]). On success, any
// output data from the kernel is populated into *arg and the error returned
// is nil. On failure, the returned error is the underlying [syscall.Errno].
func Any[T any](fd uintptr, req uint, arg *T) error {
	var errno syscall.Errno

	_, _, errno = unix.Syscall(
		unix.SYS_IOCTL,
		fd,
		uintptr(req),
		uintptr(unsafe.Pointer(arg)),
	)
	if errno != 0 {
		return errno
	}

	return nil
}
