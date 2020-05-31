package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // pg driver, needs to be imported even if not explicitly used
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "changeme"
	DB_NAME     = "postgres"
)

func connect() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	//
	return db
}

func insert() {
	fmt.Println("# Inserting values")

	db := connect()

	var lastInsertId int
	err := db.QueryRow(`
		INSERT INTO gobank.t_event(aggregate_type, stream_id, event_type, payload)
		VALUES($1, $2, $3, $4)
		returning id;
		`, "testgo", "testgo", "testgo", "testgo").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)

	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
