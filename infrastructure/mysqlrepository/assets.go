package mysqlrepository

import (
	"time"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/pkg/errors"
)

type assetsRepository struct{}

// NewAssetsRepository リポジトリのコンストラクタ
func NewAssetsRepository() repository.IAssetsRepository {
	return &assetsRepository{}
}

// Assets 資産情報
type assets struct {
	ID          int
	Name        string
	Amount      int
	Bank        string
	DeleteFlag  bool
	CreatedAt   time.Time
	CreatedUser string
	UpdatedAt   time.Time
	UpdatedUser string
}

func (a *assetsRepository) Insert(data []*model.Assets) error {
	con, err := GetConnection()
	defer con.Close()

	if err != nil {
		return errors.WithStack(err)
	}
	for _, d := range data {
		assets := &assets{
			ID:          0,
			Name:        d.GetName(),
			Amount:      d.GetAmount(),
			Bank:        d.GetBank(),
			DeleteFlag:  false,
			CreatedAt:   time.Now(),
			CreatedUser: "todo",
			UpdatedAt:   time.Now(),
			UpdatedUser: "todo",
		}
		con.Create(assets)
	}
	return nil
}

func (a *assetsRepository) All() ([]*model.Assets, error) {

	// dbとのコネクションを生成
	con, err := GetConnection()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// DBからの検索結果を代入する構造体
	assets := []*assets{}

	// 検索実行
	con.Find(&assets)

	ret := []*model.Assets{}

	for _, a := range assets {
		r := &model.Assets{}
		r.SetName(a.Name)
		r.SetAmount(a.Amount)
		r.SetBank(a.Bank)
		ret = append(ret, r)
	}

	return ret, nil
}
