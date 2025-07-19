//go:build linux

// Package input implements the userspace api [input.h] in the Linux kernel.
// It also implements [mylib.InputDevice].
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
package input
