package store

import "github.com/sboursault/gobank/events"

type inMemoryStore struct {
	events []events.Event
}

func (store *inMemoryStore) Write(event events.Event) {
	store.events = append(store.events, event)
}

func (store *inMemoryStore) Read(streamId string) (stream events.Stream) {
	array := filter(
		store.events,
		func(event events.Event) bool { return event.StreamId == streamId })
	return events.NewStream(array...)

}

func (store *inMemoryStore) Clear() {
	store.events = nil
}

func filter(elements []events.Event, f func(events.Event) bool) []events.Event {

	result := []events.Event{}
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
