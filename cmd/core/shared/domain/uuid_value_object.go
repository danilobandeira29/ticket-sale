package domain

import (
	"fmt"
	"github.com/google/uuid"
)

type UUID struct {
	uid uuid.UUID
}

func NewUUID() (*UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("new uuid: %v", err)
	}
	return &UUID{uid: id}, nil
}

func NewUUIDFromString(v string) (*UUID, error) {
	id, err := uuid.Parse(v)
	if err != nil {
		return nil, fmt.Errorf("new uuid: %v", err)
	}
	return &UUID{uid: id}, nil
}

func (u UUID) String() string {
	return u.uid.String()
}

func (u UUID) Equal(o UUID) bool {
	return u.uid.String() == o.uid.String()
}
