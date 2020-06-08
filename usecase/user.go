package usecase

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/jinzhu/gorm"
)

// UserUseCase TODO わかりやすいコメントを書きたい
type UserUseCase interface {
	GetAll(db *gorm.DB) ([]*model.User, error)
	GetUserByID(db *gorm.DB, userID int) (*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// NewUserUseCase TODO わかりやすいコメントを書きたい
func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: r,
	}
}

func (u userUseCase) GetAll(db *gorm.DB) ([]*model.User, error) {
	users, err := u.userRepository.GetAll(db)
	return users, err
}

func (u userUseCase) GetUserByID(db *gorm.DB, userID int) (*model.User, error) {
	user, err := u.userRepository.GetUserByID(db, userID)
	return user, err
}
