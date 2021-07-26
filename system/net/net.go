package net

import (
	"bytes"
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v3/net"
)

var EMPTY_STRING = ""
var SLEEP_TIME = time.Second * 1

// Get network statistics
func Get() (string, error) {
	var b bytes.Buffer
	devs1, err := net.IOCounters(true)
	if err != nil {
		return EMPTY_STRING, err
	}
	time.Sleep(SLEEP_TIME)
	devs2, err := net.IOCounters(true)
	if err != nil {
		return EMPTY_STRING, err
	}
	for dev, item := range devs2 {
		read := (item.BytesRecv - devs1[dev].BytesRecv) / (1)
		send := (item.BytesSent - devs1[dev].BytesSent) / (1)
		b.WriteString(fmt.Sprintf("[dev: %s read: %s/s write: %s/s] ", item.Name, humanize.Bytes(read), humanize.Bytes(send)))
	}
	return b.String(), nil
}
