package main

import (
	"database/sql"
	"fmt"
)

func main() {

}

func ConsumeCoin(db *sql.DB, userID int64, amount int64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	q1 := "SELECT coin FROM user_coin WHER user_id = ? FOR UPDATE"
	var coin int64
	row := db.QueryRow(q1, userID)
	if err := row.Scan(&coin); err != nil {
		tx.Rollback()
		return err
	}
	if coin-amount > 0 {
		q2 := "UPDATE user_coin SET coin - ? WHERE user_id = ?"
		result, err := db.Exec(q2, amount, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
		affected, _ := result.RowsAffected()
		if affected == 0 {
			tx.Rollback()
			return fmt.Errorf("no rows were affected")
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
