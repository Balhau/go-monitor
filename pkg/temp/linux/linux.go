package linux

import (
	"io/ioutil"
	"os"
	"strconv"

	"git.balhau.net/monitor/pkg/utils"
)

// LinuxTemp - Struct that represents the temperature strategy for linux system
type LinuxTemp struct{}

/**
*
* GetTemperatures: Return a map with temperature information for linux OS systems
 */
func (h *LinuxTemp) GetTemperatures() (map[string]int, error) {
	var tempMap = make(map[string]int)
	files, err := ioutil.ReadDir(KernelSysThermalClass)
	thermalTypes := GetThermals()
	if err == nil {
		i := 0
		for _, file := range files {
			thermalType, _ := utils.ReadString(Path(file, "/type"))
			if utils.Contains(thermalTypes, *thermalType) {
				thermalValueStr, _ := utils.ReadString(Path(file, "/temp"))
				thermalValue, err := strconv.Atoi(*thermalValueStr)
				if err == nil {
					tempMap[*thermalType+"_"+strconv.Itoa(i)] = utils.ParseLinuxTemp(thermalValue)
					i++
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
