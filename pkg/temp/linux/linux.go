package linux

import (
	"io/ioutil"
	"os"
	"strconv"

	"git.balhau.net/monitor/pkg/utils"
)

type LinuxTemp struct{}

/**
*
* GetTemperatures: Return a map with temperature information for linux OS systems
 */
func (h *LinuxTemp) GetTemperatures() (map[string]int, error) {
	var tempMap = make(map[string]int)
	files, err := ioutil.ReadDir(KERNEL_SYS_THERMAL_CLASS)
	thermalTypes := GetThermals()
	if err == nil {
		i := 0
		for _, file := range files {
			thermalType, _ := utils.ReadString(tpath(file, "/type"))
			if utils.Contains(thermalTypes, *thermalType) {
				thermalValueStr, _ := utils.ReadString(tpath(file, "/temp"))
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

func tpath(f os.FileInfo, t string) string {
	return KERNEL_SYS_THERMAL_CLASS + f.Name() + t
}
