package service

import (
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// IAssetsService 必要なサービスを定義するインターフェース
type IAssetsService interface {
	SearchFromHtml(multipart.File) ([]*model.Assets, error)
}

type assetsService struct{}

// NewAssetsService サービスのコンストラクタ
func NewAssetsService() IAssetsService {
	return &assetsService{}
}

func (s *assetsService) SearchFromHtml(file multipart.File) ([]*model.Assets, error) {

	// HTMLドキュメントの読み込み
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	assets := []*model.Assets{}

	// テーブルの各行を取得
	table := doc.Find("table.table-bordered")
	tbody := table.Find("tbody")
	tr := tbody.Find("tr")

	// 各行をループ
	tr.Each(func(i int, s *goquery.Selection) {

		// 最初の行になぜか不要なデータがあるのでスキップ
		if i == 0 {
			return
		}

		// データ保存用の構造体
		asset := model.NewAssets()

		// 各列を取得
		td := s.Find("td")

		// 各列をループ
		td.Each(func(i int, s *goquery.Selection) {

			// １列目は種類・名称
			if i == 0 {
				err := asset.SetName(s.Text())
				if err != nil {
					log.Println(err.Error())
					return
				}
			}

			// ２列目は残高
			if i == 1 {
				text := strings.Replace(s.Text(), "円", "", -1)
				text = strings.Replace(text, ",", "", -1)
				amount, err := strconv.Atoi(text)
				if err != nil {
					log.Println(err.Error())
					return
				}
				err = asset.SetAmount(amount)
				if err != nil {
					log.Println(err.Error())
					return
				}
			}

			// ３列目は金融機関
			if i == 2 {
				err = asset.SetBank(s.Text())
				if err != nil {
					log.Println(err.Error())
					return
				}
			}

		})

		// assetが有効なら追加する
		if asset.IsValid() {
			assets = append(assets, asset)
		}

	})
	return assets, nil
}
