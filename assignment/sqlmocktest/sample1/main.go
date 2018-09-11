package main

import (
	"database/sql"
	"log"
)

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root@/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	user, err := FetchUser(db, 100)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", user)
}

func FetchUser(db *sql.DB, id int64) (*User, error) {
	q := " SELECT id, name FROM user WHERE id = ? "
	row := db.QueryRow(q, id)
	u := User{}
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}
