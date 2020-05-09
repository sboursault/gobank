package events

type Event struct {
	AggregateType string
	StreamId      string
	EventType     string
	Payload       string
}

func New(
	aggregateType string,
	streamId string,
	eventType string,
	payload string) Event {
	return Event{aggregateType, streamId, eventType, payload}
}
