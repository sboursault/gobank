package main // TODO rename bank

import (
	"errors"
	"reflect"
	"testing"
)

func Test_openAccount(t *testing.T) {

	id := openAccount("John Snow")

	got := getAccount(id)

	want := Account{owner: "John Snow", balance: 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_deposit(t *testing.T) {

	accountId := openAccount("John Snow")

	deposit(accountId, 200)

	got := getAccount(accountId)

	want := Account{owner: "John Snow", balance: 200}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_withraw(t *testing.T) {

	accountId := openAccount("John Snow")

	deposit(accountId, 200)
	withdraw(accountId, 50)

	got := getAccount(accountId)

	want := Account{owner: "John Snow", balance: 150}

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

	got := getAccount(accountId)

	want := Account{owner: "John Snow", balance: 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}

func Test_close(t *testing.T) {

	accountId := openAccount("John Snow")

	closeAccount(accountId)

	got := getAccount(accountId)

	want := Account{owner: "John Snow", balance: 0, closed: true}

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

	got := getAccount(accountId)

	want := Account{owner: "John Snow", balance: 200, closed: false}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}
}
