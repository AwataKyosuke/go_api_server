package database

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
)

type eventRepository struct{}

// NewEventRepository 依存性を注入しPresistenceを返す
func NewEventRepository() repository.IEventRepository {
	return &eventRepository{}
}

// GetEvents イベント情報を取得する
func (u eventRepository) GetEvents(parameter repository.EventSearchParameter) ([]*model.Event, error) {

	ret := []*model.Event{}

	// conpassからイベントを検索
	events, err := getConpassEvent(parameter)
	if err != nil {
		return nil, err
	}
	ret = append(ret, events...)

	// doorkeeperからイベントを検索
	events, err = getDoorkeeperEvent(parameter)
	if err != nil {
		return nil, err
	}
	ret = append(ret, events...)

	return ret, nil
}

// getConpassEvent conpassからイベントを検索する
func getConpassEvent(parameter repository.EventSearchParameter) ([]*model.Event, error) {

	// apiResponse APIのレスポンス
	type apiResponse struct {
		ResultsStart     int `json:"results_start"`
		ResultsReturned  int `json:"results_returned"`
		ResultsAvailable int `json:"results_available"`
		Events           []struct {
			EventID     int       `json:"event_id"`
			Title       string    `json:"title"`
			Catch       string    `json:"catch"`
			Description string    `json:"description"`
			EventURL    string    `json:"event_url"`
			HashTag     string    `json:"hash_tag"`
			StartedAt   time.Time `json:"started_at"`
			EndedAt     time.Time `json:"ended_at"`
			Limit       int       `json:"limit"`
			EventType   string    `json:"event_type"`
			Series      struct {
				ID    int    `json:"id"`
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"series"`
			Address          string    `json:"address"`
			Place            string    `json:"place"`
			Lat              string    `json:"lat"`
			Lon              string    `json:"lon"`
			OwnerID          int       `json:"owner_id"`
			OwnerNickname    string    `json:"owner_nickname"`
			OwnerDisplayName string    `json:"owner_display_name"`
			Accepted         int       `json:"accepted"`
			Waiting          int       `json:"waiting"`
			UpdatedAt        time.Time `json:"updated_at"`
		} `json:"events"`
	}

	// アクセス先URL
	conpassURL := "https://connpass.com/api/v1/event/?"

	// 開始日と終了日のパラメータを設定
	start, _ := time.Parse("20060102", parameter.Start)
	end, _ := time.Parse("20060102", parameter.End)
	for j := 0; j < int(end.Sub(start).Hours()/24); j++ {
		conpassURL += "ymd="
		conpassURL += start.AddDate(0, 0, j).Format("20060102")
		conpassURL += "&"
	}

	// キーワードのパラメータを設定
	if len(parameter.Keyword) > 0 {
		conpassURL += "keyword="
		conpassURL += parameter.Keyword
		conpassURL += "&"
	}

	// 取得件数を100件に固定
	conpassURL += "count=100"

	// リクエスト送信
	resp, err := http.Get(conpassURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Bodyを読む
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// JSONを構造体にパース
	apiResp := apiResponse{}
	err = json.Unmarshal(byteArray, &apiResp)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 戻り値を作成
	event := []*model.Event{}
	for i := 0; i < len(apiResp.Events); i++ {
		tmp := apiResp.Events[i]
		addEvent := model.Event{
			EventID:     tmp.EventID,
			Title:       tmp.Title,
			Description: tmp.Description,
			EventURL:    tmp.EventURL,
			StartedAt:   tmp.StartedAt,
			EndedAt:     tmp.EndedAt,
			Limit:       tmp.Limit,
			Address:     tmp.Address,
			Place:       tmp.Place,
			Accepted:    tmp.Accepted,
			Waiting:     tmp.Waiting,
		}
		addEvent.Position.Lat, _ = strconv.ParseFloat(tmp.Lat, 64)
		addEvent.Position.Lon, _ = strconv.ParseFloat(tmp.Lon, 64)
		event = append(event, &addEvent)
	}

	return event, nil
}

// getDoorkeeperEvent doorkeeperからイベントを検索する
func getDoorkeeperEvent(parameter repository.EventSearchParameter) ([]*model.Event, error) {

	// apiResponse APIレスポンス
	type apiResponse struct {
		Event struct {
			Title        string    `json:"title"`
			ID           int       `json:"id"`
			StartsAt     time.Time `json:"starts_at"`
			EndsAt       time.Time `json:"ends_at"`
			VenueName    string    `json:"venue_name"`
			Address      string    `json:"address"`
			Lat          float64   `json:"lat"`
			Long         float64   `json:"long"`
			TicketLimit  int       `json:"ticket_limit"`
			PublishedAt  time.Time `json:"published_at"`
			UpdatedAt    time.Time `json:"updated_at"`
			Group        int       `json:"group"`
			Description  string    `json:"description"`
			PublicURL    string    `json:"public_url"`
			Participants int       `json:"participants"`
			Waitlisted   int       `json:"waitlisted"`
		} `json:"event"`
	}

	// APIアクセス先URL
	doorkeeperURL := "https://api.doorkeeper.jp/events?"

	// 開始日と終了日のパラメータを設定
	doorkeeperURL += "since="
	doorkeeperURL += parameter.Start
	doorkeeperURL += "&"

	doorkeeperURL += "until="
	doorkeeperURL += parameter.End
	doorkeeperURL += "&"

	// キーワードのパラメータを設定
	if len(parameter.Keyword) > 0 {
		doorkeeperURL += "q="
		doorkeeperURL += parameter.Keyword
		doorkeeperURL += "&"
	}

	// リクエスト送信
	resp, err := http.Get(doorkeeperURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Bodyを読む
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// JSONから構造体にパース
	apiResp := []apiResponse{}
	err = json.Unmarshal(byteArray, &apiResp)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 戻り値の作成
	event := []*model.Event{}
	for i := 0; i < len(apiResp); i++ {
		tmp := apiResp[i].Event
		addEvent := model.Event{
			EventID:     tmp.ID,
			Title:       tmp.Title,
			Description: tmp.Description,
			EventURL:    tmp.PublicURL,
			StartedAt:   tmp.StartsAt,
			EndedAt:     tmp.EndsAt,
			Limit:       tmp.TicketLimit,
			Address:     tmp.Address,
			Position: struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lng"`
			}{
				Lat: tmp.Lat,
				Lon: tmp.Long,
			},
			Place:    tmp.VenueName,
			Accepted: tmp.Participants,
			Waiting:  tmp.Waitlisted,
		}
		event = append(event, &addEvent)
	}

	return event, nil
}
