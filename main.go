package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type typetime time.Time
type typedate time.Time
type Article struct {
	Id      string   `json:"Id"`
	Title   string   `json:"Title"`
	Desc    string   `json:"desc"`
	Content string   `json:"content"`
	Time    typetime `json:"time"`
	Date    string   `json:"date"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

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
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/{id}", returnSingleArticle)
	myRouter.HandleFunc("/", homePage)
	/*myRouter.HandleFunc("/time", returnTime)
	myRouter.HandleFunc("/story", returnStory)
	myRouter.HandleFunc("/all", returnAllArticles)*/
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func returnTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnTime")
	json.NewEncoder(w).Encode(Articles)
}

func returnStory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnStory")
	json.NewEncoder(w).Encode(Articles)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "time", Title: "Time", Desc: "Current time and date", Time: typetime(time.Now()), Date: (time.Now().Format("01-02-2008"))},
		{Id: "story", Title: "Story", Desc: "Article Description 2", Content: "Article Content 2"},
	}

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
