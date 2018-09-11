package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

}

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func sqlxFetchUser(db *sqlx.DB, id int64) (*User, error) {
	q := " SELECT id, name FROM user WHERE id = ? "
	var user User
	if err := db.Get(&user, q, id); err != nil {
		return nil, err
	}
	return &user, nil
}
