package bank

import (
	"fmt"

	"github.com/lithammer/shortuuid"

	"encoding/json"

	"github.com/sboursault/gobank/bank/accounts"
	es "github.com/sboursault/gobank/eventsourcing"

	"github.com/sboursault/gobank/eventsourcing/store"
)

// types

type Stream = es.Stream
type Event = es.Event
type Aggregate = es.Aggregate
type EventStore = es.EventStore

var eventStore = store.PgConnection()

func openAccount(owner string) string {
	accountId := shortuuid.New()

	event, _ := json.Marshal(accounts.NewOpenedEvent(owner))

	eventStore.Write(es.NewEvent("account", accountId, "opened", string(event)))

	return accountId
}

func deposit(accountId string, amount float32) {

	event, _ := json.Marshal(accounts.NewDepositedEvent(amount))
	eventStore.Write(es.NewEvent("account", accountId, "deposited", string(event)))
}

func withdraw(accountId string, amount float32) error {

	aggregate := accounts.Get(eventStore, accountId)

	if amount > aggregate.Balance {
		return fmt.Errorf("Not enough money to withdraw %g (account balance: %g)", amount, aggregate.Balance)
	}

	event, _ := json.Marshal(accounts.NewWithdrawnEvent(amount))
	eventStore.Write(es.NewEvent("account", accountId, "withdrawn", string(event)))

	return nil
}

func closeAccount(accountId string) error {

	aggregate := accounts.Get(eventStore, accountId)

	if aggregate.Balance != 0 {
		return fmt.Errorf("Can't close account (account balance: %g)", aggregate.Balance)
	}

	event, _ := json.Marshal(accounts.NewClosedEvent())
	eventStore.Write(es.NewEvent("account", accountId, "closed", string(event)))

	return nil
}
