package db

import (
	"database/sql"

	"barr.io/graph/model" // module's generated models
	_ "github.com/lib/pq" // import Postgres
)

var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

func SaveMessage(message *model.Message) error {
	query := "INSERT INTO messages (text) VALUES ($1) RETURNING id"
	return db.QueryRow(query, message.Text).Scan(&message.ID)
}
