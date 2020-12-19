package linux

import (
	"fmt"
	"io/ioutil"
)

const (
	KERNEL_SYS_THERMAL_CLASS    = "/sys/class/thermal/"
	KERNEL_SYS_THERMAL_TYPE_FAN = "fan"
)

type LinuxTemp struct{}

func (h *LinuxTemp) GetTemperatures() (map[string]int, error) {
	files, err := ioutil.ReadDir(KERNEL_SYS_THERMAL_CLASS)
	if err == nil {
		fmt.Println("FILES")
		fmt.Println(files)
		fmt.Printf("END")
		//return nil, errors.New(KERNEL_SYS_THERMAL_CLASS)
	} else {
		fmt.Println(err)
	}
	return map[string]int{"cpu0": 69, "cpu1": 24}, nil
}
