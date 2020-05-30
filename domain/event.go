package domain

import "time"

// Event イベント情報
type Event struct {
	EventID     string
	Title       string
	Catch       string
	Description string
	EventURL    string
	HashTag     string
	StartedAt   time.Time
	EndedAt     time.Time
	Limit       int
	Address     string
	Place       string
	Lat         float32
	Lon         float32
	Accepted    int
	Waiting     int
}
