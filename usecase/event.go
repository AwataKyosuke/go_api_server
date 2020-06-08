package usecase

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
)

// EventUseCase Eventに関するユースケースを定義するインターフェース
type EventUseCase interface {
	GetEventsBySortedForDistance(lat float64, lon float64, start string, end string, keyword string) ([]*model.Event, error)
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

func (u eventUseCase) GetEventsBySortedForDistance(lat float64, lon float64, start string, end string, keyword string) ([]*model.Event, error) {
	events, err := u.eventRepository.GetEvents(start, end, keyword)
	if err != nil {

	}

	return events, nil
}
