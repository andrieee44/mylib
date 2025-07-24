//go:build linux

package main

import "github.com/andrieee44/mylib/linux/input"

var devices []*input.Device = func() []*input.Device {
	var (
		devs []*input.Device
		err  error
	)

	devs, err = input.Devices()
	exitIf(err)

	return devs
}()
