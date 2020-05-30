package persistence

import (
	"errors"

	"github.com/AwataKyosuke/go_api_server/domain"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
)

type eventPersistence struct{}

// NewEventPresistence TODO わかりやすいコメントを書きたい
func NewEventPresistence() repository.EventRepository {
	return &eventPersistence{}
}

func (p eventPersistence) GetAll() ([]*domain.Event, error) {
	var events []*domain.Event
	var error = errors.New("custom error")
	return events, error
}
