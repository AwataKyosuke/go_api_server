package usecase

import (
	"log"

	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/AwataKyosuke/go_api_server/domain/service"
)

// IAssetsUseCase 必要なユースケースを定義するインターフェース
type IAssetsUseCase interface {
	Import(session string) error
}

type assetsUseCase struct {
	repository repository.IAssetsRepository
	service    service.IAssetsService
}

// NewUseCase ユースケースのコンストラクタ
func NewUseCase(repository repository.IAssetsRepository, service service.IAssetsService) IAssetsUseCase {
	return &assetsUseCase{
		repository: repository,
		service:    service,
	}
}

func (u *assetsUseCase) Import(session string) error {
	assets, err := u.service.Search(session)
	if err != nil {
		return err
	}
	// TODO
	log.Println(assets)
	err = u.repository.Insert(assets)
	if err != nil {
		return err
	}
	return nil
}
