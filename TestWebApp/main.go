package main

import (
	//"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
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
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}

func main(){
	router := mux.NewRouter();
	posts = append(posts, Post{ID: "1", Title: "AwesomePost", Body: "Authentication and Authorization"});
	router.HandleFunc("/posts", getPosts).Methods("GET");
	router.HandleFunc("/posts", createPost).Methods("POST")
	http.ListenAndServe(":5000",router);
}