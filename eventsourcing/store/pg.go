package store

/*
source https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.4.html
*/

import (
	"database/sql"
	"fmt"

	pg "github.com/lib/pq" // pg driver, needs to be imported even if not explicitly used
	es "github.com/sboursault/gobank/eventsourcing"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "changeme"
	DB_NAME     = "postgres"

	debug = false
)

// types

type Stream = es.Stream
type Event = es.Event
type Aggregate = es.Aggregate

type pgStore struct {
}

// consctructor

func PgConnection() es.EventStore {
	return &pgStore{}
}

// public functions

func (store *pgStore) Write(event es.Event) {
	db := connect()
	defer db.Close() // will be executed at the end of the surrounding function

	var lastInsertId int
	err := db.QueryRow(`
		INSERT INTO gobank.t_event(aggregate_type, stream_id, date, event_type, payload)
		VALUES($1, $2, $3, $4, $5)
		returning id;
		`, event.AggregateType, event.StreamId, event.Date, event.EventType, event.Payload).Scan(&lastInsertId)
	checkErr(err)
	log("last inserted id", lastInsertId)

}

func (store *pgStore) ReadStream(streamId string) (stream Stream) {
	db := connect()
	defer db.Close() // will be executed at the end of the surrounding function

	rows, err := db.Query(`
		SELECT aggregate_type, stream_id, date, event_type, payload
		FROM gobank.t_event
		WHERE stream_id = $1`, streamId)
	checkErr(err)

	var events []Event

	for rows.Next() {
		var aggregateType string
		var streamId string
		var date pg.NullTime
		var eventType string
		var payload string
		err = rows.Scan(&aggregateType, &streamId, &date, &eventType, &payload)
		checkErr(err)

		events = append(events, es.Event{
			AggregateType: aggregateType,
			StreamId:      streamId,
			Date:          date.Time,
			EventType:     eventType,
			Payload:       payload})

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

func log(label string, something interface{}) {
	if debug {
		fmt.Printf(label+": %+v\n", something)
	}
}
