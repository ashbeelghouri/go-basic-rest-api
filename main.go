package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TYPES STRUCT

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article


func allArticles(w http.ResponseWriter, r *http.Request){

	articles := Articles{
		Article{
			Title: "Test title",
			Desc: "Test Description",
			Content: "This is a simple content",
		},
	}
	json.NewEncoder(w).Encode(articles)
}


func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homepage)
	router.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	handleRequests()
}