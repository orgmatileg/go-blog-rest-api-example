package model

import (
	"time"
)

// Example Struct
type Example struct {
	ExampleID string    `json:"example_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Example list
type ExampleList []Example

// NewExample func
func NewExample() *Example {
	return &Example{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
