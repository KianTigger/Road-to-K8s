package main

import (
	"fmt"
	"time"
)

func main() {

	runIndefinite()
}

func runIndefinite() {
	printLogs := true

	for {
		if printLogs == true {
			fmt.Println(time.Now().String() + " is the current time.")
			time.Sleep(10 * time.Second)
		}
	}
}
