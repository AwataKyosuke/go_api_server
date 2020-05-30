package repository

import "github.com/AwataKyosuke/go_api_server/domain"

// EventRepository TODO わかりやすいコメントを書きたい
type EventRepository interface {
	GetAll() ([]*domain.Event, error)
}
