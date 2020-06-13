package persistence

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type userPersistence struct{}

// NewUserPersistence 依存性を注入しPresistenceを返す
func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

// GetAll 全てのユーザーを取得する
func (p userPersistence) GetUsers(db *gorm.DB) ([]*model.User, error) {

	// DBからの検索結果を代入する構造体
	users := []*model.User{}

	// 検索実行
	db.Find(&users)

	return users, nil
}

// GetUserByID UserIDに一致するユーザーを1件取得
func (p userPersistence) GetUserByID(db *gorm.DB, userID int) (*model.User, error) {

	// DBからの検索結果を代入する構造体
	user := model.User{}

	// 検索実行
	if db.First(&user, userID).RecordNotFound() {
		return &user, errors.New("not found user")
	}

	return &user, nil
}
