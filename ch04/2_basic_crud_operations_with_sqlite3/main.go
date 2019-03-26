package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Book struct {
	id int
	name string
	author string
}

func main(){
	dbClient, dbErr := sql.Open("sqlite3", "./books.db")
	if dbErr != nil{
		log.Fatalf("Error while connect to Database: %v\n", dbErr)
		os.Exit(1)
	}
	log.Println(dbClient)
	defer dbClient.Close()

	// Create Table
	statement, err := dbClient.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil{
		log.Fatalf("error while prepare create statement: %v\n", err)
	} else {
		log.Println("successfully created table books!")
	}
	statement.Exec()

	// Insert Data
	statement, err = dbClient.Prepare("INSERT INTO books(name, author, isbn) VALUES(?,?,?)")
	if err != nil{
		log.Fatalf("error while insert data to database: %v", err)
	} else {
		log.Println("successfully insert data.")
	}
	statement.Exec("Detective Conan", "Aoyama Gosho", 12345)

	// Read Data
	rows, err := dbClient.Query("SELECT id, name, author FROM books")
	var tempBook Book
	for rows.Next(){
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("%d: Title %s, By %s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	// Update Data
	statement, err = dbClient.Prepare("update books set name=? where id=?")
	if err != nil{
		log.Fatalf("error while update data to database: %v", err)
	} else {
		log.Println("successfully update data.")
	}
	statement.Exec("Detective Jonen", 11)
}
