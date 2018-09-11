package main

import (
	"github.com/jmoiron/sqlx"
)

func main() {

}

func DoubleCoin(db *sqlx.DB, userID int64) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	q1 := "SELECT coin FROM user_coin WHERE user_id = ? FOR UPDATE"
	var coin int64
	if err := tx.Get(&coin, q1, userID); err != nil {
		tx.Rollback()
		return err
	}
	q2 := "UPDATE user_coin SET coin = ? WHERE user_id = ?"
	_, err = tx.Exec(q2, coin*2, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func CreateOrUpdateUser(db *sqlx.DB, name, serviceToken string) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	q1 := "INSERT INTO user (name) VALUES (?)"
	result, err := tx.Exec(q1, name)
	if err != nil {
		tx.Rollback()
		return err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	q2 := "INSERT INTO user_service (user_id, service_token) VALUES (?, ?) ON DUPLICATE KEY UPDATE user_id = ?"
	result2, err := tx.Exec(q2, userID, serviceToken, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	affected, err := result2.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affected == 2 {
		q3 := "UPDATE service_link SET user_id = ? WHERE service_token = ?"
		_, err := tx.Exec(q3, userID, serviceToken)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
