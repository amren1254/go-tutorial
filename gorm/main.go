package main

import (
    "fmt"
    "log"
    "net/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
    "github.com/gorilla/mux"
	"database/sql"
	"time"
	"encoding/json"
)
type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
  }
  type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
  }
var user User
var users []User  

func allUsers(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "All Users Endpoint Hit")
	dsn := "host=localhost user=drax password=drax dbname=netbanking port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//defer db.Close();
	var users []User
    db.Find(&users)
    fmt.Println("{}", users)

    json.NewEncoder(w).Encode(users)
	//result := db.First(&user)
	      
}

func newUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "New User Endpoint Hit")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Update User Endpoint Hit")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/users", allUsers).Methods("GET")
    myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
    myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
    myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
    log.Fatal(http.ListenAndServe(":8081", myRouter))
}


func main() {
    fmt.Println("Go ORM Tutorial")

    // Handle Subsequent requests
    handleRequests()
}