package repository

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
)

// EventSearchParameter イベント検索時のパラメータ
// TODO こいつをどこに定義するべきなのか分からない
type EventSearchParameter struct {
	Lat     float64
	Lon     float64
	Start   string
	End     string
	Keyword string
	Count   int
	Online  bool
	Offline bool
}

// IEventRepository 永続化を提供する処理を定義するインターフェース
type IEventRepository interface {
	GetEvents(parameter EventSearchParameter) ([]*model.Event, error)
}
