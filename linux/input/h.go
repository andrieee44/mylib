//go:build linux

// Package input implements [input.h] in the Linux kernel.
// It also implements [mylib.InputDevice].
//
// [input.h]: https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h
package input

import "github.com/andrieee44/mylib/linux/ioctl"

// ID identifies an input device by its bus type, vendor ID, product ID,
// and version. It mirrors the struct input_id definition used by Linux input
// ioctls.
type ID struct {
	// Bustype is the bus type for the device (for example, BUS_USB or
	// BUS_SPI).
	Bustype uint16

	// Vendor is the vendor identifier assigned by the bus.
	Vendor uint16

	// Product is the product identifier assigned by the vendor.
	Product uint16

	// Version is the version or revision number of the device.
	Version uint16
}

// KeymapEntry is used by the EVIOCGKEYCODE and EVIOCSKEYCODE ioctls
// to retrieve or modify keymap data. Lookup may be performed either
// by the Scancode itself or by the Index in the keymap entry.
// When using EVIOCGKEYCODE, the kernel will return either Scancode
// or Index, depending on which field was used for the lookup.
type KeymapEntry struct {
	// Flags specifies how the kernel should handle the request.
	// Setting INPUT_KEYMAP_BY_INDEX instructs the kernel to look up
	// by Index instead of Scancode.
	Flags uint8

	// Len is the length of the scancode that resides in the Scancode buffer.
	Len uint8

	// Index is the keymap index; may be used instead of Scancode.
	Index uint16

	// Keycode is the key code assigned to this scancode.
	Keycode uint32

	// Scancode is the scancode represented in machine-endian form.
	Scancode [32]uint8
}

// INPUT_KEYMAP_BY_INDEX is a flag for the EVIOCGKEYCODE_V2/EVIOCSKEYCODE_V2
// ioctls. It tells the kernel to identify the keymap entry by its Index
// field. When set, the ioctl uses [KeymapEntry.Index] to select which key
// mapping to get or set.
const INPUT_KEYMAP_BY_INDEX = 1 << 0

var (
	// EVIOCGVERSION is the ioctl request code to get the evdev
	// driver version. It reads an int into the provided variable
	// (e.g. 0x010000 == version 1.0.0).
	EVIOCGVERSION = ioctl.IOR('E', 0x01, int(0))

	// EVIOCGID is the ioctl request code to retrieve the device identifier.
	// It reads into an [ID] struct containing Bustype, Vendor, Product,
	// and Version.
	EVIOCGID = ioctl.IOR('E', 0x02, ID{})

	// EVIOCGREP is the ioctl request code to get keyboard auto‐repeat
	// settings. It reads a [2]uint: [0] = delay in ms, [1] = period in ms.
	EVIOCGREP = ioctl.IOR('E', 0x03, [2]uint{})

	// EVIOCSREP is the ioctl request code to set keyboard auto‐repeat
	// settings. It writes a [2]uint: [0] = delay in ms, [1] = period in ms.
	EVIOCSREP = ioctl.IOW('E', 0x03, [2]uint{})

	// EVIOCGKEYCODE is the ioctl request code to get a simple keycode
	// mapping. It reads a [2]uint: [0] = scancode, [1] = keycode.
	EVIOCGKEYCODE = ioctl.IOR('E', 0x04, [2]uint{})

	// EVIOCGKEYCODE_V2 is the ioctl request code to get an extended
	// keymap entry. It reads into an [KeymapEntry] struct for
	// flags, index, keycode, and scancode.
	EVIOCGKEYCODE_V2 = ioctl.IOR('E', 0x04, KeymapEntry{})

	// EVIOCSKEYCODE is the ioctl request code to set a simple keycode
	// mapping. It writes a [2]uint: [0] = scancode, [1] = keycode.
	EVIOCSKEYCODE = ioctl.IOW('E', 0x04, [2]uint{})

	// EVIOCSKEYCODE_V2 is the ioctl request code to set an extended
	// keymap entry. It writes an [KeymapEntry] struct for flags,
	// index, keycode, and scancode.
	EVIOCSKEYCODE_V2 = ioctl.IOW('E', 0x04, KeymapEntry{})
)

// EVIOCGNAME returns the ioctl request code to retrieve the device name.
// The length parameter specifies the size of the buffer (in bytes) that
// will hold the returned name string.
func EVIOCGNAME(length uint) uint {
	return ioctl.IOC(ioctl.IOC_READ, 'E', 0x06, length)
}

// EVIOCGPHYS returns the ioctl request code to retrieve the device
// physical location path. The length parameter specifies the size of the
// buffer (in bytes) that will hold the returned physical path string.
func EVIOCGPHYS(length uint) uint {
	return ioctl.IOC(ioctl.IOC_READ, 'E', 0x07, length)
}

// EVIOCGUNIQ returns the ioctl request code to retrieve the device’s
// unique identifier. The length parameter specifies the size of the
// buffer (in bytes) that will hold the returned unique ID string.
func EVIOCGUNIQ(length uint) uint {
	return ioctl.IOC(ioctl.IOC_READ, 'E', 0x08, length)
}

// EVIOCGPROP returns the ioctl request code to retrieve the device’s
// property bitmask. The length parameter specifies the size of the
// buffer (in bytes) that will hold the returned bitmask.
func EVIOCGPROP(length uint) uint {
	return ioctl.IOC(ioctl.IOC_READ, 'E', 0x09, length)
}
