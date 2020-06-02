// パッケージ名指定
package main

// 必要なライブラリのインポート
import (
	"log"
	"net/http"

	"github.com/AwataKyosuke/go_api_server/infrastructure/persistence"
	"github.com/AwataKyosuke/go_api_server/interfaces/handler"
	"github.com/AwataKyosuke/go_api_server/usecase"
	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	eventPersistence := persistence.NewEventPresistence()
	eventUseCase := usecase.NewEventUseCase(eventPersistence)
	eventHandler := handler.NewEventHandler(eventUseCase)

	// おまじない、、
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	// ルーティング設定
	router, err := rest.MakeRouter(
		rest.Get("/users", userHandler.GetAllUser),
		rest.Get("/users/:id", userHandler.GetUserByID),
		rest.Get("/events", eventHandler.HandleGetEvent),
	)

	if err != nil {
		log.Fatal(err)
	}

	// サーバー起動
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8888", api.MakeHandler()))
}
