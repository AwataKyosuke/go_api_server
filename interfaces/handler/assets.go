package handler

import (
	"net/http"

	"github.com/AwataKyosuke/go_api_server/interfaces/respond"
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
)

// IAssetsHandler 必要なハンドラーを定義するインターフェース
type IAssetsHandler interface {
	Import(http.ResponseWriter, *http.Request)
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

func (h *assetsHandler) Import(w http.ResponseWriter, r *http.Request) {

	responder := respond.NewNetHTTPResponder(w)

	file, _, err := r.FormFile("report")

	if err != nil {
		responder.BadRequest(err.Error())
		return
	}

	err = h.usecase.Import(file)
	if err != nil {
		responder.InternalServerError(err.Error())
		return
	}

	responder.Success(nil)
}

func (h *assetsHandler) GetAll(w rest.ResponseWriter, r *rest.Request) {

	responder := respond.NewGoJSONRestResponder(w)

	assets, err := h.usecase.GetAll()
	if err != nil {
		responder.InternalServerError(err.Error())
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
	responder.Success(ret)
}
