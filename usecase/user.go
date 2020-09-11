package usecase

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
)

type UserUseCase interface {
	GetAll() ([]*model.User, error)
	GetUserByID(userID int) (*model.User, error)
}

// userUseCase Userが依存するリポジトリ
type userUseCase struct {
	userRepository repository.IUserRepository
}

// NewUserUseCase 依存性を注入しユースケースを返す
func NewUserUseCase(r repository.IUserRepository) UserUseCase {
	return &userUseCase{
		userRepository: r,
	}
}

// GetAll 全てのユーザーを取得する
func (u userUseCase) GetAll() ([]*model.User, error) {
	return u.userRepository.GetUsers()
}

// GetUserByID ユーザーIDに一致するユーザーを返す
func (u userUseCase) GetUserByID(userID int) (*model.User, error) {
	return u.userRepository.GetUserByID(userID)
}
