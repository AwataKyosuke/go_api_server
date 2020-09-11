package service

import "github.com/AwataKyosuke/go_api_server/domain/model"

// IHistoryService 必要なサービスを定義したインターフェース
type IHistoryService interface {
	Search() ([]model.History, error)
}

type service struct{}

func (s service) Search() ([]model.History, error) {
	return nil, nil
}
