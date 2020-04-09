package eventsourcing

type InMemoryStore struct {
	events []Event
}

func (store *InMemoryStore) Write(event Event) {
	store.events = append(store.events, event)
}

func (store *InMemoryStore) Read(streamId string) (stream EventStream) {
	events := filter(
		store.events,
		func(event Event) bool { return event.StreamId == streamId })
	return EventStream{events}
}

func (store *InMemoryStore) Clear() {
	store.events = nil
}

func filter(events []Event, predicate func(Event) bool) []Event {
	result := []Event{}
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
func NewInMemory() *InMemoryStore {
	return &InMemoryStore{}
}
