package osx

//TempOsx - Strategy for temperature gathering for OSX based systems
type TempOsx struct{}

// GetTemperatures - export osx based os temperature information
func (h *TempOsx) GetTemperatures() (map[string]int, error) {
	return map[string]int{"osx0": 12, "osx1": 24}, nil
}
