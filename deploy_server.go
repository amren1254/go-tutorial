package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("hello World")
	http.HandleFunc("/",homepage_handler)
	http.HandleFunc("/about",aboutpage_handler)
	http.ListenAndServe("192.168.1.38:8000",nil)
}
func homepage_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hey server created")
}

func aboutpage_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hey Welcome to about page")
}
