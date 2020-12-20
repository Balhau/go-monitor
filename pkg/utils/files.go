package utils

import (
	"io/ioutil"
	"strings"
)

//ReadString - Read file content and convert bytes into string
func ReadString(path string) (*string, error) {
	dataBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	trimStr := strings.ReplaceAll(string(dataBytes), "\n", "")
	return &trimStr, nil
}
