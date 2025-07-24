// Package main implements the inputdevices CLI, which discovers and displays
// input devices.
//
// It enumerates all available devices, retrieves their ID and name, prints
// the results to standard output, and closes each device handle.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/andrieee44/mylib"
)

func exitIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "inputdevices:", err)
		os.Exit(1)
	}
}

func main() {
	var (
		devs     []mylib.InputDevice
		dev      mylib.InputDevice
		id, name string
		events   []mylib.InputEvent
		event    mylib.InputEvent
		codes    []mylib.InputCode
		code     mylib.InputCode
		builder  strings.Builder
		err      error
	)

	devs = make([]mylib.InputDevice, 0, len(devices))
	for _, dev = range devices {
		devs = append(devs, dev)
	}

	for _, dev = range devs {
		id, err = dev.ID()
		exitIf(err)

		name, err = dev.Name()
		exitIf(err)

		events, err = dev.Events()
		exitIf(err)

		builder.WriteString(fmt.Sprintf("ID: %s\nName: %s\n", id, name))
		builder.WriteString("Supported Events:\n")

		for _, event = range events {
			codes, err = dev.Codes(event)
			exitIf(err)

			builder.WriteString(fmt.Sprintf("  Event Type %d (TBD):\n", event))

			for _, code = range codes {
				builder.WriteString(fmt.Sprintf("    Event code %d (TBD)\n", code))
			}
		}

		err = dev.Close()
		exitIf(err)

		builder.WriteString(strings.Repeat("-", 60))
		builder.WriteByte('\n')
	}

	fmt.Print(builder.String())
}
