package linux

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type LinuxTemp struct{}

func contains(array []string, str string) bool {
	for _, a := range array {
		if a == str {
			return true
		}
	}
	return false
}

func (h *LinuxTemp) GetTemperatures() (map[string]int, error) {
	files, err := ioutil.ReadDir(KERNEL_SYS_THERMAL_CLASS)
	thermal_types := GetThermals()
	if err == nil {
		for _, file := range files {
			thermal_type_path := KERNEL_SYS_THERMAL_CLASS + file.Name() + "/type"
			dbytes, _ := ioutil.ReadFile(thermal_type_path)
			thermal_type := strings.TrimSuffix(string(dbytes), "\n")

			if contains(thermal_types, thermal_type) {
				fmt.Print(file.Name())
				fmt.Println(" " + thermal_type)
			}

		}
		//return nil, errors.New(KERNEL_SYS_THERMAL_CLASS)
	} else {
		fmt.Println(err)
	}
	return map[string]int{"cpu0": 69, "cpu1": 24}, nil
}
