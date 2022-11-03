package main

import (
	"assignment1/models"
	"fmt"
	"os"
)

func main() {
	a := os.Args[1]
	process(a, models.Checking)
}

func process(name string, callback func(string) bool) {
	isOnList := callback(name)
	if isOnList {
		models.DataProcess(name)
	} else {
		fmt.Println("not on the list")
	}
}
