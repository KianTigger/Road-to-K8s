package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Success bool     `json:"success"`
	Message string   `json:"message"`
}

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
	fmt.Fprintf(w, "Welcome to the base homepage.")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/story", createArticle).Methods("POST")
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/{id}", returnSingleArticle)
	myRouter.HandleFunc("/", homePage)

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

func createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Point Hit: create article")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	article.Success = true
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func main() {

	//runIndefinite()

	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "time", Title: "Time", Desc: "Current time and date", Time: typetime(time.Now()), Date: (time.Now().Format("01-02-2008"))},
		{Id: "story2", Title: "Story", Desc: "Article Description 2", Content: "Article Content 2"},
	}

	handleRequests()
}
