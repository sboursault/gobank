package accounts

import (
	"reflect"
	"testing"

	es "github.com/sboursault/gobank/eventsourcing"
)

func Test_OpenedEvent(t *testing.T) {

	openedEvent := es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`)

	got := onOpenedEvent(account{}, openedEvent)

	want := account{Owner: "Snow John", Balance: 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_DepositedEvent(t *testing.T) {

	openedEvent := es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`)
	depositedEvent := es.NewEvent("account", "account:00001", "deposited", `{"amount":100}`)

	aggregate := onOpenedEvent(account{}, openedEvent)
	got := onDepositedEvent(aggregate, depositedEvent)

	want := account{Owner: "Snow John", Balance: 100}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_WithdrawnEvent(t *testing.T) {

	openedEvent := es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`)
	depositedEvent := es.NewEvent("account", "account:00001", "deposited", `{"amount":100}`)
	withdrawnEvent := es.NewEvent("account", "account:00001", "withdrawn", `{"amount":30}`)

	aggregate := onOpenedEvent(account{}, openedEvent)
	aggregate = onDepositedEvent(aggregate, depositedEvent)
	got := onWithdrawnEvent(aggregate, withdrawnEvent)

	want := account{Owner: "Snow John", Balance: 70}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_ClosedEvent(t *testing.T) {

	openedEvent := es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`)
	closedEvent := es.NewEvent("account", "account:00001", "closed", `{}`)

	aggregate := onOpenedEvent(account{}, openedEvent)
	got := onClosedEvent(aggregate, closedEvent)

	want := account{Owner: "Snow John", Closed: true}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_LeftFold(t *testing.T) {

	stream := es.NewStream(
		es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`),
		es.NewEvent("account", "account:00001", "deposited", `{"amount":100}`),
		es.NewEvent("account", "account:00001", "withdrawn", `{"amount":30}`),
		es.NewEvent("account", "account:00001", "closed", `{}`))

	got := LeftFold(stream)

	want := account{Owner: "Snow John", Balance: 70, Closed: true}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}
