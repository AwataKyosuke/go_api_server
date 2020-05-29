// パッケージ名指定
package main

// 必要なライブラリのインポート
import (
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// User ユーザー情報
type User struct {
	ID          int
	Username    string
	Email       string
	Password    string
	DeleteFlag  bool
	CreatedAt   time.Time
	CreatedUser string
	UpdatedAt   time.Time
	UpdatedUser string
}

// DbConfig データベース接続情報
type DbConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
}

// DbError データベースに関するエラー情報
type DbError struct {
	TableName      string
	ParameterKey   string
	ParameterValue interface{}
	Message        string
}

func main() {

	// おまじない、、
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	// ルーティング設定
	router, err := rest.MakeRouter(
		rest.Get("/users", GetAllUser),
		rest.Get("/users/:id", GetUserByID),
	)

	if err != nil {
		log.Fatal(err)
	}

	// サーバー起動
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8888", api.MakeHandler()))
}

// GetConnection データベースとのコネクションをオープンする
func GetConnection() *gorm.DB {

	config := &DbConfig{
		Host:     "mysql",
		Username: "root",
		Password: "password",
		DBName:   "development",
	}

	db, err := gorm.Open("mysql", config.Username+":"+config.Password+"@tcp("+config.Host+")/"+config.DBName+"?parseTime=true&&loc=Asia%2FTokyo&charset=utf8")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	return db
}

// GetAllUser 全てのユーザー情報を取得する
func GetAllUser(w rest.ResponseWriter, r *rest.Request) {

	connection := GetConnection()

	// DBからの検索結果を代入する構造体
	users := []User{}

	// 検索実行
	connection.Find(&users)

	if len(users) == 0 {

		// ヘッダーに失敗ステータスを書き込む
		w.WriteHeader(http.StatusNotFound)

		// エラー情報
		dbError := DbError{
			TableName: "users",
			Message:   "レコードが1件もありません。",
		}

		// レスポンスボディを書き込み
		w.WriteJson(dbError)
	} else {

		// ヘッダーに成功ステータスを書き込む
		w.WriteHeader(http.StatusOK)

		// レスポンスボディを書き込み
		w.WriteJson(&users)
	}

}

// GetUserByID IDに一致するユーザーを取得する
func GetUserByID(w rest.ResponseWriter, r *rest.Request) {

	// データベースとのコネクション生成
	connection := GetConnection()

	// DBからの検索結果を代入する構造体
	user := User{}

	// 検索実行
	if connection.First(&user, r.PathParam("id")).RecordNotFound() {

		// ヘッダーに失敗ステータスを書き込む
		w.WriteHeader(http.StatusNotFound)

		// エラー情報
		dbError := DbError{
			TableName:      "users",
			ParameterKey:   "id",
			ParameterValue: r.PathParam("id"),
			Message:        "対象のレコードが見つかりませんでした。",
		}

		// レスポンスボディを書き込み
		w.WriteJson(dbError)
	} else {

		// ヘッダーに成功ステータスを書き込む
		w.WriteHeader(http.StatusOK)

		// レスポンスボディを書き込み
		w.WriteJson(&user)
	}
}
