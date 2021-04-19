package main

import (
	//"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	//"math/rand"
	"net/http"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(posts);
}

func main(){
	router := mux.NewRouter();
	posts = append(posts, Post{ID: "1", Title: "AwesomePost", Body: "Authentication and Authorization"});
	router.HandleFunc("/posts", getPosts).Methods("GET");
	http.ListenAndServe(":5000",router);
}