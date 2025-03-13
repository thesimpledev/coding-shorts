package main

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID    string
	Name  string
	Email string
}

func main() {
	db, _ := sql.Open("mysql", "user:password@/dbname")
	fetchAndPrintUser(db, "123")
}

func fetchAndPrintUser(db *sql.DB, id string) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
