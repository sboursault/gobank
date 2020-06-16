package eventsourcing

import (
	"time"
)

type Event struct {
	AggregateType string
	StreamId      string
	Date          time.Time
	EventType     string
	Payload       string
}

func NewEvent(
	aggregateType string,
	streamId string,
	eventType string,
	payload string) Event {
	return Event{aggregateType, streamId, time.Now(), eventType, payload}
}
