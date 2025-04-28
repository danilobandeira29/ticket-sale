package domain

import (
	"encoding/json"
	"fmt"
)

type Entity struct{}

func (e *Entity) String(v any) string {
	entityJson, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("error marshalling entity: %v", err)
	}
	return string(entityJson)
}
