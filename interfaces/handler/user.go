package handler

import (
	"net/http"
	"strconv"

	"github.com/AwataKyosuke/go_api_server/interfaces/respond"
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
)

// IUserHandler Userに関するハンドラーを定義するインターフェース
type IUserHandler interface {
	GetUsers(w rest.ResponseWriter, r *rest.Request)
	GetUserByID(w rest.ResponseWriter, r *rest.Request)
}

// userHandler 依存しているユースケースを返す
type userHandler struct {
	userUseCase usecase.IUserUseCase
}

// NewUserHandler 依存性を注入しハンドラーを作成
func NewUserHandler(u usecase.IUserUseCase) IUserHandler {
	return &userHandler{
		userUseCase: u,
	}
}

// GetAllUser 全てのユーザーを返す
func (h userHandler) GetUsers(w rest.ResponseWriter, r *rest.Request) {

	// 全てのユーザーを取得
	users, err := h.userUseCase.GetAll()
	if err != nil {
		respond.InternalServerError(w, err.Error())
		return
	}

	// 成功
	w.WriteHeader(http.StatusOK)
	w.WriteJson(&users)
}

// GetUserByID ユーザーIDに一致するユーザーを返す
func (h userHandler) GetUserByID(w rest.ResponseWriter, r *rest.Request) {

	userID, err := strconv.Atoi(r.PathParam("id"))
	if err != nil {
		respond.BadRequest(w, err.Error())
		return
	}

	user, err := h.userUseCase.GetUserByID(userID)
	if err != nil {
		respond.InternalServerError(w, err.Error())
		return
	}

	// 成功
	w.WriteHeader(http.StatusOK)
	w.WriteJson(&user)
}
