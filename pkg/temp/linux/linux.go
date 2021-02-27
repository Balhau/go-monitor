package linux

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"git.balhau.net/monitor/pkg/utils"
)

// TempLinux - Struct that represents the temperature strategy for linux system
type TempLinux struct{}

// GetTemperatures - Return a map with temperature information for linux OS systems
func (h *TempLinux) GetTemperatures() (map[string]int, error) {
	var tempMap = make(map[string]int)
	fmt.Println(KernelSysThermalClass)
	files, err := ioutil.ReadDir(KernelSysThermalClass)
	thermalTypes := GetThermals()
	if err == nil {
		i := 0
		for _, file := range files {
			thermalType, _ := utils.ReadString(Path(file, "/type"))
			if utils.Contains(thermalTypes, *thermalType) {
			thermalValueStr, err := utils.ReadString(Path(file, "/temp"))
			if err == nil {
				thermalValue, err := strconv.Atoi(*thermalValueStr)
				if err == nil {
					tempMap[*thermalType+"_"+strconv.Itoa(i)] = utils.ParseLinuxTemp(thermalValue)
					i++
				}
			}
			}
		}
	} else {
		return nil, err
	}
	return tempMap, nil
}

//Path - Utility method to build thermal file descriptor paths
func Path(f os.FileInfo, t string) string {
	return KernelSysThermalClass + f.Name() + t
}
