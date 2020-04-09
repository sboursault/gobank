package eventstore

import es "github.com/sboursault/gobank/eventsourcing"

type InMemory struct {
	events []es.Event
}

func (store *InMemory) Write(event es.Event) {
	store.events = append(store.events, event)
}

func (store *InMemory) Read(aggregate string, streamId string) (events []es.Event) {
	return filter(
		store.events,
		func(event es.Event) bool { return event.StreamId == streamId })
}

func (store *InMemory) Clear() {
	store.events = nil
}

func filter(events []es.Event, predicate func(es.Event) bool) []es.Event {
	result := []es.Event{}
	for _, event := range events {
		if predicate(event) {
			result = append(result, event)
		}
	}
	return result
}

/*
NewInMemory creates an InMemory structure, with a nil Event slice.
It returns a pointer to the created structure.
*/
func NewInMemory() *InMemory {
	return &InMemory{}
}
