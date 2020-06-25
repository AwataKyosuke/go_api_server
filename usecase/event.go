package usecase

import (
	"sort"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/AwataKyosuke/go_api_server/domain/service"
	"github.com/AwataKyosuke/go_api_server/util/calculate"
)

// EventUseCase Eventに関するユースケースを定義するインターフェース
type EventUseCase interface {
	GetEventsBySortedForDistance(parameter repository.EventSearchParameter) ([]*model.Event, error)
}

// eventUseCase ユースケースが依存するリポジトリ
type eventUseCase struct {
	eventRepository repository.IEventRepository
}

// NewEventUseCase 依存性を注入
func NewEventUseCase(r repository.IEventRepository) EventUseCase {
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

	// 開催方法による絞り込み
	events = service.Filtered(events, parameter.Online, parameter.Offline)

	// 距離を取得
	for _, event := range events {
		event.Distance = calculate.GetDistance(parameter.Lat, parameter.Lon, event.Position.Lat, event.Position.Lon)
	}

	// 近い順にソート
	sort.Slice(events, func(i, j int) bool { return events[i].Distance < events[j].Distance })

	// 件数を制限
	if len(events) > parameter.Count {
		events = events[0:parameter.Count]
	}

	return events, nil
}
