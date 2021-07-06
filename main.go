package main

import (
	"fmt"
	"time"

	"log"
	"net/http"
)

func runIndefinite() {
	printLogs := true

	for {
		if printLogs == true {
			fmt.Println(time.Now().String() + " is the current time.")
			time.Sleep(10 * time.Second)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/time", returnTimeAndDate)
	http.HandleFunc("/story", returnStory)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnTimeAndDate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	//json.NewEncoder(w).Encode(Articles)
}

func returnStory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	//json.NewEncoder(w).Encode(Articles)
}

func main() {
	handleRequests()

	//runIndefinite()
}
