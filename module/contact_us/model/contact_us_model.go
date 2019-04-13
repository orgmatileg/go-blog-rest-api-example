package model

import (
	"time"
)

// ContactUs Struct
type ContactUs struct {
	ContactUsID string    `json:"contact_us_id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Subject     string    `json:"subject"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
}

// ContactUs list
type ContactUsList []ContactUs

// ContactUs func
func NewContactUs() *ContactUs {
	return &ContactUs{
		CreatedAt: time.Now(),
	}
}
