package service

import (
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// IHistoryService 必要なサービスを定義したインターフェース
type IHistoryService interface {
	SearchFromHTML(multipart.File) ([]*model.History, error)
}

type service struct{}

// NewHistoryService コンストラクタ
func NewHistoryService() IHistoryService {
	return &service{}
}

// SearchFromHTML HTMLファイルから必要な履歴情報を取得する
func (s service) SearchFromHTML(file multipart.File) ([]*model.History, error) {

	// HTMLドキュメントの読み込み
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	histories := []*model.History{}

	// テーブルの各行を取得
	table := doc.Find("table#cf-detail-table")
	tbody := table.Find("tbody")
	tr := tbody.Find("tr")

	// 各行をループ
	tr.Each(func(i int, s *goquery.Selection) {

		// 各列を取得
		td := s.Find("td")

		// データ保存用の構造体
		history := model.NewHistory()

		// 各列をループ
		td.Each(func(i int, s *goquery.Selection) {
			if i == 1 {
				date := strings.ReplaceAll(strings.TrimSpace(s.Find("div").Text()), "\n", "")[:5]
				date = strconv.Itoa(time.Now().Year()) + "/" + date
				layout := "2006/01/02"
				t, _ := time.Parse(layout, date)
				history.SetDate(t)
			}
			if i == 2 {
				history.SetContent(strings.ReplaceAll(strings.TrimSpace(s.Find("div").Text()), "\n", ""))
			}
			if i == 3 {
				amt, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(s.Find("div").Text()), "\n", ""), ",", ""))
				history.SetAmount(amt)
			}
			if i == 4 {
				bank := strings.ReplaceAll(strings.TrimSpace(s.Find("div.noform").Text()), "\n", "")
				if len(bank) > 0 {
					history.SetBank(bank)
					return
				}
				bank = strings.ReplaceAll(strings.TrimSpace(s.Text()), "\n", "")
				if len(bank) > 0 {
					history.SetBank(bank)
					return
				}
			}
			if i == 5 {
				history.SetMajorType(strings.ReplaceAll(strings.TrimSpace(s.Find("div").Text()), "\n", ""))
			}
			if i == 6 {
				history.SetMediumType(strings.ReplaceAll(strings.TrimSpace(s.Find("div").Text()), "\n", ""))
			}
			if i == 7 {
				history.SetMemo(strings.ReplaceAll(strings.TrimSpace(s.Find("div").Text()), "\n", ""))
			}
		})

		// assetが有効なら追加する
		histories = append(histories, history)

	})
	return histories, nil
}
