package domain

import (
	"encoding/json"
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

func (u *UUID) String() string {
	return u.uid.String()
}

func (u *UUID) Equal(o UUID) bool {
	return u.uid.String() == o.uid.String()
}

func (u *UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.uid.String())
}

func (u *UUID) Scan(src interface{}) error {
	switch source := src.(type) {
	case []byte:
		parsed, err := uuid.ParseBytes(source)
		if err != nil {
			return fmt.Errorf("uuid parse slice of bytes: %w", err)
		}
		u.uid = parsed
		return nil
	case string:
		parsed, err := uuid.Parse(source)
		if err != nil {
			return fmt.Errorf("uuid parse string: %w", err)
		}
		u.uid = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan UUID from %T", source)
	}
}
