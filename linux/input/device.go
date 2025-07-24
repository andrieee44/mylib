//go:build linux

package input

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andrieee44/mylib"
	"github.com/andrieee44/mylib/linux/ioctl"
	"golang.org/x/sys/unix"
)

// Device represents an evdev input device.
// It wraps the opened /dev/input/eventN file.
type Device struct {
	file *os.File
	fd   uintptr
}

var _ mylib.InputDevice = (*Device)(nil)

// NewDevice opens the evdev device at the given path and returns a Device.
// The path is cleaned before opening, and the device file is opened
// in read-write mode. The caller is responsible for closing the device
// when no longer needed.
func NewDevice(path string) (*Device, error) {
	var (
		device *Device
		file   *os.File
		err    error
	)

	file, err = os.OpenFile(filepath.Clean(path), os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("input.NewDevice: %w", err)
	}

	device = &Device{
		file: file,
		fd:   file.Fd(),
	}

	return device, nil
}

// Devices scans /dev/input for event devices, opens each one, and
// returns a slice of Device pointers. If any device fails to open,
// an error is returned and no devices are returned.
func Devices() ([]*Device, error) {
	var (
		devices []*Device
		device  *Device
		paths   []string
		path    string
		err     error
	)

	paths, err = filepath.Glob("/dev/input/event*")
	if err != nil {
		return nil, fmt.Errorf("input.Devices: %w", err)
	}

	devices = make([]*Device, 0, len(paths))
	for _, path = range paths {
		device, err = NewDevice(path)
		if err != nil {
			return nil, fmt.Errorf("input.Devices: %w", err)
		}

		devices = append(devices, device)
	}

	return devices, nil
}

// Name returns the human-readable name of the evdev device.
// It sends the [EVIOCGNAME] ioctl to read up to 256 bytes and
// converts the null-terminated result into a Go string.
func (dev *Device) Name() (string, error) {
	var (
		buf []byte
		err error
	)

	buf = make([]byte, 256)

	err = ioctl.Any(dev.fd, EVIOCGNAME(256), &buf[0])
	if err != nil {
		return "", fmt.Errorf("Device.Name: %w", err)
	}

	return unix.ByteSliceToString(buf), nil
}

// ID returns the platform-specific identifier for this evdev device.
// It issues the EVIOCGID ioctl to fetch the bus, vendor, product, and version fields.
// The result is formatted as:
// "bus 0x<bustype> vendor 0x<vendor> product 0x<product> version 0x<version>".
// e.g. "bus 0x3 vendor 0x46d product 0xc24f version 0x111".
func (dev *Device) ID() (string, error) {
	var (
		id  ID
		err error
	)

	err = ioctl.Any(dev.fd, EVIOCGID, &id)
	if err != nil {
		return "", fmt.Errorf("Device.ID: %w", err)
	}

	return fmt.Sprintf(
		"bus 0x%x vendor 0x%x product 0x%x version 0x%x",
		id.Bustype,
		id.Vendor,
		id.Product,
		id.Version,
	), nil
}

// Events returns a slice of all supported event types for the device.
func (dev *Device) Events() ([]mylib.InputEvent, error) {
	var (
		buf       []byte
		events    []mylib.InputEvent
		eventType mylib.InputEvent
		err       error
	)

	buf = make([]byte, (EV_MAX+7)/8)

	err = ioctl.Any(
		dev.fd,
		EVIOCGBIT(0, uint(len(buf))),
		&buf[0],
	)
	if err != nil {
		return nil, fmt.Errorf("Device.Events: %w", err)
	}

	events = make([]mylib.InputEvent, 0, EV_CNT)

	for eventType = range EV_CNT {
		if !TestBit(buf, uint(eventType)) {
			continue
		}

		if eventType == EV_REP {
			continue
		}

		events = append(events, eventType)
	}

	return events, nil
}

// Codes returns all supported [mylib.InputCode] values for the given
// eventType.
func (dev *Device) Codes(eventType mylib.InputEvent) ([]mylib.InputCode, error) {
	var (
		buf            []byte
		codes          []mylib.InputCode
		maxCodes, code uint
		ok             bool
		err            error
	)

	maxCodes, ok = MaxCodes(eventType)
	if !ok {
		return nil, fmt.Errorf("Device.Codes: %w %d", ErrInvalidEventType, eventType)
	}

	buf = make([]byte, (maxCodes+7)/8)

	err = ioctl.Any(
		dev.fd,
		EVIOCGBIT(uint(eventType), uint(len(buf))),
		&buf[0],
	)
	if err != nil {
		return nil, fmt.Errorf("Device.Codes: %w", err)
	}

	codes = make([]mylib.InputCode, 0, maxCodes+1)

	for code = range maxCodes + 1 {
		if !TestBit(buf, code) {
			continue
		}

		codes = append(codes, mylib.InputCode(code))
	}

	return codes, nil
}

// Close closes the evdev device by closing its underlying file handle.
func (dev *Device) Close() error {
	var err error

	err = dev.file.Close()
	if err != nil {
		return fmt.Errorf("Device.Close: %w", err)
	}

	return nil
}
