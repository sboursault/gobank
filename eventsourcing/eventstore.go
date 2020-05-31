package eventsourcing

type EventStore interface {
	Write(event Event)
	Read(streamId string) (stream Stream)
}
