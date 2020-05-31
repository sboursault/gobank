package bank

import (
	"errors"
	"reflect"
	"testing"

	"github.com/sboursault/gobank/bank/accounts"
)

func Test_openAccount(t *testing.T) {

	id := openAccount("John Snow")

	got := accounts.Get(eventStore, id)

	want := accounts.New()
	want.Owner = "John Snow"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_deposit(t *testing.T) {

	accountId := openAccount("John Snow")

	deposit(accountId, 200)

	got := accounts.Get(eventStore, accountId)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 200

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_withraw(t *testing.T) {

	accountId := openAccount("John Snow")

	deposit(accountId, 200)
	withdraw(accountId, 50)

	got := accounts.Get(eventStore, accountId)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 150

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_withraw_refused(t *testing.T) {

	accountId := openAccount("John Snow")

	gotErr := withdraw(accountId, 50)

	wantedErr := errors.New("Not enough money to withdraw 50 (account balance: 0)")

	if !reflect.DeepEqual(gotErr, wantedErr) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", wantedErr, gotErr)
		return
	}

	got := accounts.Get(eventStore, accountId)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 0

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_close(t *testing.T) {

	accountId := openAccount("John Snow")

	closeAccount(accountId)

	got := accounts.Get(eventStore, accountId)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 0
	want.Closed = true

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_close_refused(t *testing.T) {

	accountId := openAccount("John Snow")
	deposit(accountId, 200)

	gotErr := closeAccount(accountId)

	wantedErr := errors.New("Can't close account (account balance: 200)")

	if !reflect.DeepEqual(gotErr, wantedErr) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", wantedErr, gotErr)
		return
	}

	got := accounts.Get(eventStore, accountId)

	want := accounts.New()
	want.Owner = "John Snow"
	want.Balance = 200
	want.Closed = false

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}
