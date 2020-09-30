package handler

import (
	"net/http"
	"time"

	"github.com/AwataKyosuke/go_api_server/interfaces/respond"
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
)

// IHistoryHandler 必要なハンドラーを定義したインターフェース
type IHistoryHandler interface {
	Import(http.ResponseWriter, *http.Request)
	GetAll(rest.ResponseWriter, *rest.Request)
}

type historyHandler struct {
	usecase usecase.IHistoryUseCase
}

func NewHistoryHandler(usecase usecase.IHistoryUseCase) IHistoryHandler {
	return &historyHandler{
		usecase: usecase,
	}
}

type historiesToJSON struct {
	Date       time.Time `json:"date"`
	Content    string    `json:"content"`
	Amount     int       `json:"amount"`
	Bank       string    `json:"bank"`
	MajorType  string    `json:"major_type"`
	MediumType string    `json:"medium_type"`
	Memo       string    `json:"memo"`
}

func (h historyHandler) Import(w http.ResponseWriter, r *http.Request) {

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

func (h historyHandler) GetAll(w rest.ResponseWriter, r *rest.Request) {

	responder := respond.NewGoJSONRestResponder(w)

	histories, err := h.usecase.GetAll()
	if err != nil {
		responder.InternalServerError(err.Error())
		return
	}
	ret := []*historiesToJSON{}
	for _, h := range histories {
		r := *&historiesToJSON{
			Date:       h.GetDate(),
			Content:    h.GetContent(),
			Amount:     h.GetAmount(),
			Bank:       h.GetBank(),
			MajorType:  h.GetMajorType(),
			MediumType: h.GetMediumType(),
			Memo:       h.GetMediumType(),
		}
		ret = append(ret, &r)
	}
	responder.Success(ret)
}
