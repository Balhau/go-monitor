package osx

type OsxTemp struct{}

/**
 * GetTemperatures - export osx based os temperature information
 */
func (h *OsxTemp) GetTemperatures() (map[string]int, error) {
	return map[string]int{"osx0": 12, "osx1": 24}, nil
}
