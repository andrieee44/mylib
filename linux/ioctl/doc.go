//go:build linux

// Package ioctl implements the userspace api [ioctl.h] in the Linux kernel.
//
// From [ioctl.h]:
//
// ioctl command encoding: 32 bits total, command in lower 16 bits,
// size of the parameter structure in the lower 14 bits of the
// upper 16 bits.
// Encoding the size of the parameter structure in the ioctl request
// is useful for catching programs compiled with old versions
// and to avoid overwriting user space outside the user buffer area.
// The highest 2 bits are reserved for indicating the “access mode”.
// NOTE: This limits the max parameter size to 16kB -1 !
//
// The following is for compatibility across the various Linux
// platforms. The generic ioctl numbering scheme doesn't really enforce
// a type field. De facto, however, the top 8 bits of the lower 16
// bits are indeed used as a type field, so we might just as well make
// this explicit here. Please be sure to use the decoding macros
// below from now on.
//
// [ioctl.h]: https://github.com/torvalds/linux/blob/master/include/uapi/asm-generic/ioctl.h
package ioctl
