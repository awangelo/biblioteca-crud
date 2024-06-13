package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDatabase() error {
	connection := "user=root password=root dbname=databasezuda host=localhost sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connectado!")

	return nil
}
