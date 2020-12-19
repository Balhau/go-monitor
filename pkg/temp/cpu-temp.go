package temp

import (
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	KERNEL_SYS_THERMAL_CLASS = "/sys/class/thermal/"
)

func GetTemp() ([]int, error) {
	files, err := ioutil.ReadDir(KERNEL_SYS_THERMAL_CLASS)
	if err != nil {
		fmt.Println(len(files))
		
		return nil, errors.New(KERNEL_SYS_THERMAL_CLASS)
	}
	return []int{1, 2}, nil
}
