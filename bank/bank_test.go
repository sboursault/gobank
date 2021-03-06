package bank

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/sboursault/gobank/bank/accounts"
	"github.com/sboursault/gobank/eventsourcing/store"
)

/*
TestMain runs in the main goroutine and can do whatever setup and teardown is necessary around a call to m.Run.
It should then call os.Exit with the result of m.Run.
*/
func TestMain(m *testing.M) {

	// setup
	eventStore = store.NewInMemory()
	// run test
	code := m.Run()
	// teardown
	// ...
	os.Exit(code)
}

func Test_OpenAccount(t *testing.T) {

	id := OpenAccount("John Snow")

	got := accounts.Get(eventStore, id)

	want := accounts.New()
	want.Owner = "John Snow"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_Deposit(t *testing.T) {

	accountNumber := OpenAccount("John Snow")

	Deposit(accountNumber, 200)

	got := accounts.Get(eventStore, accountNumber)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 200

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_withraw(t *testing.T) {

	accountNumber := OpenAccount("John Snow")

	Deposit(accountNumber, 200)
	Withdraw(accountNumber, 50)

	got := accounts.Get(eventStore, accountNumber)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 150

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_withraw_refused(t *testing.T) {

	accountNumber := OpenAccount("John Snow")

	gotErr := Withdraw(accountNumber, 50)

	wantedErr := errors.New("Not enough money to withdraw 50 (account balance: 0)")

	if !reflect.DeepEqual(gotErr, wantedErr) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", wantedErr, gotErr)
		return
	}

	got := accounts.Get(eventStore, accountNumber)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 0

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_close(t *testing.T) {

	accountNumber := OpenAccount("John Snow")

	CloseAccount(accountNumber)

	got := accounts.Get(eventStore, accountNumber)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 0
	want.Closed = true

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_close_refused(t *testing.T) {

	accountNumber := OpenAccount("John Snow")
	Deposit(accountNumber, 200)

	gotErr := CloseAccount(accountNumber)

	wantedErr := errors.New("Can't close account (account balance: 200)")

	if !reflect.DeepEqual(gotErr, wantedErr) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", wantedErr, gotErr)
		return
	}

	got := accounts.Get(eventStore, accountNumber)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 200
	want.Closed = false

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}
