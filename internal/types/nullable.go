package types

import (
	"encoding/json"
	"fmt"
)

// Nullable is a generic type that can be used to represent nullable values that
// are interopable with JSON.
type Nullable[T any] struct {
	Value T
	Valid bool
}

func (n Nullable[T]) String() string {
	if !n.Valid {
		return "null"
	}

	return fmt.Sprintf("%v", n.Value)
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Value)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	n.Valid = true
	n.Value = value

	return nil
}
