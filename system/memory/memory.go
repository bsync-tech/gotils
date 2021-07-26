package memory

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v3/mem"
)

var EMPTY_STRING = ""

// Get memory statistics
func Get() (string, error) {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return EMPTY_STRING, err
	}
	s := fmt.Sprintf("total memory: %s, used memory: %s", humanize.Bytes(mem.Total), humanize.Bytes(mem.Used))
	return s, nil
}
