package model

import "time"

// Event イベント情報
type Event struct {
	EventID     int       `json:"event_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventURL    string    `json:"event_url"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`
	Limit       int       `json:"limit"`
	Address     string    `json:"address"`
	Place       string    `json:"place"`
	Position    struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lng"`
	} `json:"position"`
	Distance float64 `json:"distance"`
	Accepted int     `json:"accepted"`
	Waiting  int     `json:"waiting"`
}
