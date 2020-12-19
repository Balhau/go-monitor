package main

import (
	"fmt"

	"git.balhau.net/monitor/pkg/temp"
)

func main() {
	fmt.Println("Hello world")
	var temps []int
	var err error
	temps, err = temp.GetTemp()
	if err != nil {
		fmt.Println("Deu merda")
	} else {
		fmt.Println(temps[0])
	}
}
