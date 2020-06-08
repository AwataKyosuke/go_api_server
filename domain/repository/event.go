package repository

import "github.com/AwataKyosuke/go_api_server/domain/model"

// EventRepository TODO わかりやすいコメントを書きたい
type EventRepository interface {
	GetEvents(start string, end string, keyword string) ([]*model.Event, error)
}
