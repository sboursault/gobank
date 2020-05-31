package store

import (
	es "github.com/sboursault/gobank/eventsourcing"
)

type inMemoryStore struct {
	events []es.Event
}

// consctructor

func NewInMemory() es.EventStore {
	return &inMemoryStore{}
}

// public functions

func (store *inMemoryStore) Write(event es.Event) {
	store.events = append(store.events, event)
}

func (store *inMemoryStore) ReadStream(streamId string) (stream es.Stream) {
	array := filter(
		store.events,
		func(event es.Event) bool { return event.StreamId == streamId })
	return es.NewStream(array...)
}

// private functions

func (store *inMemoryStore) clear() {
	store.events = nil
}

func filter(elements []es.Event, f func(es.Event) bool) []es.Event {

	result := []es.Event{}
	for _, event := range elements {
		if f(event) {
			result = append(result, event)
		}
	}
	return result
}
