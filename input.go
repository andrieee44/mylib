package mylib

// Device represents a physical or virtual input device.
type Device struct {
	// Name is the human-readable name (e.g. "Xbox Controller",
	// "Logitech Dual Action").
	Name string

	// ID is a platform-specific identifier: on Linux it might be
	// "/dev/input/event5", on Windows it could be a GUID string,
	// and on macOS an IOKit registry path.
	ID string

	// Capabilities describes the features this device supports.
	Capabilities Capabilities
}

// Capabilities describes the feature set supported by an input device.
type Capabilities struct {
	// HasAbsoluteAxes reports whether the device provides absolute
	// axis input (EV_ABS).
	HasAbsoluteAxes bool

	// HasButtons reports whether the device provides button or
	// key input (EV_KEY).
	HasButtons bool

	// IsJoystick reports whether the device is considered a joystick or
	// gamepad. It is true when the device has both absolute axes and
	// buttons.
	IsJoystick bool
}
