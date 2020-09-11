package handler

import (
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
)

// IHistoryHandler 必要なハンドラーを定義したインターフェース
type IHistoryHandler interface {
	Import(rest.ResponseWriter, *rest.Request)
}

type handler struct {
	usecase usecase.IHistoryUseCase
}

func (h handler) Import(w rest.ResponseWriter, r *rest.Request) {}
