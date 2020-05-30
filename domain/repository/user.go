package repository

import (
	"github.com/AwataKyosuke/go_api_server/domain"
	"github.com/jinzhu/gorm"
)

// UserRepository TODO わかりやすいコメントを書きたい
type UserRepository interface {
	GetAll(db *gorm.DB) ([]*domain.User, error)
	GetUserByID(db *gorm.DB, userID int) (*domain.User, error)
}
