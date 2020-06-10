package persistence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/AwataKyosuke/go_api_server/domain/model"
	"github.com/AwataKyosuke/go_api_server/domain/repository"
)

type eventPersistence struct{}

// NewEventPresistence TODO わかりやすいコメントを書きたい
func NewEventPresistence() repository.EventRepository {
	return &eventPersistence{}
}

func (p eventPersistence) GetEvents(parameter repository.EventSearchParameter) ([]*model.Event, error) {

	ret := []*model.Event{}

	events, err := getConpassEvent(parameter)
	if err != nil {

	}

	ret = append(ret, events...)

	events, err = getDoorkeeperEvent(parameter)
	if err != nil {

	}

	ret = append(ret, events...)

	return ret, nil
}

func getConpassEvent(parameter repository.EventSearchParameter) ([]*model.Event, error) {

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

	conpassURL := "https://connpass.com/api/v1/event/?"

	start, _ := time.Parse("20060102", parameter.StartDate)
	end, _ := time.Parse("20060102", parameter.EndDate)
	for j := 0; j < int(end.Sub(start).Hours()/24); j++ {
		conpassURL += "ymd="
		conpassURL += start.AddDate(0, 0, j).Format("20060102")
		conpassURL += "&"
	}

	if len(parameter.Keyword) > 0 {
		conpassURL += "keyword_or="
		conpassURL += parameter.Keyword
		conpassURL += "&"
	}

	fmt.Println(conpassURL)

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
		addEvent.Lat, _ = strconv.ParseFloat(tmp.Lat, 64)
		addEvent.Lon, _ = strconv.ParseFloat(tmp.Lon, 64)
		event = append(event, &addEvent)
	}

	return event, nil
}

func getDoorkeeperEvent(parameter repository.EventSearchParameter) ([]*model.Event, error) {

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

	doorkeeperURL := "https://api.doorkeeper.jp/events?"

	if len(parameter.Keyword) > 0 {
		doorkeeperURL += "q="
		doorkeeperURL += parameter.Keyword
		doorkeeperURL += "&"
	}

	doorkeeperURL += "since="
	doorkeeperURL += parameter.StartDate
	doorkeeperURL += "&"

	doorkeeperURL += "until="
	doorkeeperURL += parameter.EndDate
	doorkeeperURL += "&"

	fmt.Println(doorkeeperURL)

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
			Place:       tmp.VenueName,
			Accepted:    tmp.Participants,
			Waiting:     tmp.Waitlisted,
		}
		addEvent.Lat, _ = strconv.ParseFloat(tmp.Lat, 64)
		addEvent.Lon, _ = strconv.ParseFloat(tmp.Long, 64)
		event = append(event, &addEvent)
	}

	return event, nil
}
