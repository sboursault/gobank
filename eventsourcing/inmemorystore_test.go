package eventsourcing

import (
	"reflect"
	"testing"
)

var store = NewInMemory()

func Test_InMemory_storesEvents(t *testing.T) {

	store.Clear()
	store.Write(NewEvent("account", "account:001", "opened", `{"owner":"Snow John"}`))
	store.Write(NewEvent("account", "account:001", "deposited", "{\"amount\":1000}"))
	store.Write(NewEvent("account", "account:002", "opened", "{\"owner\":\"Arya Stark\"}"))

	got := store.Read("account:001")

	want := EventStream{
		[]Event{
			NewEvent("account", "account:001", "opened", `{"owner":"Snow John"}`),
			NewEvent("account", "account:001", "deposited", "{\"amount\":1000}")}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}
