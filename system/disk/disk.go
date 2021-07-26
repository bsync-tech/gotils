package disk

import (
	"bytes"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
)

var EMPTY_STRING = ""

func Get() (string, error) {
	var b bytes.Buffer
	devs1, err := disk.IOCounters()
	if err != nil {
		return EMPTY_STRING, err
	}
	time.Sleep(500 * time.Millisecond)
	devs2, err := disk.IOCounters()
	if err != nil {
		return EMPTY_STRING, err
	}
	for dev, item := range devs2 {
		util := 100 * float64(item.IoTime-devs1[dev].IoTime) / (500)
		b.WriteString(fmt.Sprintf("[dev: %s util: %.2f %%] ", dev, util))
		fmt.Println(item.IoTime, devs1[dev].IoTime)
	}
	return b.String(), nil
}
