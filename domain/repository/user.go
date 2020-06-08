package repository

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/jinzhu/gorm"
)

// UserRepository TODO わかりやすいコメントを書きたい
type UserRepository interface {
	GetAll(db *gorm.DB) ([]*model.User, error)
	GetUserByID(db *gorm.DB, userID int) (*model.User, error)
}
