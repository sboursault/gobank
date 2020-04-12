package eventsourcing

import (
	"strconv"
	"testing"
)

type counter struct {
	value int
}

func incr(a Aggregate, e Event) Aggregate {
	c := a.(counter)
	howmuch, _ := strconv.Atoi(e.Payload)
	c.value = c.value + howmuch
	return c
}

func decr(a Aggregate, e Event) Aggregate {
	c := a.(counter)
	howmuch, _ := strconv.Atoi(e.Payload)
	c.value = c.value - howmuch
	return c
}

func Test_ReadAggregateFromStream2(t *testing.T) {

	// https://tour.golang.org/methods/15
	//var i interface{} = "hello"

	stream := NewEventStream(
		NewEvent("counter", "counter:001", "incr", `5`),
		NewEvent("counter", "counter:001", "decr", `3`))

	handlers := map[string]func(Aggregate, Event) Aggregate{
		"incr": incr,
		"decr": decr}

	got := stream.LeftFold(counter{}, handlers).(counter)
	want := 2

	if got.value != want {
		t.Errorf("want:\n%+v\n, but got:\n%+v", want, got.value)
	}
}
