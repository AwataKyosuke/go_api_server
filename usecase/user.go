package usecase

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/jinzhu/gorm"
)

// UserUseCase Userに関するユースケースを定義するインターフェース
type UserUseCase interface {
	GetAll(db *gorm.DB) ([]*model.User, error)
	GetUserByID(db *gorm.DB, userID int) (*model.User, error)
}

// userUseCase Userが依存するリポジトリ
type userUseCase struct {
	userRepository repository.UserRepository
}

// NewUserUseCase 依存性を注入しユースケースを返す
func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: r,
	}
}

// GetAll 全てのユーザーを取得する
func (u userUseCase) GetAll(db *gorm.DB) ([]*model.User, error) {
	return u.userRepository.GetAll(db)
}

// GetUserByID ユーザーIDに一致するユーザーを返す
func (u userUseCase) GetUserByID(db *gorm.DB, userID int) (*model.User, error) {
	return u.userRepository.GetUserByID(db, userID)
}
