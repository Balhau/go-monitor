package main

import (
	"fmt"

	"git.balhau.net/monitor/pkg/temp"
)

func main() {
	temp, err := temp.NewTemperature()
	if err == nil {
		temps, _ := temp.GetTemperatures()
		fmt.Println(temps)
	}
}
