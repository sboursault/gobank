package eventsourcing

type Event struct {
	AggregateType string
	StreamId      string
	EventType     string
	Payload       string
}

func NewEvent(
	aggregateType string,
	streamId string,
	eventType string,
	payload string) Event {
	return Event{aggregateType, streamId, eventType, payload}
}
