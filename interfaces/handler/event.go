package handler

import (
	"net/http"

	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
)

// EventHandler TODO わかりやすいコメントを書きたい
type EventHandler interface {
	HandleGetEvent(rest.ResponseWriter, *rest.Request)
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

func (h eventHandler) HandleGetEvent(w rest.ResponseWriter, r *rest.Request) {
	events, err := h.eventUseCase.GetAll()
	if err != nil {
		// エラー処理
		return
	}
	// ヘッダーに成功ステータスを書き込む
	w.WriteHeader(http.StatusOK)

	// レスポンスボディを書き込み
	w.WriteJson(&events)
}
