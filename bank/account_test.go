package main // TODO rename bank

import (
	"reflect"
	"testing"

	es "github.com/sboursault/gobank/eventsourcing"
)

func Test_OpenedEvent(t *testing.T) {

	event := es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`)

	aggregate := Account{}

	got := onOpenedEvent(aggregate, event)

	want := Account{"Snow John", 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_DepositedEvent(t *testing.T) {

	event1 := es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`)
	event2 := es.NewEvent("account", "account:00001", "deposited", `{"amount":100}`)

	aggregate := Account{}

	got := onOpenedEvent(aggregate, event1)
	got = onDepositedEvent(got, event2)

	want := Account{"Snow John", 100}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}
