package handler

import (
	"net/http"

	"github.com/AwataKyosuke/go_api_server/interfaces/respond"
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
)

// IAssetsHandler 必要なハンドラーを定義するインターフェース
type IAssetsHandler interface {
	Import(rest.ResponseWriter, *rest.Request)
	GetAll(rest.ResponseWriter, *rest.Request)
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

type assetsToJSON struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Bank   string `json:"bank"`
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
	w.WriteHeader(http.StatusOK)
	w.WriteJson(nil)
}

func (h *assetsHandler) GetAll(w rest.ResponseWriter, r *rest.Request) {
	assets, err := h.usecase.GetAll()
	if err != nil {
		respond.InternalServerError(w, err.Error())
		return
	}
	ret := []*assetsToJSON{}
	for _, t := range assets {
		r := *&assetsToJSON{
			Name:   t.GetName(),
			Amount: t.GetAmount(),
			Bank:   t.GetBank(),
		}
		ret = append(ret, &r)
	}
	respond.Success(w, ret)
}
