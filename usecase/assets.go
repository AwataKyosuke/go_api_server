package usecase

import (
	"mime/multipart"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/AwataKyosuke/go_api_server/domain/service"
)

// IAssetsUseCase 必要なユースケースを定義するインターフェース
type IAssetsUseCase interface {
	Import(multipart.File) error
	GetAll() ([]*model.Assets, error)
}

type assetsUseCase struct {
	repository repository.IAssetsRepository
	service    service.IAssetsService
}

// NewAssetsUseCase ユースケースのコンストラクタ
func NewAssetsUseCase(repository repository.IAssetsRepository, service service.IAssetsService) IAssetsUseCase {
	return &assetsUseCase{
		repository: repository,
		service:    service,
	}
}

func (u *assetsUseCase) Import(file multipart.File) error {
	assets, err := u.service.SearchFromHtml(file)
	if err != nil {
		return err
	}
	err = u.repository.Insert(assets)
	if err != nil {
		return err
	}
	return nil
}

func (u *assetsUseCase) GetAll() ([]*model.Assets, error) {
	assets, err := u.repository.All()
	if err != nil {
		return nil, err
	}
	return assets, nil
}
