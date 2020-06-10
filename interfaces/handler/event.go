package handler

import (
	"net/http"
	"strconv"

	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
)

// EventHandler TODO わかりやすいコメントを書きたい
type EventHandler interface {
	GetEvents(rest.ResponseWriter, *rest.Request)
}

// eventHandler TODO わかりやすいコメントを書きたい
type eventHandler struct {
	eventUseCase usecase.EventUseCase
}

// NewEventHandler TODO わかりやすいコメントを書きたい
func NewEventHandler(u usecase.EventUseCase) EventHandler {
	return &eventHandler{
		eventUseCase: u,
	}
}

func (h eventHandler) GetEvents(w rest.ResponseWriter, r *rest.Request) {
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		// TODO パラメータエラー
	}
	lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil {
		// TOOD パラメータエラー
	}
	start := r.URL.Query().Get("start")
	if len(start) < 1 {
		// TOOD パラメータエラー
	}
	end := r.URL.Query().Get("end")
	if len(end) < 1 {
		// TOOD パラメータエラー
	}
	keyword := r.URL.Query().Get("keyword")
	events, err := h.eventUseCase.GetEventsBySortedForDistance(lat, lon, start, end, keyword)
	if err != nil {
		// TODO サーバー側のエラー
		return
	}
	// ヘッダーに成功ステータスを書き込む
	w.WriteHeader(http.StatusOK)
	// レスポンスボディを書き込み
	w.WriteJson(&events)
}
