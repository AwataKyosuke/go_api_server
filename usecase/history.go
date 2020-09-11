package usecase

import (
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/AwataKyosuke/go_api_server/domain/service"
)

// IHistoryUseCase 必要なユースケースを定義したインターフェース
type IHistoryUseCase interface {
	Import() error
}

type historyUseCase struct {
	repository repository.IHistoryRepository
	service    service.IHistoryService
}

func (u historyUseCase) Import() error {
	return nil
}
