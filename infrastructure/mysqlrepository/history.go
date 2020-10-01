package mysqlrepository

import (
	"time"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
	"github.com/pkg/errors"
)

type historyRepository struct{}

type history struct {
	ID          int
	Date        time.Time
	Content     string
	Amount      int
	Bank        string
	MajorType   string
	MediumType  string
	Memo        string
	DeleteFlag  bool
	CreatedAt   time.Time
	CreatedUser string
	UpdatedAt   time.Time
	UpdatedUser string
}

// NewHistoryRepository コンストラクタ
func NewHistoryRepository() repository.IHistoryRepository {
	return &historyRepository{}
}

func (h *historyRepository) Insert(data []*model.History) error {
	con, err := GetConnection()
	defer con.Close()

	if err != nil {
		return errors.WithStack(err)
	}
	for _, d := range data {
		history := &history{
			ID:          0,
			Date:        d.Date(),
			Content:     d.Content(),
			Amount:      d.Amount(),
			Bank:        d.Bank(),
			MajorType:   d.MajorType(),
			MediumType:  d.MediumType(),
			Memo:        d.Memo(),
			DeleteFlag:  false,
			CreatedAt:   time.Now(),
			CreatedUser: "todo",
			UpdatedAt:   time.Now(),
			UpdatedUser: "todo",
		}
		con.Create(history)
	}
	return nil
}

func (h *historyRepository) All() ([]*model.History, error) {
	con, err := GetConnection()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	histories := []*history{}

	con.Find(&histories)

	ret := []*model.History{}

	for _, h := range histories {
		r := &model.History{}
		r.SetDate(h.Date)
		r.SetContent(h.Content)
		r.SetAmount(h.Amount)
		r.SetBank(h.Bank)
		r.SetMajorType(h.MajorType)
		r.SetMediumType(h.MediumType)
		r.SetMemo(h.Memo)
		ret = append(ret, r)
	}

	return ret, nil
}
