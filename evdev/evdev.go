// Package evdev implements input.h in the Linux kernel.
//
//revive:disable:var-naming
package evdev

import (
	"github.com/andrieee44/mylib/ioctl"
)

var (
	// EVIOCGVERSION is the ioctl request code to get the evdev
	// driver version. It reads an int into the provided variable
	// (e.g. 0x010000 == version 1.0.0).
	EVIOCGVERSION = ioctl.IOR('E', 0x01, int(0))

	// EVIOCGID is the ioctl request code to retrieve the device identifier.
	// It reads into an Input_id struct containing Bustype, Vendor, Product,
	// and Version.
	EVIOCGID = ioctl.IOR('E', 0x02, Input_id{})

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
	// keymap entry. It reads into an Input_keymap_entry struct for
	// flags, index, keycode, and scancode.
	EVIOCGKEYCODE_V2 = ioctl.IOR('E', 0x04, Input_keymap_entry{})

	// EVIOCSKEYCODE is the ioctl request code to set a simple keycode
	// mapping. It writes a [2]uint: [0] = scancode, [1] = keycode.
	EVIOCSKEYCODE = ioctl.IOW('E', 0x04, [2]uint{})

	// EVIOCSKEYCODE_V2 is the ioctl request code to set an extended
	// keymap entry. It writes an Input_keymap_entry struct for flags,
	// index, keycode, and scancode.
	EVIOCSKEYCODE_V2 = ioctl.IOW('E', 0x04, Input_keymap_entry{})
)

type Input_id struct{}

type Input_keymap_entry struct{}

func EVIOCGNAME(length uint) uint {
	return ioctl.IOC(ioctl.IOC_READ, 'E', 0x06, length)
}

func EVIOCGPHYS(length uint) uint {
	return ioctl.IOC(ioctl.IOC_READ, 'E', 0x07, length)
}

func EVIOCGUNIQ(length uint) uint {
	return ioctl.IOC(ioctl.IOC_READ, 'E', 0x08, length)
}

func EVIOCGPROP(length uint) uint {
	return ioctl.IOC(ioctl.IOC_READ, 'E', 0x09, length)
}
