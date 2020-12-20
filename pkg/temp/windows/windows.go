package windows

type WindowsTemp struct{}

/**
 * GetTemperatures - export osx based os temperature information
 */
func (h *WindowsTemp) GetTemperatures() (map[string]int, error) {
	return map[string]int{"windows0": 12, "windows1": 24}, nil
}
