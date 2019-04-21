package model

import (
	"time"
)

// Subscribe Struct
type Subscribe struct {
	SubscribeID string    `json:"subscribe_id"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
}

// SubscribeList type
type SubscribeList []Subscribe

// NewSubscribe func
func NewSubscribe() *Subscribe {
	return &Subscribe{
		CreatedAt: time.Now(),
	}
}
