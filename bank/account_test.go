package main // TODO rename bank

import (
	"reflect"
	"testing"

	es "github.com/sboursault/gobank/eventsourcing"
)

func Test_OpenedEvent(t *testing.T) {

	openedEvent := es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`)

	got := onOpenedEvent(Account{}, openedEvent)

	want := Account{"Snow John", 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_DepositedEvent(t *testing.T) {

	openedEvent := es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`)
	depositedEvent := es.NewEvent("account", "account:00001", "deposited", `{"amount":100}`)

	aggregate := onOpenedEvent(Account{}, openedEvent)
	got := onDepositedEvent(aggregate, depositedEvent)

	want := Account{"Snow John", 100}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_WithdrawnEvent(t *testing.T) {

	openedEvent := es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`)
	depositedEvent := es.NewEvent("account", "account:00001", "deposited", `{"amount":100}`)
	withdrawnEvent := es.NewEvent("account", "account:00001", "withdrawn", `{"amount":30}`)

	aggregate := onOpenedEvent(Account{}, openedEvent)
	aggregate = onDepositedEvent(aggregate, depositedEvent)
	got := onWithdrawnEvent(aggregate, withdrawnEvent)

	want := Account{"Snow John", 70}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_LeftFold(t *testing.T) {

	stream := es.NewEventStream(
		es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`),
		es.NewEvent("account", "account:00001", "deposited", `{"amount":100}`),
		es.NewEvent("account", "account:00001", "withdrawn", `{"amount":30}`))

	got := leftFold(Account{}, stream)

	want := Account{"Snow John", 70}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

// TODO
// verif balance positive
// blagues : John Snow ouvre un compte en banque, verif balance un peu bête car empeche les agio
// brancher sur base PG
// brancher sur ligne de commande (après BD, ainsi il y aura un stockage persistant)
