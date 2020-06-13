package repository

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
)

// EventSearchParameter イベント検索時のパラメータ
type EventSearchParameter struct {
	Lat     float64
	Lon     float64
	Start   string
	End     string
	Keyword string
	Count   int
}

// EventRepository 永続化を提供する処理を定義するインターフェース
type EventRepository interface {
	GetEvents(parameter EventSearchParameter) ([]*model.Event, error)
}
