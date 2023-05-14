package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

// User represents a user entity.
type User struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	//Set up database connection.
	db, err := sqlx.Open("postgres", "postgres://postgres:postgres@cicd_db_1:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Run database migrations.
	if err := goose.Up(db.DB, "../migration"); err != nil {
		log.Print(err)
	}

	// Create a new user.
	newUser := User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	_, err = db.NamedExec(`
        INSERT INTO users (name, email)
        VALUES (:name, :email)
    `, &newUser)
	if err != nil {
		panic(err)
	}

	//Get all users.
	var users []User
	err = db.Select(&users, `SELECT * FROM users`)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}
