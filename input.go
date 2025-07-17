package mylib

// InputDevice represents a physical or virtual input device.
type InputDevice interface {
	// Name is the human-readable name
	// (e.g. "Xbox Controller", "Logitech Dual Action").
	Name() (string, error)

	// ID returns a stable, platform‐specific identifier for this device.
	//
	// On Linux (evdev), it’s formatted as
	// "<basename>|<vendor>:<product> v<version>",
	// e.g. "event5|045e:028e v0111".
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
