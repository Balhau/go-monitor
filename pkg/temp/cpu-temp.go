package temp

import (
	"fmt"
	"runtime"

	"git.balhau.net/monitor/pkg/temp/windows"

	"git.balhau.net/monitor/pkg/temp/linux"
	"git.balhau.net/monitor/pkg/temp/osx"
	"git.balhau.net/monitor/pkg/utils"
)

// Temperature - Strategies should obey this contract
type Temperature interface {
	GetTemperatures() (map[string]int, error)
}

//NewTemperature - Factory method that return the Temperature instance based on the type of OS
func NewTemperature() (Temperature, error) {
	os := runtime.GOOS
	fmt.Println(os)
	switch os {
	case utils.OsLinux:
		return &linux.TempLinux{}, nil
	case utils.OsOsx:
		return &osx.TempOsx{}, nil
	case utils.OsWindows:
		return &windows.TempWindows{}, nil
	default:
		return nil, fmt.Errorf("No temperature implementation found")
	}
}
