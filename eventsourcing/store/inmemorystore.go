package store

import (
	es "github.com/sboursault/gobank/eventsourcing"
)

type inMemoryStore struct {
	events []es.Event
}

func (store *inMemoryStore) Write(event es.Event) {
	store.events = append(store.events, event)
}

func (store *inMemoryStore) Read(streamId string) (stream es.Stream) {
	array := filter(
		store.events,
		func(event es.Event) bool { return event.StreamId == streamId })
	return es.NewStream(array...)

}

func (store *inMemoryStore) Clear() {
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

/*
NewInMemory creates an InMemory structure, with a nil Event slice.
It returns a pointer to the created structure.
*/
func NewInMemory() *inMemoryStore {
	return &inMemoryStore{}
}
