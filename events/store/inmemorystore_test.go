package store

import (
	"reflect"
	"testing"

	events "github.com/sboursault/gobank/events"
)

var store = NewInMemory()

func Test_InMemory_storesEvents(t *testing.T) {

	store.Clear()
	store.Write(events.New("account", "account:001", "opened", `{"owner":"Snow John"}`))
	store.Write(events.New("account", "account:001", "deposited", "{\"amount\":1000}"))
	store.Write(events.New("account", "account:002", "opened", "{\"owner\":\"Arya Stark\"}"))

	got := store.Read("account:001")

	want := events.NewStream(
		events.New("account", "account:001", "opened", `{"owner":"Snow John"}`),
		events.New("account", "account:001", "deposited", "{\"amount\":1000}"))

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}
