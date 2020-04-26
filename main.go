// パッケージ名指定
package main

// 必要なライブラリのインポート
import(
		"log"
		"net/http"
		"github.com/ant0ine/go-json-rest/rest"
    _ "github.com/go-sql-driver/mysql"
		"github.com/jinzhu/gorm"
)

// DB接続情報を持つ構造体
type Impl struct {
	DB *gorm.DB
}

// テーブル情報の構造体
type User struct {
	Id   int  	`json:id`
	Name string `json:name`
}

func main() {

	// DB周り初期設定
	i := Impl{}
	// DBとの接続
	i.InitDB()

	// おまじない、、
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	// ルーティング設定
	router, err := rest.MakeRouter(
			rest.Get("/test", i.GetTestMessage),
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
	i.DB, err = gorm.Open("mysql", "root:password@tcp(mysql:3306)/test?parseTime=true&&loc=Asia%2FTokyo&charset=utf8")
	if err != nil {
			log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	i.DB.LogMode(true)
}

// /testにアクセスしたさいの処理
func (i *Impl) GetTestMessage(w rest.ResponseWriter, r *rest.Request) {

	// DBからの検索結果を代入する構造体
	users := []User{}

	// 検索実行
	i.DB.Find(&users)

	// ヘッダーに成功ステータスを書き込む
	w.WriteHeader(http.StatusOK)

	// レスポンスボディを書き込み
	w.WriteJson(&users)
}
