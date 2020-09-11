// パッケージ名指定
package main

// 必要なライブラリのインポート
import (
	"net/http"

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

	// ログ書き込み設定
	logger.Setting(config.Config.LogFile)

	// ルーティング設定
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/users", userHandler.GetUsers),
		rest.Get("/users/:id", userHandler.GetUserByID),
	)

	if err != nil {
		logger.Fatal(errors.WithStack(err))
	}

	// サーバー起動
	api.SetApp(router)
	http.ListenAndServe(":8888", api.MakeHandler())
}
