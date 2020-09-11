package repository

import "github.com/AwataKyosuke/go_api_server/domain/model"

// IAssetsRepository 必要なリポジトリを定義するインターフェース
type IAssetsRepository interface {
	Insert([]model.Assets) error
	All() ([]model.Assets, error)
}

type assetsRepository struct{}

// NewRepository リポジトリのコンストラクタ
func NewRepository() IAssetsRepository {
	return &assetsRepository{}
}

func (a *assetsRepository) Insert(data []model.Assets) error {
	return nil
}

func (a *assetsRepository) All() ([]model.Assets, error) {
	return nil, nil
}
