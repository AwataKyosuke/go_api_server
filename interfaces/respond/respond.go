package respond

import (
	"encoding/json"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// IResponder レスポンスを返す際に使用する処理を定義するインターフェース
type IResponder interface {
	Success(json interface{})
	InternalServerError(message string)
	BadRequest(message string)
}

type netHTTPReponder struct {
	writer   http.ResponseWriter
	response response
}

type goJSONRestResponder struct {
	writer   rest.ResponseWriter
	response response
}

type response struct {
	Code int
	JSON interface{}
}

type errorResponse struct {
	Message string
}

// NewNetHTTPResponder net/httpパッケージを使う際のレスポンダー
func NewNetHTTPResponder(w http.ResponseWriter) IResponder {
	return &netHTTPReponder{
		writer:   w,
		response: response{},
	}
}

func (n *netHTTPReponder) Success(j interface{}) {
	n.response.Code = http.StatusOK
	n.response.JSON = j
	res, err := json.Marshal(n.response)
	if err != nil {
		http.Error(n.writer, err.Error(), http.StatusInternalServerError)
		return
	}
	n.writer.Header().Set("Content-Type", "application/json")
	n.writer.WriteHeader(http.StatusOK)
	n.writer.Write(res)
}

func (n *netHTTPReponder) InternalServerError(message string) {
	n.response.Code = http.StatusInternalServerError
	n.response.JSON = errorResponse{
		Message: message,
	}
	res, err := json.Marshal(n.response)
	if err != nil {
		http.Error(n.writer, err.Error(), http.StatusInternalServerError)
		return
	}
	n.writer.Header().Set("Content-Type", "application/json")
	n.writer.WriteHeader(http.StatusInternalServerError)
	n.writer.Write(res)
}

func (n *netHTTPReponder) BadRequest(message string) {
	n.response.Code = http.StatusBadRequest
	n.response.JSON = errorResponse{
		Message: message,
	}
	res, err := json.Marshal(n.response)
	if err != nil {
		http.Error(n.writer, err.Error(), http.StatusInternalServerError)
		return
	}
	n.writer.Header().Set("Content-Type", "application/json")
	n.writer.WriteHeader(http.StatusBadRequest)
	n.writer.Write(res)
}

// NewGoJSONRestResponder go-json-restパッケージを使う際のレスポンダー
func NewGoJSONRestResponder(w rest.ResponseWriter) IResponder {
	return &goJSONRestResponder{
		writer:   w,
		response: response{},
	}
}

// Success ステータスコード200としてレスポンスを返す。
func (g *goJSONRestResponder) Success(json interface{}) {
	g.writer.WriteHeader(http.StatusOK)
	g.response.Code = http.StatusOK
	g.response.JSON = json
	g.writer.WriteJson(g.response)
}

// InternalServerError ステータスコード500としてレスポンスを返す。
func (g *goJSONRestResponder) InternalServerError(message string) {
	g.writer.WriteHeader(http.StatusInternalServerError)
	g.response.Code = http.StatusInternalServerError
	g.response.JSON = errorResponse{
		Message: message,
	}
	g.writer.WriteJson(g.response)
}

// BadRequest ステータスコード400としてレスポンスを返す。
func (g *goJSONRestResponder) BadRequest(message string) {
	g.writer.WriteHeader(http.StatusBadRequest)
	g.response.Code = http.StatusBadRequest
	g.response.JSON = errorResponse{
		Message: message,
	}
	g.writer.WriteJson(g.response)
}
