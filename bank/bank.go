package bank

import (
	"fmt"
	"math/rand"
	"time"

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

// public functions

/*
init function is called at startup
*/
func init() {
	rand.Seed(time.Now().UnixNano()) // initialize rand module
}

func OpenAccount(owner string) string {
	accountNumber := generateAccountNumber()

	event, _ := json.Marshal(accounts.NewOpenedEvent(owner))

	eventStore.Write(es.NewEvent("account", accountNumber, "opened", string(event)))

	return accountNumber
}

func Deposit(accountNumber string, amount float32) {

	event, _ := json.Marshal(accounts.NewDepositedEvent(amount))
	eventStore.Write(es.NewEvent("account", accountNumber, "deposited", string(event)))
}

func Withdraw(accountNumber string, amount float32) error {

	aggregate := accounts.Get(eventStore, accountNumber)

	if amount > aggregate.Balance {
		return fmt.Errorf("Not enough money to withdraw %g (account balance: %g)", amount, aggregate.Balance)
	}

	event, _ := json.Marshal(accounts.NewWithdrawnEvent(amount))
	eventStore.Write(es.NewEvent("account", accountNumber, "withdrawn", string(event)))

	return nil
}

func CloseAccount(accountNumber string) error {

	aggregate := accounts.Get(eventStore, accountNumber)

	if aggregate.Balance != 0 {
		return fmt.Errorf("Can't close account (account balance: %g)", aggregate.Balance)
	}

	event, _ := json.Marshal(accounts.NewClosedEvent())
	eventStore.Write(es.NewEvent("account", accountNumber, "closed", string(event)))

	return nil
}

func GetAccountInfo(accountNumber string) string {

	aggregate := accounts.Get(eventStore, accountNumber)

	info := "Account " + accountNumber + "\n"
	info += "Balance: " + fmt.Sprintf("%f", aggregate.Balance) + "\n"

	if aggregate.Closed {
		info += "Closed\n"
	}

	return info
}

// private functions

func generateAccountNumber() string {
	random_number := rand.Int63n(1e11)         // generate number
	return fmt.Sprintf("%011d", random_number) // left pad with 0
}
