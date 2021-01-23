package windows

//TempWindows - Strategy for temperature gathering for Windows based systems
type TempWindows struct{}

// GetTemperatures - export osx based os temperature information
func (h *TempWindows) GetTemperatures() (map[string]int, error) {
	return map[string]int{"windows0": 12, "windows1": 24}, nil
}
