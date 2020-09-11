package mysqlrepository

import (
	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/pkg/errors"
)

type userRepository struct{}

// NewUserRepository 依存性を注入しPresistenceを返す
func NewUserRepository() repository.IUserRepository {
	return &userRepository{}
}

// GetAll 全てのユーザーを取得する
func (u userRepository) GetUsers() ([]*model.User, error) {

	// dbとのコネクションを生成
	con, err := GetConnection()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// DBからの検索結果を代入する構造体
	users := []*model.User{}

	// 検索実行
	con.Find(&users)

	return users, nil
}

// GetUserByID UserIDに一致するユーザーを1件取得
func (u userRepository) GetUserByID(userID int) (*model.User, error) {

	// dbとのコネクションを生成
	con, err := GetConnection()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// DBからの検索結果を代入する構造体
	user := model.User{}

	// 検索実行
	if con.First(&user, userID).RecordNotFound() {
		return &user, errors.New("not found user")
	}

	return &user, nil
}
