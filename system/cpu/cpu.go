// +build linux

package cpu

import (
	"time"

	"github.com/bsync-tech/gotils/convert"
	"github.com/shirou/gopsutil/v3/cpu"
)

var EMPTY_STRING = ""

// Get cpu statistics
func Get() (string, error) {
	c, err := cpu.Percent(time.Second, false)
	if err != nil {
		return EMPTY_STRING, err
	}
	cu, err := convert.JSONString(c)
	if err != nil {
		return EMPTY_STRING, err
	}
	return cu, nil
}
