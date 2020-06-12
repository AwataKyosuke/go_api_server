package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AwataKyosuke/go_api_server/interfaces/response"
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/AwataKyosuke/go_api_server/util/logger"
	"github.com/ant0ine/go-json-rest/rest"
)

// EventHandler Eventに関するハンドラーを定義するインターフェース
type EventHandler interface {
	GetEvents(rest.ResponseWriter, *rest.Request)
}

// eventHandler 依存しているUsecaseを返す
type eventHandler struct {
	eventUseCase usecase.EventUseCase
}

// NewEventHandler 依存性を注入しEventHandlerを作成する
func NewEventHandler(u usecase.EventUseCase) EventHandler {
	return &eventHandler{
		eventUseCase: u,
	}
}

// GetEvents イベントを取得する
func (h eventHandler) GetEvents(w rest.ResponseWriter, r *rest.Request) {

	// 必須パラメータチェック
	lat := r.URL.Query().Get("lat")
	if len(lat) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			response.Error{
				Message: "Latitude parameter is required.",
				Code:    400,
			})
		return
	}

	// 必須パラメータチェック
	lon := r.URL.Query().Get("lon")
	if len(lon) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			response.Error{
				Message: "Longitude parameter is required.",
				Code:    400,
			})
		return
	}

	// 必須パラメータチェック
	start := r.URL.Query().Get("start")
	if len(start) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			response.Error{
				Message: "Start parameter is required.",
				Code:    400,
			})
		return
	}

	// 必須パラメータチェック
	end := r.URL.Query().Get("end")
	if len(end) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			response.Error{
				Message: "End parameter is required.",
				Code:    400,
			})
		return
	}

	// パラメータフォーマットチェック
	convLat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			response.Error{
				Message: err.Error(),
				Code:    400,
			})
		return
	}

	// パラメータフォーマットチェック
	convLon, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			response.Error{
				Message: err.Error(),
				Code:    400,
			})
		return
	}

	// パラメータフォーマットチェック
	_, err = time.Parse("20060102", start)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			response.Error{
				Message: err.Error(),
				Code:    400,
			})
		return
	}

	// パラメータフォーマットチェック
	_, err = time.Parse("20060102", end)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			response.Error{
				Message: err.Error(),
				Code:    400,
			})
		return
	}

	// キーワードパラメータ取得
	keyword := r.URL.Query().Get("keyword")

	// 距離順に並び替えてイベントを取得する
	events, err := h.eventUseCase.GetEventsBySortedForDistance(convLat, convLon, start, end, keyword)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(
			response.Error{
				Message: err.Error(),
				Code:    500,
			})
		return
	}

	// 成功
	w.WriteHeader(http.StatusOK)
	w.WriteJson(&events)
}
