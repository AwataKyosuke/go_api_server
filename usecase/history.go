package usecase

import (
	"mime/multipart"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/AwataKyosuke/go_api_server/domain/service"
)

// IHistoryUseCase 必要なユースケースを定義したインターフェース
type IHistoryUseCase interface {
	Import(multipart.File) error
	GetAll() ([]*model.History, error)
}

type historyUseCase struct {
	repository repository.IHistoryRepository
	service    service.IHistoryService
}

// NewHistoryUseCase コンストラクタ
func NewHistoryUseCase(repository repository.IHistoryRepository, service service.IHistoryService) IHistoryUseCase {
	return &historyUseCase{
		repository: repository,
		service:    service,
	}
}

func (u historyUseCase) Import(file multipart.File) error {
	histories, err := u.service.SearchFromHTML(file)
	if err != nil {
		return err
	}
	err = u.repository.Insert(histories)
	if err != nil {
		return err
	}
	return nil
}

func (u historyUseCase) GetAll() ([]*model.History, error) {
	histories, err := u.repository.All()
	if err != nil {
		return nil, err
	}
	return histories, nil
}
