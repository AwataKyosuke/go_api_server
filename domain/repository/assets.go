package repository

import "github.com/AwataKyosuke/go_api_server/domain/model"

// IAssetsRepository 必要なリポジトリを定義するインターフェース
type IAssetsRepository interface {
	Insert([]model.Assets) error
	All() ([]model.Assets, error)
}
