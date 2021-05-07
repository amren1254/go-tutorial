package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"	
)

func main(){
	router := mux.NewRouter();
	router.HandleFunc("/test", test).Methods("GET");
	router.HandleFunc("/add", addItem).Methods("POST");

	http.ListenAndServe(":8090", router);
}

func test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	json.NewEncoder(w).Encode(struct {
		ID string
	}{"213"})
}

type Users struct{
	FullName string `json:"name"`
	UserName string `json:"username"`
	Email string `json:"email"`

}
//
type Posts struct{
	Title string `json:"title"`
	Body string `json:"body`
	Author Users `json:"author"`
}
var posts []Posts = []Posts{}

func addItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	//routeVariable := mux.Vars(r)["item"];
	var newPost Posts;
	json.NewDecoder(r.Body).Decode(&newPost);
	posts = append(posts, newPost);
	json.NewEncoder(w).Encode(posts);
}	