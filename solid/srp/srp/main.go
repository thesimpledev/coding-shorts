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
	user, err := fetchUser(db, "123")
	if err != nil {
		log.Fatal(err)
	}

	printUser(user)
}

func fetchUser(db *sql.DB, id string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func printUser(user User) {
	fmt.Println(user)
}
