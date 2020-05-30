package usecase

import (
	"github.com/AwataKyosuke/go_api_server/domain"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
)

// EventUseCase Eventに関するユースケースを定義するインターフェース
type EventUseCase interface {
	GetAll() ([]*domain.Event, error)
}

type eventUseCase struct {
	eventRepository repository.EventRepository
}

// NewEventUseCase Eventデータに対するUseCaseを生成
func NewEventUseCase(r repository.EventRepository) EventUseCase {
	return &eventUseCase{
		eventRepository: r,
	}
}

func (u eventUseCase) GetAll() ([]*domain.Event, error) {
	events, err := u.eventRepository.GetAll()
	return events, err
}
