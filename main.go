package main

import (
	"assignment1/services"
	"fmt"
	"os"
)

func main() {
	a := os.Args[1]
	process(a, services.Checking)
}

func process(name string, callback func(string) bool) {
	isOnList := callback(name)
	if isOnList {
		services.DataProcess(name)
	} else {
		fmt.Println("not on the list")
	}
}
