package domain

import "time"

// Event イベント情報
type Event struct {
	EventID     int
	Title       string
	Description string
	EventURL    string
	StartedAt   time.Time
	EndedAt     time.Time
	Limit       int
	Address     string
	Place       string
	Distance    float64
	Accepted    int
	Waiting     int
}
