package model

import (
	"fmt"
	"time"
)

// Message - model type of a message.
type Message struct {
	ID           string
	Name         string
	Email        string
	Text         string
	CreationTime time.Time
}

// SetCreationTime - set the creation time field by string.
func (m *Message) SetCreationTime(value string) error {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return fmt.Errorf("wrong time format (expected layout=%s, value=%s): %w", time.RFC3339, value, err)
	}
	m.CreationTime = t
	return nil
}
