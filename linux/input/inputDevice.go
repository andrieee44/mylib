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

type Device struct {
	file *os.File
}

var _ mylib.InputDevice = (*Device)(nil)

func NewDevice(path string) (Device, error) {
	var (
		device Device
		err    error
	)

	device.file, err = os.Open(filepath.Clean(path))
	if err != nil {
		return Device{}, err
	}

	return device, nil
}

func Devices() ([]Device, error) {
	var (
		devices []Device
		device  Device
		paths   []string
		path    string
		err     error
	)

	paths, err = filepath.Glob("/dev/input/event*")
	if err != nil {
		return nil, err
	}

	devices = make([]Device, 0, len(paths))
	for _, path = range paths {
		device, err = NewDevice(path)
		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (dev *Device) Name() (string, error) {
	var (
		buf []byte
		err error
	)

	buf = make([]byte, 256)

	err = ioctl.Any(dev.file.Fd(), EVIOCGNAME(256), &buf)
	if err != nil {
		return "", fmt.Errorf("Device.Name: %w", err)
	}

	return unix.ByteSliceToString(buf), nil
}

func (dev *Device) ID() (string, error) {
	var (
		buf []byte
		err error
	)

	buf = make([]byte, 256)

	err = ioctl.Any(dev.file.Fd(), EVIOCGNAME(256), &buf)
	if err != nil {
		return "", fmt.Errorf("Device.Name: %w", err)
	}

	return unix.ByteSliceToString(buf), nil
}

func (dev *Device) Close() error {
	return dev.file.Close()
}
