package domain

type AggregateRoot struct {
	Entity
	events *Set[string, Event]
}

func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{
		Entity: Entity{},
		events: NewSet[string, Event](),
	}
}

func (a *AggregateRoot) AddEvent(e Event) {
	a.events.Add(e.AggregateID(), e)
}

func (a *AggregateRoot) ClearEvents() {
	a.events = NewSet[string, Event]()
}
