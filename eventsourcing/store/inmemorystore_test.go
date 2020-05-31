package store

import (
	"reflect"
	"testing"

	es "github.com/sboursault/gobank/eventsourcing"
)

var store = inMemoryStore{}

func Test_InMemory_storesEvents(t *testing.T) {

	store.clear()
	store.Write(es.NewEvent("account", "account:001", "opened", `{"owner":"Snow John"}`))
	store.Write(es.NewEvent("account", "account:001", "deposited", "{\"amount\":1000}"))
	store.Write(es.NewEvent("account", "account:002", "opened", "{\"owner\":\"Arya Stark\"}"))

	got := store.ReadStream("account:001")

	want := es.NewStream(
		es.NewEvent("account", "account:001", "opened", `{"owner":"Snow John"}`),
		es.NewEvent("account", "account:001", "deposited", "{\"amount\":1000}"))

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}
