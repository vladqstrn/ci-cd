package main

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestUserCRUD(t *testing.T) {
	// Connect to test database.
	db, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Run migrations.
	// if err := goose.Up(db.DB, "../migration"); err != nil {
	// 	t.Fatal(err)
	// }

	// Create a new user.
	newUser := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	_, err = db.NamedExec(`
        INSERT INTO users (name, email)
        VALUES (:name, :email)
    `, &newUser)
	if err != nil {
		t.Fatal(err)
	}

	// Get the user by ID.
	var retrievedUser User
	err = db.Get(&retrievedUser, `SELECT * FROM users WHERE id=1`)
	if err != nil {
		t.Fatal(err)
	}
	if retrievedUser.Name != newUser.Name || retrievedUser.Email != newUser.Email {
		t.Errorf("Expected user %+v, but got %+v", newUser, retrievedUser)
	}

	// Update the user.
	updatedUser := User{
		ID:    newUser.ID,
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	}
	_, err = db.NamedExec(`
        UPDATE users SET name=:name, email=:email WHERE id=:id
    `, &updatedUser)
	if err != nil {
		t.Fatal(err)
	}

	// Get the updated user by ID.
	var retrievedUpdatedUser User
	err = db.Get(&retrievedUpdatedUser, `SELECT * FROM users WHERE id=1`)
	if err != nil {
		t.Fatal(err)
	}
	if retrievedUpdatedUser.Name != updatedUser.Name || retrievedUpdatedUser.Email != updatedUser.Email {
		t.Errorf("Expected user %+v, but got %+v", updatedUser, retrievedUpdatedUser)
	}

	// Delete the user.
	_, err = db.Exec(`DELETE FROM users WHERE id=$1`, newUser.ID)
	if err != nil {
		t.Fatal(err)
	}

	// Try to get the deleted user by ID.
	var retrievedDeletedUser User
	err = db.Get(&retrievedDeletedUser, `SELECT * FROM users WHERE id=$1`, newUser.ID)
	if err == nil {
		t.Errorf("Expected error, but got user %+v", retrievedDeletedUser)
	}
}
