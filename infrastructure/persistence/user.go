package persistence

import (
	"errors"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/jinzhu/gorm"
)

type userPersistence struct{}

// NewUserPersistence TODO わかりやすいコメントを書きたい
func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (p userPersistence) GetAll(db *gorm.DB) ([]*model.User, error) {

	// DBからの検索結果を代入する構造体
	users := []*model.User{}

	// 検索実行
	db.Find(&users)

	if len(users) == 0 {
		return users, errors.New("not found users")
	}

	return users, nil
}

func (p userPersistence) GetUserByID(db *gorm.DB, userID int) (*model.User, error) {

	// DBからの検索結果を代入する構造体
	user := model.User{}

	// 検索実行
	if db.First(&user, userID).RecordNotFound() {
		return &user, errors.New("not found user")
	}

	return &user, nil
}
