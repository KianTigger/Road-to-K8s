package main

import (
	"fmt"
	"time"
)

func main() {

	go runIndefinite()
}

func runIndefinite() {
	printLogs := true
	for {
		if printLogs == true {
			fmt.Println(time.Now().String() + " is the current time.")
		}
	}
}
