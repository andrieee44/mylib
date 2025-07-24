//go:build linux

// Package input implements the userspace api [input.h] and event constants
// in [input-event-codes.h] in the Linux kernel.
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
// [input-event-codes.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h
package input
