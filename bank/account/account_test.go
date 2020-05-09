package account

import (
	"reflect"
	"testing"

	"github.com/sboursault/gobank/events"
)

func Test_OpenedEvent(t *testing.T) {

	openedEvent := events.New("account", "account:00001", "opened", `{"owner":"Snow John"}`)

	got := onOpenedEvent(Account{}, openedEvent)

	want := Account{Owner: "Snow John", Balance: 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_DepositedEvent(t *testing.T) {

	openedEvent := events.New("account", "account:00001", "opened", `{"owner":"Snow John"}`)
	depositedEvent := events.New("account", "account:00001", "deposited", `{"amount":100}`)

	aggregate := onOpenedEvent(Account{}, openedEvent)
	got := onDepositedEvent(aggregate, depositedEvent)

	want := Account{Owner: "Snow John", Balance: 100}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_WithdrawnEvent(t *testing.T) {

	openedEvent := events.New("account", "account:00001", "opened", `{"owner":"Snow John"}`)
	depositedEvent := events.New("account", "account:00001", "deposited", `{"amount":100}`)
	withdrawnEvent := events.New("account", "account:00001", "withdrawn", `{"amount":30}`)

	aggregate := onOpenedEvent(Account{}, openedEvent)
	aggregate = onDepositedEvent(aggregate, depositedEvent)
	got := onWithdrawnEvent(aggregate, withdrawnEvent)

	want := Account{Owner: "Snow John", Balance: 70}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_ClosedEvent(t *testing.T) {

	openedEvent := events.New("account", "account:00001", "opened", `{"owner":"Snow John"}`)
	closedEvent := events.New("account", "account:00001", "closed", `{}`)

	aggregate := onOpenedEvent(Account{}, openedEvent)
	got := onClosedEvent(aggregate, closedEvent)

	want := Account{Owner: "Snow John", Closed: true}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_LeftFold(t *testing.T) {

	stream := events.NewStream(
		events.New("account", "account:00001", "opened", `{"owner":"Snow John"}`),
		events.New("account", "account:00001", "deposited", `{"amount":100}`),
		events.New("account", "account:00001", "withdrawn", `{"amount":30}`),
		events.New("account", "account:00001", "closed", `{}`))

	got := LeftFold(stream)

	want := Account{Owner: "Snow John", Balance: 70, Closed: true}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

// TODO
// how to switch memory store
//   soit c'est account qui masque la base
//   soit c'est un objet banque/service qui permet de récupérer un compte, faire un dépot
//      -> essayer cette option, dans le dernier test, ne plus faire apparaître l'event store
// verif balance positive
// blagues : John Snow ouvre un compte en banque, verif balance un peu bête car empeche les agio
// brancher sur base PG
// brancher sur ligne de commande (après BD, ainsi il y aura un stockage persistant)
