package eventsourcing

type EventStream struct {
	events []Event
}

/*
rehydrate / replay / left fold
*/
//func (stream *Stream) Replay(handlers map[string]Handler) Aggregate {

func (stream *EventStream) LeftFold(
	init Aggregate,
	handlers map[string]func(Aggregate, Event) Aggregate) Aggregate {

	aggregate := init

	for _, event := range stream.events {
		f := handlers[event.EventType]
		aggregate = f(aggregate, event)
	}
	return aggregate
}
