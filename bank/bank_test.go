package main // TODO rename bank

import (
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
