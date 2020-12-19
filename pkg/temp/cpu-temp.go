package temp

import (
	"fmt"
	"runtime"

	"git.balhau.net/monitor/pkg/temp/linux"
	"git.balhau.net/monitor/pkg/temp/osx"
)

/** Temperature Strategies should obey this contract */
type Temperature interface {
	GetTemperatures() (map[string]int, error)
}

func NewTemperature() (Temperature, error) {
	os := runtime.GOOS
	switch os {
	case "linux":
		return &linux.LinuxTemp{}, nil
	case "darwin":
		return &osx.OsxTemp{}, nil
	default:
		return nil, fmt.Errorf("No temperature implementation found")
	}
}
