package repository

import "github.com/AwataKyosuke/go_api_server/domain/model"

// IHistoryRepository 必要なリポジトリを定義したインターフェース
type IHistoryRepository interface {
	Insert([]model.History) error
	All() ([]model.History, error)
}

type historyRepository struct{}

func (r historyRepository) Insert(data []model.History) error {
	return nil
}

func (r historyRepository) All() ([]model.History, error) {
	return nil, nil
}
