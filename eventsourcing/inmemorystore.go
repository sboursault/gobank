package eventsourcing

type inMemoryStore struct {
	events []Event
}

func (store *inMemoryStore) Write(event Event) {
	store.events = append(store.events, event)
}

func (store *inMemoryStore) Read(streamId string) (stream EventStream) {
	events := filter(
		store.events,
		func(event Event) bool { return event.StreamId == streamId })
	return EventStream{events}
}

func (store *inMemoryStore) Clear() {
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
func NewInMemory() *inMemoryStore {
	return &inMemoryStore{}
}
