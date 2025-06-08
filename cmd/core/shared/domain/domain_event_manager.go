package domain

// this is a mediator

type Handler func(event Event)
type EventManager struct {
	emitter map[string][]Handler
}

func NewEventManager() *EventManager {
	return &EventManager{emitter: make(map[string][]Handler)}
}

func (e *EventManager) Register(event string, handler Handler) {
	e.emitter[event] = append(e.emitter[event], handler)
}

func (e *EventManager) Publish(agg *AggregateRoot) {
	for _, event := range agg.events.Data {
		handlers, ok := e.emitter[event.Name()]
		if !ok {
			continue
		}
		for _, handler := range handlers {
			handler(event)
		}
	}
	agg.ClearEvents()
}
