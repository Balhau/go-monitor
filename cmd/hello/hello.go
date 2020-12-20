package main

import (
	"encoding/json"
	"fmt"

	"git.balhau.net/monitor/pkg/temp"
)

func main() {
	temp, err := temp.NewTemperature()
	if err == nil {
		temps, _ := temp.GetTemperatures()
		fmt.Print("Array Size: ")
		fmt.Println(len(temps))
		jsonTemperatures, err := json.Marshal(temps)
		if err == nil {
			fmt.Println(string(jsonTemperatures))
		}

	}
}
