package main

import (
	"fmt"

	"github.com/lithammer/shortuuid"

	"encoding/json"

	es "github.com/sboursault/gobank/eventsourcing"
)

var eventStore = es.NewInMemory()

/*

interface to open / deposit / withdrow / close an account



eventStore := es.InMemoryStore{}

	eventStore.Write(es.NewEvent("account", "account:00001", "opened", `{"owner":"Snow John"}`))
	eventStore.Write(es.NewEvent("account", "account:00001", "deposited", `{"amount":100}`))
	eventStore.Write(es.NewEvent("account", "account:00001", "withdrawn", `{"amount":30}`))

	stream := eventStore.Read("account:00001")

	got := leftFold(Account{}, stream)

	want := Account{"Snow John", 70}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got)
	}


*/

func openAccount(owner string) string {
	accountId := shortuuid.New()

	event, _ := json.Marshal(OpenedEvent{owner}) // tester sortir account.go dans un bank/account/account.go -> la syntaxe est-elle plus sympa ?

	eventStore.Write(es.NewEvent("account", accountId, "opened", string(event))) // un peu moche es.NewEvent, events.New ? masquer Event ?

	// see https://golangbot.com/go-packages/

	return accountId
}

func deposit(accountId string, amount float32) {

	event, _ := json.Marshal(DepositedEvent{amount})
	eventStore.Write(es.NewEvent("account", accountId, "deposited", string(event)))
}

func withdraw(accountId string, amount float32) error {

	account := getAccount(accountId)

	if amount > account.balance {
		return fmt.Errorf("Not enough money to withdraw %g (account balance: %g)", amount, account.balance)
	}

	event, _ := json.Marshal(WithdrawnEvent{amount})
	eventStore.Write(es.NewEvent("account", accountId, "withdrawn", string(event)))

	return nil
}

func closeAccount(accountId string) error {

	account := getAccount(accountId)

	if account.balance != 0 {
		return fmt.Errorf("Can't close account (account balance: %g)", account.balance)
	}

	event, _ := json.Marshal(ClosedEvent{})
	eventStore.Write(es.NewEvent("account", accountId, "closed", string(event)))

	return nil
}

func getAccount(id string) Account {

	stream := eventStore.Read(id)

	return leftFold(stream)

}

func PrintAccountInfo(id string) {

}

// next step : tester openAccount
