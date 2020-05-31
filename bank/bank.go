package bank

import (
	"fmt"

	"github.com/lithammer/shortuuid"

	"encoding/json"

	"github.com/sboursault/gobank/bank/accounts"
	"github.com/sboursault/gobank/events"
	"github.com/sboursault/gobank/events/store"
)

// types

type Stream = events.Stream
type Event = events.Event
type Aggregate = events.Aggregate
type EventStore = events.EventStore

var eventStore EventStore = store.NewInMemory()

func openAccount(owner string) string {
	accountId := shortuuid.New()

	event, _ := json.Marshal(accounts.NewOpenedEvent(owner))

	eventStore.Write(events.New("account", accountId, "opened", string(event)))

	return accountId
}

func deposit(accountId string, amount float32) {

	event, _ := json.Marshal(accounts.NewDepositedEvent(amount))
	eventStore.Write(events.New("account", accountId, "deposited", string(event)))
}

func withdraw(accountId string, amount float32) error {

	aggregate := accounts.Get(eventStore, accountId)

	if amount > aggregate.Balance {
		return fmt.Errorf("Not enough money to withdraw %g (account balance: %g)", amount, aggregate.Balance)
	}

	event, _ := json.Marshal(accounts.NewWithdrawnEvent(amount))
	eventStore.Write(events.New("account", accountId, "withdrawn", string(event)))

	return nil
}

func closeAccount(accountId string) error {

	aggregate := accounts.Get(eventStore, accountId)

	if aggregate.Balance != 0 {
		return fmt.Errorf("Can't close account (account balance: %g)", aggregate.Balance)
	}

	event, _ := json.Marshal(accounts.NewClosedEvent())
	eventStore.Write(events.New("account", accountId, "closed", string(event)))

	return nil
}

func PrintAccountInfo(id string) {

}
