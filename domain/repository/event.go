package repository

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
)

// EventSearchParameter イベント検索時のパラメータ
type EventSearchParameter struct {
	StartDate string
	EndDate   string
	Keyword   string
}

// EventRepository TODO わかりやすいコメントを書きたい
type EventRepository interface {
	GetEvents(parameter EventSearchParameter) ([]*model.Event, error)
}
