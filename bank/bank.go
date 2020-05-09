package bank

import (
	"fmt"

	"github.com/lithammer/shortuuid"

	"encoding/json"

	"github.com/sboursault/gobank/bank/account"
	"github.com/sboursault/gobank/events"
	"github.com/sboursault/gobank/events/store"
)

var eventStore = store.NewInMemory()

func openAccount(owner string) string {
	accountId := shortuuid.New()

	event, _ := json.Marshal(account.NewOpenedEvent(owner))

	eventStore.Write(events.New("account", accountId, "opened", string(event))) // un peu moche events.New, events.New ? masquer Event ?

	// see https://golangbot.com/go-packages/

	return accountId
}

func deposit(accountId string, amount float32) {

	event, _ := json.Marshal(account.NewDepositedEvent(amount))
	eventStore.Write(events.New("account", accountId, "deposited", string(event)))
}

func withdraw(accountId string, amount float32) error {

	aggregate := getAccount(accountId)

	if amount > aggregate.Balance {
		return fmt.Errorf("Not enough money to withdraw %g (account balance: %g)", amount, aggregate.Balance)
	}

	event, _ := json.Marshal(account.NewWithdrawnEvent(amount))
	eventStore.Write(events.New("account", accountId, "withdrawn", string(event)))

	return nil
}

func closeAccount(accountId string) error {

	aggregate := getAccount(accountId)

	if aggregate.Balance != 0 {
		return fmt.Errorf("Can't close account (account balance: %g)", aggregate.Balance)
	}

	event, _ := json.Marshal(account.NewClosedEvent())
	eventStore.Write(events.New("account", accountId, "closed", string(event)))

	return nil
}

func getAccount(id string) account.Account {

	stream := eventStore.Read(id)

	return account.LeftFold(stream)
}

func PrintAccountInfo(id string) {

}

// next step : tester openAccount
