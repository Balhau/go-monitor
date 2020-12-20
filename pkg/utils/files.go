package utils

import (
	"io/ioutil"
	"strings"
)

func ReadString(path string) (*string, error) {
	dataBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	trimStr := strings.ReplaceAll(string(dataBytes), "\n", "")
	return &trimStr, nil
}
