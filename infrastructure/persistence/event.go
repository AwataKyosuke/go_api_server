package persistence

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/AwataKyosuke/go_api_server/domain"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
)

type eventPersistence struct{}

// NewEventPresistence TODO わかりやすいコメントを書きたい
func NewEventPresistence() repository.EventRepository {
	return &eventPersistence{}
}

func (p eventPersistence) GetAll() ([]*domain.Event, error) {

	ret := []*domain.Event{}

	events, err := getConpassEvent()
	if err != nil {

	}

	ret = append(ret, events...)

	events, err = getDoorkeeperEvent()
	if err != nil {

	}

	ret = append(ret, events...)

	sort.Slice(ret, func(i, j int) bool { return ret[i].Distance < ret[j].Distance })

	return ret, nil
}

func getConpassEvent() ([]*domain.Event, error) {

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

	conpassURL := "https://connpass.com/api/v1/event/"

	resp, err := http.Get(conpassURL)
	if err != nil {

	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	apiResp := apiResponse{}

	err = json.Unmarshal(byteArray, &apiResp)
	if err != nil {

	}

	event := []*domain.Event{}

	for i := 0; i < len(apiResp.Events); i++ {
		tmp := apiResp.Events[i]
		addEvent := domain.Event{
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
		lat, _ := strconv.ParseFloat(tmp.Lat, 64)
		lon, _ := strconv.ParseFloat(tmp.Lon, 64)
		addEvent.Distance = getDistance(lat, lon, 35.690921, 139.70025799999996)
		event = append(event, &addEvent)
	}

	return event, nil
}

func getDoorkeeperEvent() ([]*domain.Event, error) {

	type apiResponse struct {
		Event struct {
			Title        string    `json:"title"`
			ID           int       `json:"id"`
			StartsAt     time.Time `json:"starts_at"`
			EndsAt       time.Time `json:"ends_at"`
			VenueName    string    `json:"venue_name"`
			Address      string    `json:"address"`
			Lat          string    `json:"lat"`
			Long         string    `json:"long"`
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

	doorkeeperURL := "https://api.doorkeeper.jp/events"

	resp, err := http.Get(doorkeeperURL)
	if err != nil {

	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	apiResp := []apiResponse{}

	err = json.Unmarshal(byteArray, &apiResp)
	if err != nil {

	}

	event := []*domain.Event{}

	for i := 0; i < len(apiResp); i++ {
		tmp := apiResp[i].Event
		addEvent := domain.Event{
			EventID:     tmp.ID,
			Title:       tmp.Title,
			Description: tmp.Description,
			EventURL:    tmp.PublicURL,
			StartedAt:   tmp.StartsAt,
			EndedAt:     tmp.EndsAt,
			Limit:       tmp.TicketLimit,
			Address:     tmp.Address,
			Place:       tmp.VenueName,
			Accepted:    tmp.Participants,
			Waiting:     tmp.Waitlisted,
		}
		lat, _ := strconv.ParseFloat(tmp.Lat, 64)
		lon, _ := strconv.ParseFloat(tmp.Long, 64)
		addEvent.Distance = getDistance(lat, lon, 35.690921, 139.70025799999996)
		event = append(event, &addEvent)
	}

	return event, nil
}

func getDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {

	if lat1 == 0 || lng1 == 0 {
		return 0
	}

	// 緯度経度をラジアンに変換
	rlat1 := lat1 * math.Pi / 180
	rlng1 := lng1 * math.Pi / 180
	rlat2 := lat2 * math.Pi / 180
	rlng2 := lng2 * math.Pi / 180

	// 2点の中心角(ラジアン)を求める
	a :=
		math.Sin(rlat1)*math.Sin(rlat2) +
			math.Cos(rlat1)*math.Cos(rlat2)*
				math.Cos(rlng1-rlng2)
	rr := math.Acos(a)

	earthRadius := 6378140.
	distance := earthRadius * rr
	return distance
}
