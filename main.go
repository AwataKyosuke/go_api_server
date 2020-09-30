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

	historyRepository := mysqlrepository.NewHistoryRepository()
	historyService := service.NewHistoryService()
	historyUseCase := usecase.NewHistoryUseCase(historyRepository, historyService)
	historyHandler := handler.NewHistoryHandler(historyUseCase)

	// ログ書き込み設定
	logger.Setting(config.Values.LogFile)

	// ルーティング設定
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return origin == "http://localhost:8080"
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})
	router, err := rest.MakeRouter(
		rest.Get("/users", userHandler.GetUsers),
		rest.Get("/users/:id", userHandler.GetUserByID),

		rest.Get("/assets", assetsHandler.GetAll),
		rest.Post("/assets", assetsHandler.Import),

		rest.Get("/histories", historyHandler.GetAll),
	)

	if err != nil {
		logger.Fatal(errors.WithStack(err))
	}

	api.SetApp(router)
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/upload/history", withCORS(historyHandler.Import))

	http.ListenAndServe(":8888", nil)
}

// withCORS CORSを有効にする
func withCORS(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		f(w, r)
	}
}
