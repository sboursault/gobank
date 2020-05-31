package eventsourcing

type EventStore interface {
	Write(event Event)
	ReadStream(streamId string) (stream Stream)
}
