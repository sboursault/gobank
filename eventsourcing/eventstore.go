package eventsourcing

type EventStore interface {
	Read(streamId string) (events []Event)
}
