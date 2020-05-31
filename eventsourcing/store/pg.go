package store

/*
source https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.4.html
*/

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // pg driver, needs to be imported even if not explicitly used
	es "github.com/sboursault/gobank/eventsourcing"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "changeme"
	DB_NAME     = "postgres"
)

// types

type Stream = es.Stream
type Event = es.Event
type Aggregate = es.Aggregate

type pgStore struct {
}

// consctructor

func NewPg() es.EventStore {
	return &pgStore{}
}

// public functions

func (store *pgStore) Write(event es.Event) {
	db := connect()

	var lastInsertId int
	err := db.QueryRow(`
		INSERT INTO gobank.t_event(aggregate_type, stream_id, event_type, payload)
		VALUES($1, $2, $3, $4)
		returning id;
		`, event.AggregateType, event.StreamId, event.EventType, event.Payload).Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)

	defer db.Close()
}

func (store *pgStore) ReadStream(streamId string) (stream Stream) {
	db := connect()
	rows, err := db.Query("SELECT * FROM gobank.t_event")
	checkErr(err)

	var events []Event

	for rows.Next() {
		var aggregateType string
		var streamId string
		var eventType string
		var payload string
		err = rows.Scan(&aggregateType, &streamId, &eventType, &payload)
		checkErr(err)

		events = append(events, es.NewEvent(aggregateType, streamId, eventType, payload))

	}
	return es.NewStream(events...)
}

// private functions

func connect() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
