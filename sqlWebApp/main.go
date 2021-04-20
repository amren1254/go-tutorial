package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/goRestService")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var (
		id int
		name string
		password string
	)
	//select 
	rows, err := db.Query("select id, username, password from user")// where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &password)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, password)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	insert
	stmt, err := db.Prepare("INSERT INTO user(username,password) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("naren","naren@")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}