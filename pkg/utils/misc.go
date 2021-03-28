package utils

import (
	"encoding/json"
)

//ParseLinuxTemp - Convert linux temp variables into degrees
func ParseLinuxTemp(linuxTemp int) int {
	return linuxTemp / 1000
}

func Jsonfy(v interface{}) (*string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		var s = string(b)
		return &s, err
	}
	return nil, err
}
