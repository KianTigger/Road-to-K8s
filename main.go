package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}

/*

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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func main() {
	r := mux.NewRouter{}
	r.HandleFunc("/", handler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
	Articles = []Article{
		Article{time: time.Now(), date: time.Now().Format("01-02-2006")},
		Article{success: true, message: "Once upon a time..."},
	}

	handleRequests()

	//runIndefinite()
}

*/
