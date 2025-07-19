package mylib

// InputDevice represents a physical or virtual input device.
type InputDevice interface {
	// Name is the human-readable name
	// (e.g. "Xbox Controller", "Logitech Dual Action").
	Name() (string, error)

	// ID returns a stable, platform‐specific identifier for this device.
	//
	// On Linux (evdev), it’s formatted as
	// "bus <bustype> vendor <vendor> product <product> version <version>"
	// e.g. "bus 0x3 vendor 0x46d product 0xc24f version 0x111".
	ID() string

	// HasAbsoluteAxes reports whether the device provides absolute
	// axis input.
	HasAbsoluteAxes() bool

	// HasButtons reports whether the device provides button or
	// key input.
	HasButtons() bool

	Subscribe(events chan<- struct{})

	Close() error
}
