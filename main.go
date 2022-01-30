package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

// Article structure
type Article struct {
	Title string `json: "title"`
	Banana string `json: "desc"`
	Content string `json: "jkhsdfhds"` // "content"
}

// array of articles
type Articles []Article

// mocking up some API endpoints to retrieve data 
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Banana: "Test Description", Content: "Hello Article"},
		Article{Title: "Real title", Banana: "Real Description", Content: "Hello World!"},
	}

	fmt.Println("endpoint hit: all articles endpoint")
	json.NewEncoder(w).Encode(articles)
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit") // actually see on page
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
	
	
}

func main() {
	handleRequests()
}