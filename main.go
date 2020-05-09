// パッケージ名指定
package main

// 必要なライブラリのインポート
import (
	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

// DB接続情報を持つ構造体
type Impl struct {
	DB *gorm.DB
}

// テーブル情報の構造体
type User struct {
	Id          int
	Username    string
	Email       string
	Password    string
	DeleteFlag  bool
	CreatedAt   time.Time
	CreatedUser string
	UpdatedAt   time.Time
	UpdatedUser string
}

func main() {

	// おまじない、、
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	// ルーティング設定
	router, err := rest.MakeRouter(
		rest.Get("/users", GetAllUser),
		rest.Get("/users/:id", GetUserById),
	)

	if err != nil {
		log.Fatal(err)
	}

	// サーバー起動
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8888", api.MakeHandler()))
}

// データベース初期処理
func (i *Impl) InitDB() {
	// エラーオブジェクト
	var err error
	// コネクションオープン
	i.DB, err = gorm.Open("mysql", "root:password@tcp(mysql:3306)/development?parseTime=true&&loc=Asia%2FTokyo&charset=utf8")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	i.DB.LogMode(true)
}

// /testにアクセスしたさいの処理
func GetAllUser(w rest.ResponseWriter, r *rest.Request) {

	// DB周り初期設定
	i := Impl{}

	// DBとの接続
	i.InitDB()

	// DBからの検索結果を代入する構造体
	users := []User{}

	// 検索実行
	i.DB.Find(&users)

	// ヘッダーに成功ステータスを書き込む
	w.WriteHeader(http.StatusOK)

	// レスポンスボディを書き込み
	w.WriteJson(&users)
}

func GetUserById(w rest.ResponseWriter, r *rest.Request) {

	// DB周り初期設定
	i := Impl{}

	// DBとの接続
	i.InitDB()

	// DBからの検索結果を代入する構造体
	user := User{}

	// 検索実行
	i.DB.First(&user, r.PathParam("id"))

	// ヘッダーに成功ステータスを書き込む
	w.WriteHeader(http.StatusOK)

	// レスポンスボディを書き込み
	w.WriteJson(&user)
}
