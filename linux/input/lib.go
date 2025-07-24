//go:build linux

package input

import (
	"errors"

	"github.com/andrieee44/mylib"
)

// ErrInvalidEventType is returned when an unsupported or unrecognized
// event type is passed to a Device method.
var ErrInvalidEventType error = errors.New("invalid event type")

// TestBit returns true if the bit numbered pos is set in b.
func TestBit(b []byte, pos uint) bool {
	return b[pos/8]&(1<<(pos%8)) != 0
}

// MaxCodes returns the highest valid code for the specified eventType.
// It looks up eventType in a predefined map of EV_* constants to their
// *_MAX values. If eventType is supported, it returns (maxCode, true).
// Otherwise it returns (0, false).
func MaxCodes(eventType mylib.InputEvent) (uint, bool) {
	var (
		maxCodes map[mylib.InputEvent]uint
		maxCode  uint
		ok       bool
	)

	maxCodes = map[mylib.InputEvent]uint{
		EV_SYN:       SYN_MAX,
		EV_KEY:       KEY_MAX,
		EV_REL:       REL_MAX,
		EV_ABS:       ABS_MAX,
		EV_MSC:       MSC_MAX,
		EV_SW:        SW_MAX,
		EV_LED:       LED_MAX,
		EV_SND:       SND_MAX,
		EV_REP:       REP_MAX,
		EV_FF:        FF_MAX,
		EV_PWR:       0,
		EV_FF_STATUS: FF_STATUS_MAX,
	}

	maxCode, ok = maxCodes[eventType]

	return maxCode, ok
}
