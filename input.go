package mylib

// InputEvent identifies a category of input events.
type InputEvent uint

// InputCode identifies a specific event code within an [InputEvent]
// category.
type InputCode uint

// InputDevice represents a physical or virtual input device.
type InputDevice interface {
	// Name is the human-readable name
	// (e.g. "Xbox Controller", "Logitech Dual Action").
	Name() (string, error)

	// ID returns a stable, platform‐specific identifier for this device.
	//
	// On Linux (evdev), it’s formatted as
	// "bus 0x<bustype> vendor 0x<vendor> product 0x<product> version 0x<version>"
	// e.g. "bus 0x3 vendor 0x46d product 0xc24f version 0x111".
	ID() (string, error)

	// Codes returns all supported event codes for the given event category.
	// eventType must be one of the values returned by InputDevice.Events.
	// The result is a slice of InputCode values, each representing
	// a distinct input within that category.
	Codes(eventType InputEvent) ([]InputCode, error)

	// Events returns a slice of supported event categories.
	// Each returned InputEvent value identifies a class of input events
	// your application can monitor (e.g., keys, relative axes).
	// Use InputDevice.Codes with one of these values to enumerate
	// specific codes.
	Events() ([]InputEvent, error)

	// Close releases any underlying resources (file descriptors,
	// threads, etc.) associated with the input device.
	// After Close returns, no other methods should be called.
	Close() error
}
