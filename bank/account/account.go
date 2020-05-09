package account

import (
	marshaller "encoding/json"
	"fmt"

	es "github.com/sboursault/gobank/eventsourcing"
)

const debug = true

// types

type Account struct {
	Owner   string
	Balance float32
	Closed  bool
}

type openedEvent struct {
	Owner string `json:"owner"`
}

type depositedEvent struct {
	Amount float32 `json:"amount"`
}

type withdrawnEvent struct {
	Amount float32 `json:"amount"`
}

type closedEvent struct {
}

// functions

func LeftFold(stream es.EventStream) Account {

	handlers := map[string]func(es.Aggregate, es.Event) es.Aggregate{
		"opened":    onOpenedEvent,
		"deposited": onDepositedEvent,
		"withdrawn": onWithdrawnEvent,
		"closed":    onClosedEvent}

	return stream.LeftFold(Account{}, handlers).(Account)
}

// creators

func NewOpenedEvent(owner string) openedEvent {
	return openedEvent{owner}
}

func NewDepositedEvent(amount float32) depositedEvent {
	return depositedEvent{amount}
}

func NewWithdrawnEvent(amount float32) withdrawnEvent {
	return withdrawnEvent{amount}
}

func NewClosedEvent() closedEvent {
	return closedEvent{}
}

// event handlers

func onOpenedEvent(aggregate es.Aggregate, event es.Event) es.Aggregate {
	account := aggregate.(Account)

	log("event", event)

	payload := unmarshalOpenedEvent(event.Payload)

	account.Owner = payload.Owner

	log("account", account)

	return account
}

func onDepositedEvent(aggregate es.Aggregate, event es.Event) es.Aggregate {
	account := aggregate.(Account)

	log("event", event)

	payload := unmarshalDepositedEvent(event.Payload)

	account.Balance += payload.Amount

	log("account", account)

	return account
}

func onWithdrawnEvent(aggregate es.Aggregate, event es.Event) es.Aggregate {
	account := aggregate.(Account)

	log("event", event)

	payload := unmarshalWithdrawnEvent(event.Payload)

	account.Balance -= payload.Amount

	log("account", account)

	return account
}

func onClosedEvent(aggregate es.Aggregate, event es.Event) es.Aggregate {
	account := aggregate.(Account)

	log("event", event)

	account.Closed = true

	log("account", account)

	return account
}

// utils

func unmarshalOpenedEvent(json string) openedEvent {
	target := openedEvent{}
	marshaller.Unmarshal([]byte(json), &target)
	return target
}

func unmarshalDepositedEvent(json string) depositedEvent {
	target := depositedEvent{}
	marshaller.Unmarshal([]byte(json), &target)
	return target
}

func unmarshalWithdrawnEvent(json string) withdrawnEvent {
	target := withdrawnEvent{}
	marshaller.Unmarshal([]byte(json), &target)
	return target
}

func log(label string, something interface{}) {
	if debug {
		fmt.Printf(label+": %+v\n", something)
	}
}
