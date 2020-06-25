package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AwataKyosuke/go_api_server/domain/repository"
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

	// 緯度パラメータ必須チェック
	if len(r.URL.Query().Get("lat")) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: "Latitude parameter is required.",
				Code:    400,
			})
		return
	}

	// 経度パラメータ必須チェック
	if len(r.URL.Query().Get("lon")) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: "Longitude parameter is required.",
				Code:    400,
			})
		return
	}

	// 開始日パラメータ必須チェック
	if len(r.URL.Query().Get("start")) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: "Start parameter is required.",
				Code:    400,
			})
		return
	}

	// 終了日パラメータ必須チェック
	if len(r.URL.Query().Get("end")) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: "End parameter is required.",
				Code:    400,
			})
		return
	}

	// 取得件数パラメータ必須チェック
	if len(r.URL.Query().Get("count")) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: "Count parameter is required.",
				Code:    400,
			})
		return
	}

	// 緯度パラメータ取得
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: err.Error(),
				Code:    400,
			})
		return
	}

	// 経度パラメータ取得
	lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: err.Error(),
				Code:    400,
			})
		return
	}

	// 取得件数パラメータ取得
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: err.Error(),
				Code:    400,
			})
		return
	}

	// 開始日パラメータ取得
	start, err := time.Parse("20060102", r.URL.Query().Get("start"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: err.Error(),
				Code:    400,
			})
		return
	}

	// 終了日パラメータ取得
	end, err := time.Parse("20060102", r.URL.Query().Get("end"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: err.Error(),
				Code:    400,
			})
		return
	}

	// 開催方法がオンラインのデータのみを対象とするパラメータ
	online, err := strconv.ParseBool(r.URL.Query().Get("online"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: err.Error(),
				Code:    400,
			})
	}

	// 開催方法がオフラインのデータのみを対象とするパラメータ
	offline, err := strconv.ParseBool(r.URL.Query().Get("offline"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(
			Error{
				Message: err.Error(),
				Code:    400,
			})
	}

	// キーワードパラメータ取得
	keyword := r.URL.Query().Get("keyword")

	// パラメータを作成
	parameter := repository.EventSearchParameter{
		Lat:     lat,
		Lon:     lon,
		Start:   start.Format("20060102"),
		End:     end.Format("20060102"),
		Keyword: keyword,
		Count:   count,
		Online:  online,
		Offline: offline,
	}

	// 距離順に並び替えてイベントを取得する
	events, err := h.eventUseCase.GetEventsBySortedForDistance(parameter)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteJson(
			Error{
				Message: err.Error(),
				Code:    500,
			})
		return
	}

	// 成功
	w.WriteHeader(http.StatusOK)
	w.WriteJson(&events)
}
