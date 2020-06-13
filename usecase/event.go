package usecase

import (
	"sort"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/AwataKyosuke/go_api_server/util/calculate"
)

// EventUseCase Eventに関するユースケースを定義するインターフェース
type EventUseCase interface {
	GetEventsBySortedForDistance(parameter repository.EventSearchParameter) ([]*model.Event, error)
}

// eventUseCase ユースケースが依存するリポジトリ
type eventUseCase struct {
	eventRepository repository.EventRepository
}

// NewEventUseCase 依存性を注入
func NewEventUseCase(r repository.EventRepository) EventUseCase {
	return &eventUseCase{
		eventRepository: r,
	}
}

// GetEventsBySortedForDistance 距離が近い順に並び替えてイベントを取得する
func (u eventUseCase) GetEventsBySortedForDistance(parameter repository.EventSearchParameter) ([]*model.Event, error) {

	events, err := u.eventRepository.GetEvents(parameter)
	if err != nil {
		return nil, err
	}

	// 距離を取得
	for i := 0; i < len(events); i++ {
		events[i].Distance = calculate.GetDistance(parameter.Lat, parameter.Lon, events[i].Lat, events[i].Lon)
	}

	// 近い順にソート
	sort.Slice(events, func(i, j int) bool { return events[i].Distance < events[j].Distance })

	// 件数を制限
	events = events[0:parameter.Count]

	return events, nil
}
