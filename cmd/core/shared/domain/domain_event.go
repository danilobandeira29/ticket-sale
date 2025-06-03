package domain

import "time"

type Event interface {
	Name() string
	AggregateID() string
	OccurredAt() time.Time
	Version() uint
}

type DEvent struct {
	id, name   string
	occurredAt time.Time
	version    uint
}

func NewEvent(id, name string, occurredAt time.Time, version uint) *DEvent {
	return &DEvent{
		id:         id,
		name:       name,
		occurredAt: occurredAt,
		version:    version,
	}
}

func (p DEvent) Name() string {
	return p.name
}

func (p DEvent) AggregateID() string {
	return p.id
}

func (p DEvent) OccurredAt() time.Time {
	return p.occurredAt
}

func (p DEvent) Version() uint {
	return p.version
}
