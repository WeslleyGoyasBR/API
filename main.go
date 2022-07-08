package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json: "desc"`
	Content string `json: "content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "test title", Desc: "test description", Content: "e ae Cloviiiisss"},
		{Title: "estudos de go", Desc: "aborda esse tutorial", Content: "tem como objetivo apenas a simulação"},
		Article{Title: "other test", Desc: "other description", Content: "llll"},
	}

	fmt.Println("Endpoint hit: all articles Endpoint")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(articles)
}

func myRest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my program")

}

func handleRequests() {
	http.HandleFunc("/", myRest)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()

}
