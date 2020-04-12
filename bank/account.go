package main

import (
	marshaller "encoding/json"
	"fmt"

	es "github.com/sboursault/gobank/eventsourcing"
)

type Account struct {
	owner   string
	balance float32
}

// account events

const debug = true

type OpenedEvent struct {
	Owner string `json:"owner"`
}

type DepositedEvent struct {
	Amount float32 `json:"amount"`
}

type WithdrawnEvent struct {
	Amount float32 `json:"amount"`
}

func leftFold(account Account, stream es.EventStream) Account {

	handlers := map[string]func(es.Aggregate, es.Event) es.Aggregate{
		"opened":    onOpenedEvent,
		"deposited": onDepositedEvent,
		"withdrawn": onWithdrawnEvent}

	return stream.LeftFold(account, handlers).(Account)
}

// event handlers

func onOpenedEvent(aggregate es.Aggregate, event es.Event) es.Aggregate {
	account := aggregate.(Account)

	log("event", event)

	payload := unmarshalOpenedEvent(event.Payload)

	account.owner = payload.Owner

	log("account", account)

	return account
}

func onDepositedEvent(aggregate es.Aggregate, event es.Event) es.Aggregate {
	account := aggregate.(Account)

	log("event", event)

	payload := unmarshalDepositedEvent(event.Payload)

	account.balance += payload.Amount

	log("account", account)

	return account
}

func onWithdrawnEvent(aggregate es.Aggregate, event es.Event) es.Aggregate {
	account := aggregate.(Account)

	log("event", event)

	payload := unmarshalWithdrawnEvent(event.Payload)

	account.balance -= payload.Amount

	log("account", account)

	return account
}

// utils

func unmarshalOpenedEvent(json string) OpenedEvent {
	target := OpenedEvent{}
	marshaller.Unmarshal([]byte(json), &target)
	return target
}

func unmarshalDepositedEvent(json string) DepositedEvent {
	target := DepositedEvent{}
	marshaller.Unmarshal([]byte(json), &target)
	return target
}

func unmarshalWithdrawnEvent(json string) WithdrawnEvent {
	target := WithdrawnEvent{}
	marshaller.Unmarshal([]byte(json), &target)
	return target
}

func log(label string, something interface{}) {
	if debug {
		fmt.Printf(label+": %+v\n", something)
	}
}
