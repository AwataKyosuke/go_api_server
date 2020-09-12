package handler

import (
	"github.com/AwataKyosuke/go_api_server/interfaces/respond"
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
)

// IAssetsHandler 必要なハンドラーを定義するインターフェース
type IAssetsHandler interface {
	Import(rest.ResponseWriter, *rest.Request)
}

type assetsHandler struct {
	usecase usecase.IAssetsUseCase
}

// NewAssetsHandler ハンドラーのコンストラクタ
func NewAssetsHandler(usecase usecase.IAssetsUseCase) IAssetsHandler {
	return &assetsHandler{
		usecase: usecase,
	}
}

func (h *assetsHandler) Import(w rest.ResponseWriter, r *rest.Request) {
	session := r.Header.Get("session")
	if len(session) == 0 {
		respond.BadRequest(w, "マネーフォワードにアクセスするために必要なセッション情報が入力されていません")
		return
	}
	err := h.usecase.Import(session)
	if err != nil {
		respond.InternalServerError(w, err.Error())
		return
	}
	respond.Success(w, nil)
}
