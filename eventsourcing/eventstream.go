package eventsourcing

type EventStream struct {
	Events []Event
}

/*
rehydrate / replay / left fold
*/
//func (stream *Stream) Replay(handlers map[string]Handler) Aggregate {

func (stream *EventStream) LeftFold(
	init Aggregate,
	handlers map[string]func(Aggregate, Event) Aggregate) Aggregate {

	aggregate := init

	for _, event := range stream.Events {
		f := handlers[event.EventType]
		aggregate = f(aggregate, event)
	}
	return aggregate
}

func NewEventStream(events ...Event) EventStream {
	return EventStream{events}
}
