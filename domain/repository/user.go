package repository

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
)

// IUserRepository 永続化を提供する処理を定義するインターフェース
type IUserRepository interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(userID int) (*model.User, error)
}
