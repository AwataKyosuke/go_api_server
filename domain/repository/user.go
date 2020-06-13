package repository

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/jinzhu/gorm"
)

// UserRepository 永続化を提供する処理を定義するインターフェース
type UserRepository interface {
	GetUsers(db *gorm.DB) ([]*model.User, error)
	GetUserByID(db *gorm.DB, userID int) (*model.User, error)
}
