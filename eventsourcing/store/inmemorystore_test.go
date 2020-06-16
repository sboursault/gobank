package store

import (
	"reflect"
	"testing"

	"time"

	es "github.com/sboursault/gobank/eventsourcing"
)

var store = inMemoryStore{}

func Test_InMemory_storesEvents(t *testing.T) {

	store.clear()
	store.Write(es.Event{AggregateType: "account", StreamId: "account:001", Date: date(2020, time.February, 10), EventType: "opened", Payload: `{"owner":"Snow John"}`})
	store.Write(es.Event{AggregateType: "account", StreamId: "account:001", Date: date(2020, time.February, 15), EventType: "deposited", Payload: "{\"amount\":1000}"})
	store.Write(es.Event{AggregateType: "account", StreamId: "account:002", Date: date(2020, time.February, 25), EventType: "opened", Payload: "{\"owner\":\"Arya Stark\"}"})

	got := store.ReadStream("account:001")

	want := es.NewStream(
		es.Event{AggregateType: "account", StreamId: "account:001", Date: date(2020, time.February, 10), EventType: "opened", Payload: `{"owner":"Snow John"}`},
		es.Event{AggregateType: "account", StreamId: "account:001", Date: date(2020, time.February, 15), EventType: "deposited", Payload: "{\"amount\":1000}"})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

// helpers
func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}
