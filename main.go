// パッケージ名指定
package main

// 必要なライブラリのインポート
import (
	"net/http"

	"github.com/AwataKyosuke/go_api_server/domain/service"
	"github.com/AwataKyosuke/go_api_server/infrastructure/mysqlrepository"
	"github.com/AwataKyosuke/go_api_server/interfaces/handler"
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/AwataKyosuke/go_api_server/util/config"
	"github.com/AwataKyosuke/go_api_server/util/logger"
	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {

	// 依存性の注入
	userRepository := mysqlrepository.NewUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	assetsRepository := mysqlrepository.NewAssetsRepository()
	assetsService := service.NewAssetsService()
	assetsUseCase := usecase.NewAssetsUseCase(assetsRepository, assetsService)
	assetsHandler := handler.NewAssetsHandler(assetsUseCase)

	// ログ書き込み設定
	logger.Setting(config.Values.LogFile)

	// ルーティング設定
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/users", withCORS(userHandler.GetUsers)),
		rest.Get("/users/:id", withCORS(userHandler.GetUserByID)),

		rest.Get("/assets", withCORS(assetsHandler.GetAll)),
		rest.Post("/assets", withCORS(assetsHandler.Import)),
	)

	if err != nil {
		logger.Fatal(errors.WithStack(err))
	}

	// サーバー起動
	api.SetApp(router)
	http.ListenAndServe(":8888", api.MakeHandler())
}

// withCORS CORSを有効にする
func withCORS(f rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		f(w, r)
	}
}
