package handler

import (
	"net/http"
	"strconv"

	"github.com/AwataKyosuke/go_api_server/infrastructure"
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
)

// UserHandler TODO わかりやすいコメントを書きたい
type UserHandler interface {
	GetAllUser(w rest.ResponseWriter, r *rest.Request)
	GetUserByID(w rest.ResponseWriter, r *rest.Request)
}

// userHandler TODO わかりやすいコメントを書きたい
type userHandler struct {
	userUseCase usecase.UserUseCase
}

// NewUserHandler TODO わかりやすいコメントを書きたい
func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: u,
	}
}

func (h userHandler) GetAllUser(w rest.ResponseWriter, r *rest.Request) {

	db := infrastructure.NewDB()
	connection := infrastructure.Connect(db)

	users, err := h.userUseCase.GetAll(connection)

	if err != nil {
		// ヘッダーに失敗ステータスを書き込む
		w.WriteHeader(http.StatusNotFound)

		// レスポンスボディを書き込み
		w.WriteJson(nil)
	}
	// ヘッダーに成功ステータスを書き込む
	w.WriteHeader(http.StatusOK)

	// レスポンスボディを書き込み
	w.WriteJson(&users)
}

func (h userHandler) GetUserByID(w rest.ResponseWriter, r *rest.Request) {

	userID, err := strconv.Atoi(r.PathParam("id"))

	if err != nil {
		// TODO パラメータエラー
	}

	db := infrastructure.NewDB()
	connection := infrastructure.Connect(db)

	user, err := h.userUseCase.GetUserByID(connection, userID)

	if err != nil {

		// ヘッダーに失敗ステータスを書き込む
		w.WriteHeader(http.StatusNotFound)

		// レスポンスボディを書き込み
		w.WriteJson(nil)
	}

	// ヘッダーに成功ステータスを書き込む
	w.WriteHeader(http.StatusOK)

	// レスポンスボディを書き込み
	w.WriteJson(&user)
}
